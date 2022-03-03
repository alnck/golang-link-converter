package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

var lock sync.Mutex

type RedisRepository struct {
	client *redis.Client
	logger *zap.Logger
}

var RedisRepositoryInstance *RedisRepository

func NewRedisRepository(logger *zap.Logger) *RedisRepository {

	if RedisRepositoryInstance == nil {

		lock.Lock()
		defer lock.Unlock()

		if RedisRepositoryInstance == nil {
			client := redis.NewClient(&redis.Options{
				Addr:     os.Getenv("CONNECTION_STRING"),
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       0,
			})
			RedisRepositoryInstance = &RedisRepository{client: client, logger: logger}
		} else {
			return RedisRepositoryInstance

		}
	}
	fmt.Println("zaten var")
	return RedisRepositoryInstance
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
		return err
	}

	err = json.Unmarshal([]byte(val), &src)
	if err != nil {
		return err
	}

	return nil
}
