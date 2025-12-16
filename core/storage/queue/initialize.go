/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package queue

import (
	"go-admin/core/config"
	"go-admin/core/runtime"
	"go-admin/core/utils/log"
)

// Setup 配置队列组件
func Setup() {
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
}
