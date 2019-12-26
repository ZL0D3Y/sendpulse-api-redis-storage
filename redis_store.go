package redisstore

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type (
	Options struct {
		MaxAge      int
		SessionName string
	}
	RedisStore struct {
		client  redis.UniversalClient
		options *Options
	}
)

func (s *RedisStore) NewStore(host string, port string, options *Options) (*RedisStore, error) {
	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
	if _, err := client.Ping().Result(); err != nil {
		return nil, err
	}

	return &RedisStore{client: client, options: options}, nil
}

// save writes session in Redis
func (s *RedisStore) Save(session string) error {

	return s.client.Set("sp"+s.options.SessionName, session, time.Duration(s.options.MaxAge)*time.Second).Err()
}
