package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/jtzjtz/kit/conn/redis_pool"
	"github.com/jtzjtz/ys_api/app/config"
	"log"
	"time"
)

var FrameworkRedis *redis.Pool

func Init() {
	options := &redis_pool.Options{
		Host:               config.Configs.SITE_REDIS_SERVER,
		PassWord:           "",
		Database:           0,
		InitCap:            10,
		MaxCap:             100,
		IsWait:             true,
		IdleTimeout:        5 * time.Second,
		DialConnectTimeout: 5 * time.Second,
		DialReadTimeout:    5 * time.Second,
		DialWriteTimeout:   5 * time.Second,
	}

	var err error

	FrameworkRedis, err = redis_pool.NewRedisPool(options)

	if err != nil {
		log.Printf("%#v\n", err)
		return
	}
}
