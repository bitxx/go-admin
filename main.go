package main

import (
	"go-admin/common/cmd"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Go-admin 后台管理系统
// @version 2.0.0
// @in header
// @name Authorization
func main() {
	cmd.Execute()
}
