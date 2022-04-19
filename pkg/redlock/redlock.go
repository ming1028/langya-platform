package redlock

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

type Scripter interface {
	Eval(script string, keys []string, args ...interface{}) *redis.Cmd
	EvalSha(sha1 string, keys []string, args ...interface{}) *redis.Cmd
	ScriptExists(scripts ...string) *redis.BoolSliceCmd
	ScriptLoad(script string) *redis.StringCmd
	SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd
	Del(key ...string) *redis.IntCmd
}

type Locker interface {
	Lock(context.Context, string) (func(), error)
}

type lockOptions struct {
	keyPrefix    string
	expiry       time.Duration
	waitTime     time.Duration
	retryDelay   time.Duration
	scripters    []Scripter
	hasScriptCap bool // FIXME(oy): 无hasScript时，有限正确性
}

type locker struct {
	opts *lockOptions
}

func NewLocker(options ...LockOption) (Locker, error) {
	opts := &lockOptions{
		expiry:       10 * time.Second,
		waitTime:     500 * time.Millisecond,
		retryDelay:   5 * time.Millisecond,
		hasScriptCap: true,
	}

	for _, o := range options {
		o.Apply(opts)
	}

	if len(opts.scripters) < 1 {
		return nil, errors.New("scripter not configured")
	}

	if opts.hasScriptCap {
		// script 预加载
		err := scriptPreLoad(opts.scripters[0])
		if err != nil {
			return nil, err
		}
	}

	return &locker{opts}, nil
}

func (l *locker) Lock(ctx context.Context, key string) (
	releaseFunc func(), err error,
) {
	// 添加key前缀--避免冲突
	key = l.opts.keyPrefix + key

	// 生成随机值，避免碰撞
	value, err := genValue()
	if err != nil {
		return nil, err
	}

	// 释放锁
	releaseFunc = func() {
		if l.opts.hasScriptCap {
			releaseLockByScript(l.opts.scripters[0], key, value)
		} else {
			releaseLock(l.opts.scripters[0], key, value)
		}
	}

	// 添加recover，防止资源泄露
	defer func() {
		if revErr := recover(); revErr != any(nil) {
			releaseFunc()
			err = fmt.Errorf("recover: %v", revErr)
		}
	}()

	var bSucc bool
	start := time.Now()
	attempts := 0
	errCnt := 0
	for {
		select {
		// Check if we should quit
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		// 获取锁间隙，休眠
		if attempts > 0 && l.opts.retryDelay > time.Millisecond {
			time.Sleep(l.opts.retryDelay)
		}

		// 超时判断
		if attempts > 0 {
			elapsed := time.Since(start)
			if elapsed > l.opts.waitTime {
				err = context.DeadlineExceeded
				break
			}
		}
		attempts++

		// 获取锁
		bSucc, err = acquireLock(l.opts.scripters[0], key, value, l.opts.expiry)
		if err != nil {
			// 可重试三次，减少网络抖动发生时，失败的概率
			errCnt++
			if errCnt > 2 {
				break
			}
			err = nil
		}

		if bSucc {
			break
		}
	}

	if err != nil {
		return nil, err
	}

	// TODO(oysl): keepalive(保活) redlock

	select {
	// Check if we should quit
	case <-ctx.Done():
		if bSucc { // 上层调用取消，直接释放锁
			releaseFunc()
		}
		return nil, ctx.Err()
	default:
	}

	return releaseFunc, nil
}

type LockOption interface {
	Apply(*lockOptions)
}

type optionFunc func(*lockOptions)

func (o optionFunc) Apply(l *lockOptions) { o(l) }

func SetKeyPrefix(keyPrefix string) LockOption {
	return optionFunc(func(l *lockOptions) {
		l.keyPrefix = keyPrefix
	})
}

func SetExpiry(expiry time.Duration) LockOption {
	return optionFunc(func(l *lockOptions) {
		l.expiry = expiry
	})
}

func SetWaitTime(waitTime time.Duration) LockOption {
	return optionFunc(func(l *lockOptions) {
		l.waitTime = waitTime
	})
}

func SetRetryDelay(retryDelay time.Duration) LockOption {
	return optionFunc(func(l *lockOptions) {
		l.retryDelay = retryDelay
	})
}

func SetScripter(scripter ...Scripter) LockOption {
	return optionFunc(func(l *lockOptions) {
		l.scripters = scripter
	})
}

func SetScriptCap(hasScriptCap bool) LockOption {
	return optionFunc(func(l *lockOptions) {
		l.hasScriptCap = hasScriptCap
	})
}

func genValue() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// lua script 预加载
func scriptPreLoad(c Scripter) error {
	_, err := deleteScript.Load(c).Result()
	return err
}

func acquireLock(c Scripter, key, value string, expir time.Duration) (bool, error) {
	return c.SetNX(key, value, expir).Result()
}

var deleteScript = redis.NewScript(`
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("DEL", KEYS[1])
	else
		return 0
	end
`)

func releaseLock(c Scripter, key, value string) bool {
	status, err := c.Del(key).Result()
	return err == nil && status != 0
}

func releaseLockByScript(c Scripter, key, value string) bool {
	status, err := deleteScript.Run(c, []string{key}, value).Result()
	return err == nil && status != 0
}

/* NOTE(oy): 未支持保活
var touchScript = redis.NewScript(`
	if redis.call("GET", KEYS[1]) == ARGV[1] then
		return redis.call("SET", KEYS[1], ARGV[1], "XX", "PX", ARGV[2])
	else
		return "ERR"
	end
`)

func touchLock(c Scripter, key, value string, expir time.Duration) bool {
	status, err := touchScript.Run(c, []string{key}, value, int(expir/time.Millisecond)).Result()
	return err != nil && status != "ERR"
}*/
