package config

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // nếu dùng Docker hoặc Redis mặc định
		Password: "",               // để trống nếu không set password
		DB:       0,                // DB số 0
	})

	// Kiểm tra kết nối
	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		fmt.Println("❌ Redis không kết nối được:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Redis đã kết nối!")
}
