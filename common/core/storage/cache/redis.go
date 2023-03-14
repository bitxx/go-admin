package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go-admin/common/core/runtime"
	"time"
)

// NewRedis redis模式
func NewRedis(client *redis.Client, options *redis.Options) (*Redis, error) {
	if client == nil {
		client = redis.NewClient(options)
	}
	r := &Redis{
		client: client,
		ctx:    context.Background(),
	}
	err := r.connect()
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Redis cache implement
type Redis struct {
	client *redis.Client
	ctx    context.Context
}

func (*Redis) String() string {
	return "redis"
}

// connect connect test
func (r *Redis) connect() error {
	var err error
	_, err = r.client.Ping(r.ctx).Result()
	return err
}

func (r *Redis) Exist(prefix, key string) bool {
	key = prefix + runtime.IntervalTenant + key
	v, _ := r.client.Exists(r.ctx, key).Result()
	if v != 1 {
		return false
	}
	return true
}

// Get from key
func (r *Redis) Get(prefix, key string) (string, error) {
	key = prefix + runtime.IntervalTenant + key
	return r.client.Get(r.ctx, key).Result()
}

// Set value with key and expire time
func (r *Redis) Set(prefix, key string, val interface{}, expire int) error {
	key = prefix + runtime.IntervalTenant + key
	return r.client.Set(r.ctx, key, val, time.Duration(expire)*time.Second).Err()
}

// Del delete key in redis
func (r *Redis) Del(prefix, key string) error {
	key = prefix + runtime.IntervalTenant + key
	return r.client.Del(r.ctx, key).Err()
}

func (r *Redis) HashSet(expire int, prefix, key string, values map[string]interface{}) error {
	key = prefix + runtime.IntervalTenant + key
	err := r.client.HSet(r.ctx, key, values).Err()
	if err != nil {
		return err
	}
	return r.client.Expire(r.ctx, key, time.Duration(expire)*time.Second).Err()
}

// HashGet from key
func (r *Redis) HashGet(prefix, key, field string) (string, error) {
	key = prefix + runtime.IntervalTenant + key
	return r.client.HGet(r.ctx, key, field).Result()
}

func (r *Redis) HashGetAll(prefix, key string) (map[string]string, error) {
	key = prefix + runtime.IntervalTenant + key
	return r.client.HGetAll(r.ctx, key).Result()
}

// HashDel delete key in specify redis's hashtable
func (r *Redis) HashDel(prefix, key string, field string) error {
	key = prefix + runtime.IntervalTenant + key
	return r.client.HDel(r.ctx, key, field).Err()
}

// Increase
func (r *Redis) Increase(prefix, key string) error {
	key = prefix + runtime.IntervalTenant + key
	return r.client.Incr(r.ctx, key).Err()
}

func (r *Redis) Decrease(prefix, key string) error {
	key = prefix + runtime.IntervalTenant + key
	return r.client.Decr(r.ctx, key).Err()
}

// Set ttl
func (r *Redis) Expire(prefix, key string, expire int) error {
	key = prefix + runtime.IntervalTenant + key
	return r.client.Expire(r.ctx, key, time.Duration(expire)*time.Second).Err()
}

// GetClient 暴露原生client
func (r *Redis) GetClient() *redis.Client {
	return r.client
}
