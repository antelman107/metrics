package nosql

import (
	"fmt"
	"github.com/go-redis/redis"
)

type PubSubSubscriberRedis struct {
	channel     string
	redisClient *redis.Client
	pubSub      *redis.PubSub
}

func (s *PubSubSubscriberRedis) Receive() ([]byte, error) {
	if s.pubSub == nil {
		return nil, fmt.Errorf("Before receive quotes you must subscribe of them.")
	}

	message, err := s.pubSub.ReceiveMessage()
	if err != nil {
		return nil, err
	}

	return []byte(message.Payload), nil
}

func (s *PubSubSubscriberRedis) Subscribe(key string) error {
	s.channel = key
	s.pubSub = s.redisClient.Subscribe(s.channel)

	return nil
}

func (s *PubSubSubscriberRedis) Unsubscribe() error {
	if err := s.pubSub.Unsubscribe(s.channel); err != nil {
		return err
	}

	res := s.pubSub.Close()
	if res != nil {
		return res
	}
	s.pubSub = nil
	return nil
}

func NewPubSubSubscriberRedis(redisClient *redis.Client) *PubSubSubscriberRedis {
	return &PubSubSubscriberRedis{
		redisClient: redisClient,
	}
}
