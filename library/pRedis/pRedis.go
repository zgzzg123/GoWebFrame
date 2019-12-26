package pRedis

import (
	"github.com/garyburd/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
	"fmt"
	"strconv"
)

var (
	rsConfig redisConfigStruct
	Pool     *redis.Pool
)

type redisConfigStruct struct {
	host     string
	port     int
	password string
}

func init() {
	redisHost := os.Getenv("REDIS_HOST")

	rsConfig = redisConfigStruct{
		host:     "127.0.0.1",
		port:     6379,
		password: "cae0f7fcf1",
	}

	fmt.Println(redisHost)
	if redisHost == "" {
		redisHost = ":6379"
	}
	Pool = newPool(rsConfig)
	cleanupHook()
}

func newPool(server redisConfigStruct) *redis.Pool {

	return &redis.Pool{

		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server.host + ":" + strconv.Itoa(server.port))
			if err != nil {
				return nil, err
			}

			if server.password != "" {
				if _, err := c.Do("AUTH", server.password); err != nil {
					c.Close()
					return nil, err
				}
			}

			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func cleanupHook() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
