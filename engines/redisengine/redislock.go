package redisengine

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

type RedisEngine struct {
	client *redis.Client
}

func New(addr, password string, db int) (*RedisEngine, error) {

	opt := &redis.Options{
		Addr:        addr,
		Password:    password, // no password set
		DB:          db,       // use default DB
		DialTimeout: time.Minute,
		ReadTimeout: time.Minute,
		IdleTimeout: time.Minute,
	}

	c := redis.NewClient(opt)

	if c == nil {
		return nil, errors.New("Redis connect failed:" + addr)
	}

	return &RedisEngine{client: c}, nil
}
func NewWithClient(c *redis.Client) (*RedisEngine, error) {

	if c == nil {
		return nil, errors.New("Redis connect failed:" + addr)
	}

	return &RedisEngine{client: c}, nil
}

func (e *RedisEngine) GetLock(lockKey string, expires time.Duration) (bool, error) {

	return e.client.SetNX(ctx, lockKey, 1, expires).Result()
}

func (e *RedisEngine) ReleaseLock(lockKey string) error {
	return e.client.Del(ctx, lockKey).Err()

}
