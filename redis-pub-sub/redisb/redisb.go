package redisb

import (
	"fmt" 
	"context"
    "github.com/go-redis/redis/v8"
	"strings" 
	"os" 
)


var ctx = context.Background()

func Save(msg string, rdb *redis.Client) string{	
	err := rdb.Set(ctx, "key", msg, 0).Err()
	fmt.Println(err)
	val, _ := rdb.Get(ctx, "key").Result()
	return val
}

func Get(rdb *redis.Client) string{	
	val, _ := rdb.Get(ctx, "key").Result()
	return val
}

func ToUpper(msg string) string{
	return strings.ToUpper(msg)
}

func CreateClient() *redis.Client{
	return redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_HOST") +":"+ os.Getenv("REDIS_PORT"),
        Password: "", 
        DB:       0, 
    })
}