package runtime

import (
	"go-admin/common/core/storage"
)

const (
	IntervalTenant = ":"
	appPrefix      = "app"
)

// NewCache 创建对应上下文缓存
func NewCache(store storage.AdapterCache) storage.AdapterCache {
	return &Cache{
		store: store,
	}
}

type Cache struct {
	store           storage.AdapterCache
	wxTokenStoreKey string
}

func (e *Cache) Exist(prefix, key string) bool {
	prefix = appPrefix + IntervalTenant + prefix
	v, err := e.store.Get(prefix, key)
	if err != nil || v == "" {
		return false
	}
	return true
}

// String string输出
func (e *Cache) String() string {
	if e.store == nil {
		return ""
	}
	return e.store.String()
}

// Connect 初始化
func (e Cache) Connect() error {
	return nil
	//return e.store.Connect()
}

// Get val in cache
func (e Cache) Get(prefix, key string) (string, error) {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.Get(prefix, key)
}

// Set val in cache
func (e Cache) Set(prefix, key string, val interface{}, expire int) error {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.Set(prefix, key, val, expire)
}

// Del delete key in cache
func (e Cache) Del(prefix, key string) error {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.Del(prefix, key)
}
func (e *Cache) HashSet(expire int, prefix, key string, values map[string]interface{}) error {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.HashSet(expire, prefix, key, values)
}

func (e *Cache) HashGetAll(prefix, key string) (map[string]string, error) {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.HashGetAll(prefix, key)
}

// HashGet get val in hashtable cache
func (e Cache) HashGet(prefix, key, field string) (string, error) {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.HashGet(prefix, key, field)
}

// HashDel delete one key:value pair in hashtable cache
func (e Cache) HashDel(prefix, key, field string) error {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.HashDel(prefix, key, field)
}

// Increase value
func (e Cache) Increase(prefix, key string) error {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.Increase(prefix, key)
}

func (e Cache) Decrease(prefix, key string) error {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.Decrease(prefix, key)
}

//
//  Expire
//  @Description:
//  @receiver e
//  @param key
//  @param expire 有效时间  秒
//  @return error
//
func (e Cache) Expire(prefix, key string, expire int) error {
	prefix = appPrefix + IntervalTenant + prefix
	return e.store.Expire(prefix, key, expire)
}
