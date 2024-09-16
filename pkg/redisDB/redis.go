package redisdb

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

// singleton
type singletonRedis struct {
	client *redis.Client
}

var (
	once          sync.Once
	instanceRedis *singletonRedis
)

func GetInstanceRedis() *singletonRedis {
	once.Do(func() {
		instanceRedis = &singletonRedis{}
		instanceRedis.init()
	})
	return instanceRedis
}

func (s *singletonRedis) init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	addr := os.Getenv("REDIS_ADDR")
	pwd := os.Getenv("REDIS_PWD")
	s.client = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pwd,
		DB:       0,
	})

	ping, err := s.client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("ping: %v\n", ping)
}

func (s *singletonRedis) SaveRefreshToken(refreshToken string) error {
	err := s.client.Set(context.Background(), "refreshToken", refreshToken, (7*24*60*60)*time.Second).Err()
	if err != nil {
		return fmt.Errorf("failed to save refresh token: %w", err) // Trả về lỗi
	}
	return nil
}

func (s *singletonRedis) CheckRefreshToken() (bool, error) {
	val, err := s.client.Get(context.Background(), "refreshToken").Result()
	if err == redis.Nil {
		// Token không tồn tại
		return false, nil
	} else if err != nil {
		return false, err
	}
	// Token tồn tại
	fmt.Printf("Token hiện tại: %s\n", val)
	return true, nil
}
