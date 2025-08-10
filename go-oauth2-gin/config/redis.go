package config

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client
var Ctx = context.Background()

func ConnectRedis() {
	port, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		log.Fatalf("Invalid REDIS_PORT in .env: %v", err)
	}

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "redis:6379", // Use "localhost:6379" if testing without Docker
		Password: "",           // No password set
		DB:       0,            // Default DB
	})

	// Ping the Redis server to check the connection
	status, err := RedisClient.Ping(Ctx).Result()
    if err != nil {
        log.Fatalf("❌ Redis connection failed: %v, %v", err, status)
    }

	log.Printf("✅ Connected to Redis: %v, %v", port, status)
}
