package xzap

import (
	"context"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
	"time"
)

type zLog struct {
	log                    *zap.Logger
	outputPaths            []string
	closeOutputPathsFunc   func()
	errorOutPaths          []string
	closeErrorOutPathsFunc func()
	level                  zapcore.Level
	sugarLogger            *zap.SugaredLogger
	logOutputPath          string
	errorLogOutputPath     string
}

var zl *zLog

func InitZLog(
	outputPaths []string, level zapcore.Level, callerLevel ...int,
) (err error) {
	if zl != nil {
		return
	}
	errorOutPaths := []string{}
	logOutputPath := ""
	errLogOutPath := ""
	for k, v := range outputPaths {
		if v == "stderr" || v == "stdin" || v == "stdout" {
			continue
		}
		ps := strings.Split(v, ".")
		ps[0] = ps[0]
		v = strings.Join(ps, ".")

		outputPaths[k] = v
		errorOutPaths = append(errorOutPaths, v+".error")
		logOutputPath = v
		errLogOutPath = v + ".error"
	}

	zl = &zLog{
		outputPaths:        outputPaths,
		errorOutPaths:      errorOutPaths,
		level:              level,
		logOutputPath:      logOutputPath,
		errorLogOutputPath: errLogOutPath,
	}
	err = zl.init()
	if err != nil {
		return
	}
	log.Println("new zlog", outputPaths, errorOutPaths)
	// zl.splitByTime()

	zl.InitZapLog()
	return
}

func (this *zLog) init() (err error) {
	var sink zapcore.WriteSyncer
	// 初始化日志输入文件
	sink, this.closeOutputPathsFunc, err = zap.Open(this.outputPaths...)
	if err != nil {
		return
	}
	allWriter := zapcore.AddSync(sink)
	// 初始化错误日志输入文件
	sink, this.closeErrorOutPathsFunc, err = zap.Open(this.errorOutPaths...)
	if err != nil {
		return
	}
	errorWriter := zapcore.AddSync(sink)
	// 初始化日志格式配置
	config := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     this.timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	encoder := zapcore.NewJSONEncoder(config)

	// 一次写行为到多个输出端
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, allWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= this.level
		})),
		zapcore.NewCore(encoder, errorWriter, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= zapcore.ErrorLevel
		})),
	)

	this.log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	this.log = this.log.WithOptions(zap.AddCallerSkip(1))
	return
}

func (this *zLog) timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000000"))
}

func (this *zLog) splitByTime() {
	log.Println("do splitByTime")

	go func() {
		var lastSplitHour = -1
		for {
			time.Sleep(200 * time.Millisecond)
			// 整点切换文件
			if time.Now().Minute() == 59 {
				currHour := time.Now().Hour()
				if currHour == lastSplitHour {
					continue
				}
				lastSplitHour = currHour

				for _, file := range this.outputPaths {
					_, err := os.Stat(file)
					if err == nil {
						newFile := file + "." + time.Now().Format("2006-01-02_15")
						err = os.Rename(file, newFile)
						if err != nil {
							log.Println(err)
						} else {
							log.Println("RenameFile", newFile)
						}
					}
				}
				if currHour == 23 {
					for _, file := range this.errorOutPaths {
						_, err := os.Stat(file)
						if err == nil {
							newFile := file + "." + time.Now().Format("2006-01-02_15")
							err = os.Rename(file, newFile)
							if err != nil {
								log.Println(err)
							} else {
								log.Println("RenameFile", newFile)
							}
						}
					}
				}

				err := this.log.Sync()
				if err != nil {
					log.Println(err)
				}
				this.closeOutputPathsFunc()
				this.closeErrorOutPathsFunc()
				err = this.init()
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()
}

func Info(msg string, fields ...zap.Field) {
	// zl.log.Info(msg, fields...)
	zl.sugarLogger.Info(msg, fields)
}

func InfoContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getLogField(ctx))
	zl.log.Info(msg, fields...)
}

func getLogField(ctx context.Context) (field zap.Field) {
	log_id := ""
	if ctx.Value("X-Request-Id") != nil {
		log_id = ctx.Value("X-Request-Id").(string)
	}
	field = zap.String("X-Request-Id", log_id)
	return
}

func ErrorContext(ctx context.Context, msg string, fields ...zap.Field) {
	fields = append(fields, getLogField(ctx))
	zl.log.Error(msg, fields...)
}

func (this *zLog) InitZapLog() {
	writeSyncer := getLogWriter(this.logOutputPath)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	zl.sugarLogger = logger.Sugar()
}

func getLogWriter(fileName string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    1,     //切割文件大小 （MB）
		MaxBackups: 5,     // 文件最大个数
		MaxAge:     30,    // 保留最长时间 (天)
		Compress:   false, // 压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
