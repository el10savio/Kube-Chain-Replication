package redisinterface

import (
	"errors"

	"github.com/gomodule/redigo/redis"
)

const (
	// PORT ...
	PORT = "6379"
)

var (
	// ErrNil ...
	ErrNil = redis.ErrNil
)

// NewPool ...
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial:      Dial,
	}
}

// Dial ...
func Dial() (redis.Conn, error) {
	connection, err := redis.Dial("tcp", ":"+PORT)
	if err != nil {
		panic(err.Error())
	}

	return connection, err
}

// Ping ...
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

// Set ...
func Set(connection redis.Conn, key string, value string) error {
	_, err := connection.Do("SET", key, value)
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func Get(connection redis.Conn, key string) (string, error) {
	value, err := redis.String(connection.Do("GET", key))
	if err != nil {
		return "", err
	}

	return value, nil
}
