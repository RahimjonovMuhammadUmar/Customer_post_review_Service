package redis

import (
	"context"
	"time"

	"exam/api_gateway/storage/repo"

	"github.com/gomodule/redigo/redis"
)

type redisRepo struct {
	rds *redis.Pool
}

func NewRedisRepo(rds *redis.Pool) repo.InMemoryStorageI {
	return &redisRepo{
		rds: rds,
	}
}
func (r *redisRepo) Set(key, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(3))
	defer cancel()
	conn := r.rds.Get()
	_, err := conn.Do("SET", ctx, key, value, 0)
	if err != nil {
		return err
	}

	return nil
}

func (r *redisRepo) SetWithTTl(key, value string, duration int64) error {
	conn := r.rds.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, duration, value)
	return err
}

func (r *redisRepo) Get(key string) (interface{}, error) {
	conn := r.rds.Get()
	defer conn.Close()
	return conn.Do("GET", key)
}
