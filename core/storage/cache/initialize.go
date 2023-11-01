/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package cache

import (
	"fmt"
	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/captchautils"
	"go-admin/core/utils/log"
)

// Setup 配置storage组件
func Setup() {
	//4. 设置缓存
	cacheAdapter, err := config.CacheConfig.Setup()
	if err != nil {
		panic(fmt.Sprintf("cache setup error, %s\n", err.Error()))
	}

	runtime.RuntimeConfig.SetCacheAdapter(cacheAdapter)
	//5. 设置验证码缓存
	captchautils.SetStore(captchautils.NewCacheStore(cacheAdapter, config.CacheConfig.Expired))

	//6. 设置队列
	if !config.QueueConfig.Empty() {
		if q := runtime.RuntimeConfig.GetQueueAdapter(); q != nil {
			q.Shutdown()
		}
		queueAdapter, err := config.QueueConfig.Setup()
		if err != nil {
			log.Fatalf("queue setup error, %s\n", err.Error())
		}
		runtime.RuntimeConfig.SetQueueAdapter(queueAdapter)
		defer func() {
			go queueAdapter.Run()
		}()
	}

	//7. 设置分布式锁
	if !config.LockerConfig.Empty() {
		lockerAdapter, err := config.LockerConfig.Setup()
		if err != nil {
			log.Fatalf("locker setup error, %s\n", err.Error())
		}
		runtime.RuntimeConfig.SetLockerAdapter(lockerAdapter)
	}
}
