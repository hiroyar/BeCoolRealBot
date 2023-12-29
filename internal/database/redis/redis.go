package redis

import (
	"BeCoolRealBot/internal/helpers"
	"github.com/go-redis/redis"
	"log"
	"os"
)

type Instance struct {
	Db *redis.Client
}

var Cache Instance

func Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       helpers.FromStringToInt(os.Getenv("REDIS_DB")),
	})

	log.Println("Connected to redis")

	Cache = Instance{
		Db: client,
	}
}
