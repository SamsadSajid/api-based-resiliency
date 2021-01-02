package config

import (
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"log"
	"os"
)

func InitRedisPool() gin.HandlerFunc{
	var pool *redis.Pool

	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6389")
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
	log.Println("init redis pool done")

	return func(c *gin.Context) {
		c.Set("redis-pool", pool)
	}
}
