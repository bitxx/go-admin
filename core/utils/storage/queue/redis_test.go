package queue

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-admin/core/utils/storage"
	redisqueue2 "go-admin/core/utils/storage/queue/redisqueue"
	"testing"
	"time"
)

func TestRedis_Append(t *testing.T) {
	type fields struct {
		ConnectOption   *redis.Options
		ConsumerOptions *redisqueue2.ConsumerOptions
		ProducerOptions *redisqueue2.ProducerOptions
		client          *redis.Client
		consumer        *redisqueue2.Consumer
		producer        *redisqueue2.Producer
	}
	type args struct {
		name    string
		message storage.Messager
	}
	client := redis.NewClient(&redis.Options{})
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"test01",
			fields{
				ConnectOption: &redis.Options{},
				ConsumerOptions: &redisqueue2.ConsumerOptions{
					VisibilityTimeout: 60 * time.Second,
					BlockingTimeout:   5 * time.Second,
					ReclaimInterval:   1 * time.Second,
					BufferSize:        100,
					Concurrency:       10,
					RedisClient:       client,
				},
				ProducerOptions: &redisqueue2.ProducerOptions{
					StreamMaxLength:      100,
					ApproximateMaxLength: false,
					RedisClient:          client,
				},
			},
			args{
				name: "test",
				message: &Message{redisqueue2.Message{
					Stream: "test",
					Values: map[string]interface{}{
						"key": "value",
					},
				}},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r, err := NewRedis(tt.fields.ProducerOptions, tt.fields.ConsumerOptions); err != nil {
				t.Errorf("SetQueue() error = %v", err)
			} else {
				if err := r.Append(tt.args.message); (err != nil) != tt.wantErr {
					t.Errorf("SetQueue() error = %v, wantErr %v", err, tt.wantErr)
				}
			}
		})
	}
}

func TestRedis_Register(t *testing.T) {
	type fields struct {
		ConnectOption   *redis.Options
		ConsumerOptions *redisqueue2.ConsumerOptions
		ProducerOptions *redisqueue2.ProducerOptions
		client          *redis.Client
		consumer        *redisqueue2.Consumer
		producer        *redisqueue2.Producer
	}
	type args struct {
		name string
		f    storage.ConsumerFunc
	}
	client := redis.NewClient(&redis.Options{})
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			"test01",
			fields{
				ConnectOption: &redis.Options{},
				ConsumerOptions: &redisqueue2.ConsumerOptions{
					VisibilityTimeout: 60 * time.Second,
					BlockingTimeout:   5 * time.Second,
					ReclaimInterval:   1 * time.Second,
					BufferSize:        100,
					Concurrency:       10,
					RedisClient:       client,
				},
				ProducerOptions: &redisqueue2.ProducerOptions{
					StreamMaxLength:      100,
					ApproximateMaxLength: true,
					RedisClient:          client,
				},
			},
			args{
				name: "login_log_queue",
				f: func(message storage.Messager) error {
					fmt.Println("ok")
					fmt.Println(message.GetValues())
					return nil
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if r, err := NewRedis(tt.fields.ProducerOptions, tt.fields.ConsumerOptions); err != nil {
				t.Errorf("SetQueue() error = %v", err)
			} else {
				r.Register(tt.args.name, tt.args.f)
				r.Run()
			}
		})
	}
}
