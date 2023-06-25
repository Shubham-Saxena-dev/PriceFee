package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"sync"
	"time"
)

var (
	ErrInvalidKey      = errors.New("key does not exist")
	ErrLockNotAcquired = errors.New("unable to acquire lock")
)

type RedisClient struct {
	client      *redis.Client
	red         *redsync.Redsync
	mutexes     map[string]*sync.Mutex
	mutexesLock sync.Mutex
}

func NewRedisClient(client *redis.Client, red *redsync.Redsync) RedisClient {
	return RedisClient{
		client:  client,
		red:     red,
		mutexes: make(map[string]*sync.Mutex),
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
	mutex := rc.getOrCreateMutex(key)
	mutex.Lock()
	defer mutex.Unlock()

	err := rc.client.Set(context.Background(), key, value, time.Second*time.Duration(3600)).Err()
	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisClient) getOrCreateMutex(key string) *sync.Mutex {
	rc.mutexesLock.Lock()
	defer rc.mutexesLock.Unlock()

	mutex, ok := rc.mutexes[key]
	if !ok {
		mutex = &sync.Mutex{}
		rc.mutexes[key] = mutex
	}

	return mutex
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	pool := goredis.NewPool(client)
	rs := redsync.New(pool)

	cache := NewRedisClient(client, rs)

	cache.Set("offset", "10")
	cache.Set("offset2", "20")
	fmt.Println(cache.Get("offset"))
	fmt.Println(cache.Get("offset2"))
	cache.Set("offset2", "30")
	cache.Set("offset3", "40")
	fmt.Println(cache.Get("offset2"))
	fmt.Println(cache.Get("offset3"))
}
