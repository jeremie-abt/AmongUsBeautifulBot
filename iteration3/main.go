package main

import "os"

import "fmt"
import "github.com/go-redis/redis"

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6700", // use default Addr
		Password: "",               // no password set
		DB:       0,                // use default DB
	})

	pong, err := rdb.Ping().Result()
	fmt.Printf("rdb : %+v\n", rdb)
	fmt.Printf("pong : %+v\n", pong)
	fmt.Printf("err : %+v\n", err)
	fmt.Printf("Os test : %s\n", os.Getenv("test"))
}
