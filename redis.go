package m2rcache

import (
	"fmt"
	"time"

	"gopkg.in/redis.v3"
)

func createRedisDB(rc RedisCfg) *redisDB {
	r := &redisDB{}
	r.initDB(rc)
	return r
}

// RedisCfg redis配置
type RedisCfg struct {
	// Password 密码
	Password string `json:"password" yaml:"password"`
	// Address 数据库连接
	Address []string `json:"address" yaml:"address"`
	// DB 数据库选择
	DB int64 `json:"db" yaml:"db"`

	// 失败重试次数
	MaxRetries int `json:"maxRetries" yaml:"maxRetries"`

	//连接超时时间
	DialTimeout time.Duration `json:"dialTimeout" yaml:"dialTimeout"`
	//读取超时时间
	ReadTimeout time.Duration `json:"readTimeout" yaml:"readTimeout"`
	// 写入超时时间
	WriteTimeout time.Duration `json:"writeTimeout" yaml:"writeTimeout"`

	// 连接池连接数
	//默认10
	PoolSize int `json:"poolSize" yaml:"poolSize"`
	// 获取连接的超时时间
	// 所有连接繁忙返回错误
	// 默认5秒
	PoolTimeout time.Duration `json:"poolTimeout" yaml:"poolTimeout"`
	// 连接空闲关闭时间
	// 空闲时间应该小于超时时间(PoolSize)
	// 默认不进行空闲关闭
	IdleTimeout time.Duration `json:"idleTimeout" yaml:"idleTimeout"`
}

type redisDB struct {
	client *redis.Client
}

func (r *redisDB) initDB(rc RedisCfg) error {
	opt := &redis.Options{
		Addr:     rc.Address[0],
		Password: rc.Password,
		DB:       rc.DB,

		// The maximum number of retries before giving up.
		// Default is to not retry failed commands.
		MaxRetries: rc.MaxRetries,

		// Sets the deadline for establishing new connections. If reached,
		// dial will fail with a timeout.
		DialTimeout: rc.DialTimeout * time.Second,
		// Sets the deadline for socket reads. If reached, commands will
		// fail with a timeout instead of blocking.
		ReadTimeout: rc.ReadTimeout * time.Second,
		// Sets the deadline for socket writes. If reached, commands will
		// fail with a timeout instead of blocking.
		WriteTimeout: rc.WriteTimeout * time.Second,

		// The maximum number of socket connections.
		// Default is 10 connections.
		PoolSize: rc.PoolSize,
		// Specifies amount of time client waits for connection if all
		// connections are busy before returning an error.
		// Default is 5 seconds.
		PoolTimeout: rc.PoolTimeout * time.Second,
		// Specifies amount of time after which client closes idle
		// connections. Should be less than server's timeout.
		// Default is to not close idle connections.
		IdleTimeout: rc.IdleTimeout * time.Second,
	}
	client := redis.NewClient(opt)
	err := client.Ping().Err()
	if err != nil {
		return fmt.Errorf("连接redis服务器失败:%v", err)
	}
	r.client = client
	return nil
}
func (r *redisDB) save(key string, cacheTime time.Duration, field, value string, pairs ...string) {
	r.client.HMSet(key, field, value, pairs...)
	fmt.Println("cachetime:", cacheTime)
	if cacheTime > 0 {

		r.client.Expire(key, cacheTime)
	}
}
func (r *redisDB) get(key string, field ...string) ([]interface{}, error) {
	fmt.Printf("redis(%s):%v\n", key, field)
	rs := r.client.HMGet(key, field...)
	fmt.Printf("redis(%s):%v  err:%v\n", key, rs.Val(), rs.Err())
	return rs.Result()

}
