package service

import (
	"go-admin/common/core/service"
	"go-admin/common/global"
	"io/ioutil"
	"os"
)

type SysRuntimeConfig struct {
	service.Service
}

func NewSysRuntimeConfigService(s *service.Service) *SysRuntimeConfig {
	var srv = new(SysRuntimeConfig)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

func (e *SysRuntimeConfig) GetConfig() (string, error) {
	fh, err := os.Open(global.ConfigPath)
	if err != nil {
		return "", err
	}
	defer fh.Close()
	result, err := ioutil.ReadAll(fh)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

// SaveConfig
// @Description: 写入风险大，暂不提供
// @param data
// @return error
func SaveConfig(data string) error {
	return ioutil.WriteFile(global.ConfigPath, []byte(data), 0666)
}
