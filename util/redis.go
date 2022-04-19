package util

import (
	"fmt"
	"gitee.com/langya_platform/pkg/redlock"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var (
	redisCli     *redis.Client
	redisCliOnce sync.Once

	redLocker     redlock.Locker
	redLockerOnce sync.Once
)

func GetRedisCli() *redis.Client {
	redisCliOnce.Do(func() {
		var err error
		redisCli, err = newClient("redis")
		if err != nil {
			panic(any("failed to connect redis, error: " + err.Error()))
		}
	})

	return redisCli
}

func GetLocker() redlock.Locker {
	redLockerOnce.Do(func() {
		redisCli, err := newClient("redis")
		if err != nil {
			panic(any("failed to connect redis, error: " + err.Error()))
		}

		redLocker, err = redlock.NewLocker(
			redlock.SetScripter(redisCli),
			redlock.SetScriptCap(false),
		)
		if err != nil {
			panic(any("failed to newLocker" + err.Error()))
		}
	})

	return redLocker
}

func newClient(key string) (client *redis.Client, err error) {
	if !viper.IsSet(key) {
		err = fmt.Errorf(fmt.Sprintf("redis config nil: %s", key))
		return
	}
	client = redis.NewClient(
		&redis.Options{
			Addr:         viper.GetString(key + ".addr"),
			Password:     viper.GetString(key + ".password"),
			DB:           viper.GetInt(key + ".db"),
			DialTimeout:  viper.GetDuration(key+".dialTimeout") * time.Millisecond,
			ReadTimeout:  viper.GetDuration(key+".readTimeout") * time.Millisecond,
			WriteTimeout: viper.GetDuration(key+".writeTimeout") * time.Millisecond,
			MaxRetries:   viper.GetInt(key + ".maxRetries"),
			PoolSize:     viper.GetInt(key + ".poolSize"),
			MinIdleConns: viper.GetInt(key + ".minIdleConns"),
		})
	err = client.Ping().Err()
	if err != nil {
		log.Println("redis ping error", err)
	}
	return
}
