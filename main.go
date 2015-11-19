package main

import (
	"flag"
	"fmt"
	//"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
)

var (
	redisAddress   = flag.String("redis-address", ":6379", "Address to the Redis server")
	maxConnections = flag.Int("max-connections", 10, "Max connections to Redis")
)

func main() {
	flag.Parse()
	//redisPool := redis.NewPool(func()(redis.Conn, error) {
	//c, err := redis.Dial("tcp", *redisAddress)
	//
	//if err != nil {
	//return nil, err
	//}
	//return c, err
	//}, *maxConnections)
	//
	//defer redisPool.close()

	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		fmt.Println("err")
		return
	}
	conn.Do("SET", "Michael", "Cool")
	defer conn.Close()

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		s, _ := redis.String(conn.Do("GET", "Michael"))
		c.String(200, s)
	})
	r.Run(":8080")
}
