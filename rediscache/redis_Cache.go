package main

/*
import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"time"
)

var (
	ErrInvalidKey      = errors.New("key does not exist")
	ErrLockNotAcquired = errors.New("unable to acquire lock")
)

type RedisClient struct {
	client *redis.Client
	red    *redsync.Redsync
}

func NewRedisClient(client *redis.Client, red *redsync.Redsync) RedisClient {
	return RedisClient{
		client: client,
		red:    red,
	}
}

func (rc *RedisClient) Get(key string) (string, error) {
	val, err := rc.client.Get(context.Background(), key).Result()
	if err != nil {
		if err == redis.Nil {
			return "", ErrInvalidKey
		}
		return "", err
	}
	return val, nil
}

func (rc *RedisClient) Set(key string, value string) error {
	mutex := rc.red.NewMutex(key)

	if err := mutex.Lock(); err != nil {
		return ErrLockNotAcquired
	}

	defer mutex.Unlock()

	err := rc.client.Set(context.Background(), key, value, time.Second*time.Duration(3600)).Err()
	if err != nil {
		return err
	}
	return nil
}
*/
