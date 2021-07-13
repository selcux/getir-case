package inmemory

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"os"
)

type MemStore interface {
	init() error
	Set(key, value string) error
	Get(key string) (string, error)
}

type storage struct {
	client *redis.Client
}

func NewStorage() (MemStore, error) {
	s := &storage{}
	err := s.init()
	if err != nil {
		return nil, err
	}

	return s, nil
}

func (s *storage) init() error {
	client := redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_URL"),
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return errors.Wrap(err, "unable to connect to Redis")
	}

	s.client = client

	return nil
}

func (s *storage) Set(key, value string) error {
	err := s.client.Set(context.Background(), key, value, 0).Err()

	return errors.Wrapf(err, "could not set (%s: %s)", key, value)
}

func (s *storage) Get(key string) (string, error) {
	value, err := s.client.Get(context.Background(), key).Result()
	if err != nil {
		return "", errors.Wrapf(err, "could not get the value of the key: %s", key)
	}

	return value, nil
}
