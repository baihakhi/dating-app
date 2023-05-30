package driver

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis"
)

type RedisClient interface {
	HGet(key, field string) ([]byte, error)
	HSet(key, field string, value interface{}, expDur time.Duration) error
	HDel(key, field string) error
}

type redisCtx struct {
	redisClient redis.Cmdable
}

func NewRedis(redisClient redis.Cmdable) RedisClient {
	return &redisCtx{
		redisClient: redisClient,
	}
}

func (c *redisCtx) HGet(key, field string) ([]byte, error) {
	data, err := c.redisClient.HGet(key, field).Result()
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return []byte(data), nil
}

// HSet set
func (c *redisCtx) HSet(key, field string, value interface{}, expDur time.Duration) error {
	payload, err := json.Marshal(value)
	if err != nil {
		log.Print(err)
		return err
	}

	err = c.redisClient.HSet(key, field, payload).Err()
	if err != nil {
		log.Print(err)
		return err
	}

	if expDur > 0 {
		err = c.redisClient.Expire(key, expDur).Err()
		if err != nil {
			log.Print(err)
			return err
		}
	}

	return nil
}

func (c *redisCtx) HDel(key, field string) error {
	err := c.redisClient.HDel(key, field).Err()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

func NewRedisConnection(
	useSentinel bool,
	master string,
	hosts []string,
	port int,
	password string,
	database int,
	poolSize int,
	timeout time.Duration,
	minIdleConns int) *redis.Client {

	if useSentinel {
		if master == "" {
			master = "master"
		}
		return newRedisFailOverConnection(
			master,
			hosts,
			port,
			password,
			database,
			poolSize,
			timeout,
			minIdleConns)
	} else {
		return newRedisDirectConnection(
			hosts[0],
			port,
			password,
			database,
			poolSize,
			timeout,
			minIdleConns)
	}
}

func newRedisDirectConnection(
	host string,
	port int,
	password string,
	database int,
	poolSize int,
	timeout time.Duration,
	minIdleConns int) *redis.Client {

	redisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		Password:     password,
		DB:           database,
		PoolSize:     poolSize,
		PoolTimeout:  timeout,
		MinIdleConns: minIdleConns,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	return redisClient
}

func newRedisFailOverConnection(
	master string,
	hosts []string,
	port int,
	password string,
	database int,
	poolSize int,
	timeout time.Duration,
	minIdleConns int) *redis.Client {

	var hostPorts []string
	for _, v := range hosts {
		hostPorts = append(hostPorts, fmt.Sprintf("%s:%d", v, port))
	}

	redisClient := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    master,
		SentinelAddrs: hostPorts,
		Password:      password,
		DB:            database,
		PoolSize:      poolSize,
		PoolTimeout:   timeout,
		MinIdleConns:  minIdleConns,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}

	return redisClient
}
