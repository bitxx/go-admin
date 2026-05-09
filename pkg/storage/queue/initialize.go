/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 3:39 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 3:39 下午
 */

package queue

import (
	"fmt"
	"go-admin/pkg/config"
	"go-admin/pkg/runtime"
)

// Setup 配置队列组件
func Setup() {
	if !config.QueueConfig.Empty() {
		if q := runtime.RuntimeConfig.GetQueueAdapter(); q != nil {
			q.Shutdown()
		}
		queueAdapter, err := config.QueueConfig.Setup()
		if err != nil {
			panic(fmt.Sprintf("queue setup error, %s\n", err.Error()))
		}
		runtime.RuntimeConfig.SetQueueAdapter(queueAdapter)
		defer func() {
			go queueAdapter.Run()
		}()
	}
}
