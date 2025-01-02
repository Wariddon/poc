package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	ctx := context.Background()

	// Connect to Redis database index 7
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   7, // Use database index 7
	})

	// Ensure the test key exists in Redis
	err := rdb.Set(ctx, "test_key", "Hello, Redis!", 0).Err()
	if err != nil {
		fmt.Printf("Error setting key: %v\n", err)
		return
	}

	N := 1000
	start := time.Now()

	for i := 0; i < N; i++ {
		_, err := rdb.Get(ctx, "test_key").Result()
		if err != nil {
			fmt.Printf("Error getting key: %v\n", err)
			return
		}
	}

	fmt.Printf("Go Redis Total Time for %d fetches from DB 7: %v seconds\n", N, time.Since(start).Seconds())
}
