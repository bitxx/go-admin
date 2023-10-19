package queue

import (
	"github.com/redis/go-redis/v9"
	"go-admin/common/utils/storage"
	redisqueue2 "go-admin/common/utils/storage/queue/redisqueue"
)

// NewRedis redis模式
func NewRedis(
	producerOptions *redisqueue2.ProducerOptions,
	consumerOptions *redisqueue2.ConsumerOptions,
) (*Redis, error) {
	var err error
	r := &Redis{}
	r.producer, err = r.newProducer(producerOptions)
	if err != nil {
		return nil, err
	}
	r.consumer, err = r.newConsumer(consumerOptions)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Redis cache implement
type Redis struct {
	client   *redis.Client
	consumer *redisqueue2.Consumer
	producer *redisqueue2.Producer
}

func (Redis) String() string {
	return "redis"
}

func (r *Redis) newConsumer(options *redisqueue2.ConsumerOptions) (*redisqueue2.Consumer, error) {
	if options == nil {
		options = &redisqueue2.ConsumerOptions{}
	}
	return redisqueue2.NewConsumerWithOptions(options)
}

func (r *Redis) newProducer(options *redisqueue2.ProducerOptions) (*redisqueue2.Producer, error) {
	if options == nil {
		options = &redisqueue2.ProducerOptions{}
	}
	return redisqueue2.NewProducerWithOptions(options)
}

func (r *Redis) Append(message storage.Messager) error {
	err := r.producer.Enqueue(&redisqueue2.Message{
		ID:     message.GetID(),
		Stream: message.GetStream(),
		Values: message.GetValues(),
	})
	return err
}

func (r *Redis) Register(name string, f storage.ConsumerFunc) {
	r.consumer.Register(name, func(message *redisqueue2.Message) error {
		m := new(Message)
		m.SetValues(message.Values)
		m.SetStream(message.Stream)
		m.SetID(message.ID)
		return f(m)
	})
}

func (r *Redis) Run() {
	r.consumer.Run()
}

func (r *Redis) Shutdown() {
	r.consumer.Shutdown()
}
