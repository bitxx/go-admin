/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package cache

import (
	"fmt"
	"go-admin/pkg/config"
	"go-admin/pkg/runtime"
	"go-admin/pkg/utils/captchautils"
)

// Setup 配置缓存组件
func Setup() {
	cacheAdapter, err := config.CacheConfig.Setup()
	if err != nil {
		panic(fmt.Sprintf("cache setup error, %s\n", err.Error()))
	}

	runtime.RuntimeConfig.SetCacheAdapter(cacheAdapter)

	//设置验证码缓存
	captchautils.SetStore(captchautils.NewCacheStore(cacheAdapter, config.CacheConfig.Expired))
}
