package redisinterface

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

const (
	// PORT defines the port value 
	// for the Redis service
	PORT = "6379"
)

var (
	// ErrNil defines the error 
	// from Redis that the
	// response is nil
	ErrNil = redis.ErrNil
)

// NewPool allocates a new Redis pool
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial:      Dial,
	}
}

// Dial connects to the 
// given Redis service
func Dial() (redis.Conn, error) {
	connection, err := redis.Dial("tcp", ":"+PORT)
	if err != nil {
		panic(err.Error())
	}

	return connection, err
}

// Ping checks if the connection to
// the Redis service is alright
func Ping(connection redis.Conn) error {
	pong, err := redis.String(connection.Do("PING"))
	if err != nil {
		return err
	}

	if pong != "PONG" {
		return errors.New("invalid response from ping")
	}

	return nil
}

// Set sets a value in Redis 
// given the key and value
func Set(connection redis.Conn, key string, value string) error {
	_, err := connection.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// Get gets a value in Redis 
// given the key
func Get(connection redis.Conn, key string) (string, error) {
	value, err := redis.String(connection.Do("GET", key))
	if err != nil {
		return "", err
	}

	return value, nil
}
