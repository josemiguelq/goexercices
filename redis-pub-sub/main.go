package main

import (
	"fmt" 
	"flag" 
    "github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"os" 
	"redis-lab/redisb"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("cade o arquivo?")
	}

	var msg string
	// var key string

	flag.StringVar(&msg, "msg", "", "empty msg")
	// flag.StringVar(&key, "key", "", "empty msg")

	flag.Parse()
	rdb := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_HOST") +":"+ os.Getenv("REDIS_PORT"),
        Password: "", 
        DB:       0, 
    })
	saved := redisb.Save(redisb.ToUpper(msg), rdb)
	fmt.Println(saved)
}
