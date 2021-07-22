package main

import (
	"redis-lab/worker"
	"github.com/joho/godotenv"
	"time"
	"fmt"

)

func main() {
	godotenv.Load()
	for {
		worker.Changed()
		time.Sleep(time.Millisecond * 500)
	}
}