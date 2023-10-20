package storage

import (
	"github.com/bsm/redislock"
)

const (
	PrefixKey = "__host"
)

type AdapterCache interface {
	String() string
	Get(prefix, key string) (string, error)
	Set(prefix, key string, val interface{}, expire int) error
	Del(prefix, key string) error
	HashSet(expire int, prefix, key string, values map[string]interface{}) error
	HashGet(prefix, key, field string) (string, error)
	HashGetAll(prefix, key string) (map[string]string, error)
	HashDel(prefix, key, field string) error
	Increase(prefix, key string) error
	Decrease(prefix, key string) error
	Expire(prefix, key string, dur int) error
	Exist(prefix, key string) bool
}

type AdapterQueue interface {
	String() string
	Append(message Messager) error
	Register(name string, f ConsumerFunc)
	Run()
	Shutdown()
}

type Messager interface {
	SetID(string)
	SetStream(string)
	SetValues(map[string]interface{})
	GetID() string
	GetStream() string
	GetValues() map[string]interface{}
	GetPrefix() string
	SetPrefix(string)
}

type ConsumerFunc func(Messager) error

type AdapterLocker interface {
	String() string
	Lock(key string, ttl int64, options *redislock.Options) (*redislock.Lock, error)
}
