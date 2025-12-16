/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package locker

import (
	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/log"
)

// Setup 配置storage组件
func Setup() {
	// 设置分布式锁
	if !config.LockerConfig.Empty() {
		lockerAdapter, err := config.LockerConfig.Setup()
		if err != nil {
			log.Fatalf("locker setup error, %s\n", err.Error())
		}
		runtime.RuntimeConfig.SetLockerAdapter(lockerAdapter)
	}
}
