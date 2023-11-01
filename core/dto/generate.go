package dto

import (
	vd "github.com/bytedance/go-tagexpr/v2/validator"
	"go-admin/core/utils/log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ObjectById struct {
	Id  int64   `uri:"id" json:"-"`
	Ids []int64 `json:"ids"`
}

func (s *ObjectById) Bind(ctx *gin.Context) error {
	var err error
	rLog := log.GetRequestLogger(ctx)
	err = ctx.ShouldBindUri(s)
	if err != nil {
		rLog.Warnf("ShouldBindUri error: %s", err.Error())
		return err
	}
	if ctx.Request.Method == http.MethodDelete {
		err = ctx.ShouldBind(&s)
		if err != nil {
			rLog.Warnf("ShouldBind error: %s", err.Error())
			return err
		}
		if len(s.Ids) > 0 {
			return nil
		}
		if s.Ids == nil {
			s.Ids = make([]int64, 0)
		}
		if s.Id != 0 {
			s.Ids = append(s.Ids, s.Id)
		}
	}
	if err = vd.Validate(s); err != nil {
		rLog.Errorf("Validate error: %s", err.Error())
		return err
	}
	return err
}

func (s *ObjectById) GetId() interface{} {
	if len(s.Ids) > 0 {
		s.Ids = append(s.Ids, s.Id)
		return s.Ids
	}
	return s.Id
}

type ObjectGetReq struct {
	Id int64 `uri:"id" json:"-"`
}

func (s *ObjectGetReq) Bind(ctx *gin.Context) error {
	var err error
	rLog := log.GetRequestLogger(ctx)
	err = ctx.ShouldBindUri(s)
	if err != nil {
		rLog.Warnf("ShouldBindUri error: %s", err.Error())
		return err
	}
	if err = vd.Validate(s); err != nil {
		rLog.Errorf("Validate error: %s", err.Error())
		return err
	}
	return err
}

func (s *ObjectGetReq) GetId() interface{} {
	return s.Id
}

type ObjectDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ObjectDeleteReq) Bind(ctx *gin.Context) error {
	var err error
	rLog := log.GetRequestLogger(ctx)
	err = ctx.ShouldBind(&s)
	if err != nil {
		rLog.Warnf("ShouldBind error: %s", err.Error())
		return err
	}
	if len(s.Ids) > 0 {
		return nil
	}
	if s.Ids == nil {
		s.Ids = make([]int, 0)
	}

	if err = vd.Validate(s); err != nil {
		rLog.Errorf("Validate error: %s", err.Error())
		return err
	}
	return err
}

func (s *ObjectDeleteReq) GetId() interface{} {
	return s.Ids
}
