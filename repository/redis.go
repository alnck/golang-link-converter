package repository

import (
	"encoding/json"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

type RedisRepository struct {
	client *redis.Client
	logger *zap.Logger
}

func NewRedisRepository(logger *zap.Logger) *RedisRepository {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("CONNECTION_STRING"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	return &RedisRepository{client: client, logger: logger}
}

func (repository *RedisRepository) SetKey(key string, value interface{}, ttl int) {
	byteData, err := json.Marshal(value)

	if err != nil {
		repository.logger.Error("SetKey - error:", zap.Error(err))
		return
	}

	duration, _ := time.ParseDuration(strconv.FormatInt(int64(ttl), 10))
	status := repository.client.Set(key, byteData, duration)
	_, err = status.Result()
	if err != nil {
		repository.logger.Error("SetKey - error:", zap.Error(err))
	}
}

func (repository *RedisRepository) GetKey(key string, src interface{}) error {
	val, err := repository.client.Get(key).Result()
	if err != nil {
		//repository.logger.Error("SetKey - error:", zap.Error(err))
		return err
	}

	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}

	return nil
}
