package conn

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/nanwp/jajan-yuk/user/config"
	"log"
)

type Cache struct {
	Pool *redis.Pool
}

type CacheService interface {
	Ping() error
	Get(key string) ([]byte, error)
	Set(key string, value []byte, ttl int64) error
	Exists(key string) (bool, error)
	Delete(key string) error
}

var cache CacheService

func NewCacheService(pool *redis.Pool) CacheService {
	if cache == nil {
		cache = &Cache{pool}
	}
	return cache
}

func CreateRedisPool(addr, password string, maxIdle int) (*redis.Pool, error) {
	redis := &redis.Pool{
		MaxIdle: maxIdle,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}

			if len(password) > 0 {
				if _, err := c.Do("AUTH", password); err != nil {
					return nil, err
				}

			}
			return c, nil
		},
	}
	conn := redis.Get()
	defer conn.Close()

	_, err := conn.Do("PING")
	if err != nil {
		return redis, err
	}

	return redis, nil
}

func (c *Cache) Ping() error {
	conn := c.Pool.Get()
	defer conn.Close()
	_, err := conn.Do("PING")
	if err != nil {
		return fmt.Errorf("error when ping to redis")
	}

	return err
}

func (c *Cache) Get(key string) ([]byte, error) {
	conn := c.Pool.Get()
	defer conn.Close()

	var data []byte
	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error getting key %s %v", key, err)
	}

	return data, nil
}

func (c *Cache) Set(key string, value []byte, ttl int64) error {
	conn := c.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
		v := string(value)
		if len(v) > 15 {
			v = v[0:12] + "..."
		}
		return fmt.Errorf("error setting key %s to %s: %v", key, value, err)
	}
	_, err = conn.Do("EXPIRE", key, ttl)

	if err != nil {
		return fmt.Errorf("error setting expire key %s to %s: %v", key, value, err)
	}

	return nil
}

func (c *Cache) Exists(key string) (bool, error) {
	conn := c.Pool.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return ok, fmt.Errorf("error checking key %s: %v", key, err)
	}
	return ok, nil
}

func (c *Cache) Delete(key string) error {
	conn := c.Pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	return err
}

func InitRedis(cfg config.Config) (CacheService, *redis.Pool) {
	redisAddress := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)
	pool, errPool := CreateRedisPool(redisAddress, cfg.RedisPassword, cfg.RedisMaxIdle)
	coreRedis := NewCacheService(pool)

	if errPool != nil {
		panic(errPool.Error())
	}

	log.Println("success connect to redis")

	return coreRedis, pool
}
