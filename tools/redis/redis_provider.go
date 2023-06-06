package redis

import	(
	"github.com/redis/go-redis/v9"
	"context"
)

func NewRedisClient() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:	  "localhost:6379",
        Password: "", // no password set
        DB:		  0,  // use default DB
    })
	return client
}

func Ping(client *redis.Client) error {
	return client.Ping(context.Background()).Err()
}