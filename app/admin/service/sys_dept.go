package service

import (
	sysLang "go-admin/app/admin/lang"
	"go-admin/app/admin/models"
	"go-admin/core/dto/service"
	"go-admin/core/global"
	"go-admin/core/lang"
	"go-admin/core/middleware"
	"gorm.io/gorm"
	"strconv"
	"time"

	"go-admin/app/admin/service/dto"
	cDto "go-admin/core/dto"
)

type SysDept struct {
	service.Service
}

func NewSysDeptService(s *service.Service) *SysDept {
	var srv = new(SysDept)
	srv.Orm = s.Orm
	srv.Log = s.Log
	return srv
}

// GetPage 获取SysDept列表
func (e *SysDept) GetPage(c *dto.SysDeptQueryReq, p *middleware.DataPermission) ([]models.SysDept, int64, int, error) {
	var list []models.SysDept
	var data models.SysDept
	var count int64

	err := e.Orm.Order("created_at desc").Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			middleware.Permission(data.TableName(), p),
		).Find(&list).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, count, lang.SuccessCode, nil
}

// Get 获取SysDept对象
func (e *SysDept) Get(id int64, p *middleware.DataPermission) (*models.SysDept, int, error) {
	if id <= 0 {
		return nil, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	data := &models.SysDept{}
	err := e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).First(data, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

// QueryOne 通过自定义条件获取一条记录
func (e *SysDept) QueryOne(queryCondition *dto.SysDeptQueryReq, p *middleware.DataPermission) (*models.SysDept, int, error) {
	data := &models.SysDept{}
	err := e.Orm.Model(&models.SysDept{}).
		Scopes(
			cDto.MakeCondition(queryCondition.GetNeedSearch()),
			middleware.Permission(data.TableName(), p),
		).First(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return nil, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return data, lang.SuccessCode, nil
}

// Count 获取条数
func (e *SysDept) Count(c *dto.SysDeptQueryReq) (int64, int, error) {
	var err error
	var count int64
	err = e.Orm.Model(&models.SysDept{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Limit(-1).Offset(-1).
		Count(&count).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return 0, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	if err == gorm.ErrRecordNotFound {
		return 0, lang.DataNotFoundCode, lang.MsgErr(lang.DataNotFoundCode, e.Lang)
	}
	return count, lang.SuccessCode, nil
}

// Insert 创建SysDept对象
func (e *SysDept) Insert(c *dto.SysDeptInsertReq) (int64, int, error) {
	if c.CurrUserId <= 0 {
		return 0, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.ParentId.IntPart() < 0 {
		return 0, sysLang.SysDeptParentIdEmptyCode, lang.MsgErr(sysLang.SysDeptParentIdEmptyCode, e.Lang)
	}
	if c.DeptName == "" {
		return 0, sysLang.SysDeptNameEmptyCode, lang.MsgErr(sysLang.SysDeptNameEmptyCode, e.Lang)
	}
	if c.Leader == "" {
		return 0, sysLang.SysDeptLeaderEmptyCode, lang.MsgErr(sysLang.SysDeptLeaderEmptyCode, e.Lang)
	}

	//确保部门名称不存在
	req := dto.SysDeptQueryReq{}
	req.DeptName = c.DeptName
	count, respCode, err := e.Count(&req)
	if err != nil && respCode != lang.DataNotFoundCode {
		return 0, respCode, err
	}
	if count > 0 {
		return 0, sysLang.SysDeptNameExistCode, lang.MsgErr(sysLang.SysDeptNameExistCode, e.Lang)
	}
	//path组装
	deptPath := "0,"
	if c.ParentId.IntPart() > 0 {
		dept := &models.SysDept{}
		dept, respCode, err = e.Get(c.ParentId.IntPart(), nil)
		if err != nil {
			return 0, respCode, err
		}
		deptPath = dept.DeptPath + "," + strconv.FormatInt(dept.Id, 10) + ","
	}

	now := time.Now()
	data := models.SysDept{}
	data.DeptName = c.DeptName
	data.ParentId = c.ParentId.IntPart()
	data.DeptPath = deptPath
	data.Sort = c.Sort
	data.Leader = c.Leader
	data.Phone = c.Phone
	data.Email = c.Email
	data.Status = global.SysStatusOk
	data.CreateBy = c.CurrUserId
	data.UpdateBy = c.CurrUserId
	data.CreatedAt = &now
	data.UpdatedAt = &now
	err = e.Orm.Create(&data).Error
	if err != nil {
		return 0, lang.DataInsertLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataInsertCode, lang.DataInsertLogCode, err)
	}
	return data.Id, lang.SuccessCode, nil
}

// Update 修改SysDept对象
func (e *SysDept) Update(c *dto.SysDeptUpdateReq, p *middleware.DataPermission) (bool, int, error) {
	if c.Id <= 0 || c.CurrUserId <= 0 {
		return false, lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}
	if c.DeptName == "" {
		return false, sysLang.SysDeptNameEmptyCode, lang.MsgErr(sysLang.SysDeptNameEmptyCode, e.Lang)
	}
	if c.Leader == "" {
		return false, sysLang.SysDeptLeaderEmptyCode, lang.MsgErr(sysLang.SysDeptLeaderEmptyCode, e.Lang)
	}
	data, respCode, err := e.Get(c.Id, p)
	if err != nil {
		return false, respCode, err
	}

	//最小化变更改动过的数据
	updates := map[string]interface{}{}
	if c.DeptName != "" && data.DeptName != c.DeptName {
		req := dto.SysDeptQueryReq{}
		req.DeptName = c.DeptName
		resp, respCode, err := e.QueryOne(&req, nil)
		if err != nil && respCode != lang.DataNotFoundCode {
			return false, respCode, err
		}
		if respCode == lang.SuccessCode && resp.Id != data.Id {
			return false, sysLang.SysDeptNameExistCode, lang.MsgErr(sysLang.SysDeptNameExistCode, e.Lang)
		}
		updates["dept_name"] = c.DeptName
	}
	if c.Sort > 0 && data.Sort != c.Sort {
		updates["sort"] = c.Sort
	}
	if c.Leader != "" && data.Leader != c.Leader {
		updates["leader"] = c.Leader
	}
	if c.Phone != "" && data.Phone != c.Phone {
		updates["phone"] = c.Phone
	}
	if c.Email != "" && data.Email != c.Email {
		updates["email"] = c.Email
	}

	if len(updates) > 0 {
		updates["updated_at"] = time.Now()
		updates["update_by"] = c.CurrUserId
		err = e.Orm.Model(&data).Where("id=?", data.Id).Updates(&updates).Error
		if err != nil {
			return false, lang.DataUpdateLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataUpdateCode, lang.DataUpdateLogCode, err)
		}
		return true, lang.SuccessCode, nil
	}
	return false, lang.SuccessCode, nil
}

// Remove 删除SysDept
func (e *SysDept) Remove(ids []int64, p *middleware.DataPermission) (int, error) {
	if len(ids) <= 0 {
		return lang.ParamErrCode, lang.MsgErr(lang.ParamErrCode, e.Lang)
	}

	for _, id := range ids {
		//若被使用，不得删除
		dataReq := dto.SysDeptQueryReq{}
		dataReq.DeptPath = "," + strconv.FormatInt(id, 10) + ","
		count, respCode, err := e.Count(&dataReq)
		if err != nil && respCode != lang.DataNotFoundCode {
			return respCode, err
		}
		if count > 0 {
			return sysLang.SysDeptChildExistNoDelCode, lang.MsgErr(sysLang.SysDeptChildExistNoDelCode, e.Lang)
		}
	}

	var err error
	var data models.SysDept
	err = e.Orm.Scopes(
		middleware.Permission(data.TableName(), p),
	).Delete(&data, ids).Error
	if err != nil {
		return lang.DataDeleteLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataDeleteCode, lang.DataDeleteLogCode, err)
	}
	return lang.SuccessCode, nil
}

// SetDeptTree 设置组织数据
func (e *SysDept) SetDeptTree(c *dto.SysDeptQueryReq) ([]dto.DeptLabel, int, error) {
	list, respCode, err := e.getList(c)
	if err != nil {
		return nil, respCode, err
	}
	m := make([]dto.DeptLabel, 0)
	for i := 0; i < len(list); i++ {
		if list[i].ParentId != 0 {
			continue
		}
		e := dto.DeptLabel{}
		e.Id = list[i].Id
		e.Label = list[i].DeptName
		deptsInfo := deptTreeCall(&list, e)

		m = append(m, deptsInfo)
	}
	return m, lang.SuccessCode, nil
}

// SetDeptPage 设置dept页面数据
func (e *SysDept) SetDeptPage(c *dto.SysDeptQueryReq) ([]models.SysDept, int, error) {
	list, respCode, err := e.getList(c)
	if err != nil {
		return nil, respCode, err
	}
	m := make([]models.SysDept, 0)
	for i := 0; i < len(list); i++ {
		/*		if list[i].ParentId != 0 {
				continue
			}*/
		if list[i].IsFlag == true {
			continue
		}
		info := e.deptPageCall(&list, list[i])
		m = append(m, info)
	}
	return m, lang.SuccessCode, nil
}

// GetWithRoleId 获取角色的部门ID集合
func (e *SysDept) GetWithRoleId(roleId int64) ([]int64, int, error) {
	deptIds := make([]int64, 0)
	deptList := make([]dto.SysRoleDeptResp, 0)
	if err := e.Orm.Table("sys_role_dept").
		Select("sys_role_dept.dept_id").
		Joins("LEFT JOIN sys_dept on sys_dept.id=sys_role_dept.dept_id").
		Where("role_id = ? ", roleId).
		Where(" sys_role_dept.dept_id not in(select sys_dept.parent_id from sys_role_dept LEFT JOIN sys_dept on sys_dept.id=sys_role_dept.dept_id where role_id =? )", roleId).
		Find(&deptList).Error; err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	for i := 0; i < len(deptList); i++ {
		deptIds = append(deptIds, deptList[i].DeptId)
	}
	return deptIds, lang.SuccessCode, nil
}

func (e *SysDept) SetDeptLabel() ([]dto.DeptLabel, int, error) {
	list := make([]models.SysDept, 0)
	err := e.Orm.Find(&list).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	m := make([]dto.DeptLabel, 0)
	var item dto.DeptLabel
	for i := range list {
		if list[i].ParentId != 0 {
			continue
		}
		item = dto.DeptLabel{}
		item.Id = list[i].Id
		item.Label = list[i].DeptName
		deptInfo := deptLabelCall(&list, item)
		m = append(m, deptInfo)
	}
	return m, lang.SuccessCode, nil
}

// Call 递归构造组织数据
func deptTreeCall(deptList *[]models.SysDept, dept dto.DeptLabel) dto.DeptLabel {
	list := *deptList
	min := make([]dto.DeptLabel, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId {
			continue
		}
		mi := dto.DeptLabel{Id: list[j].Id, Label: list[j].DeptName, Children: []dto.DeptLabel{}}
		ms := deptTreeCall(deptList, mi)
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}

// deptLabelCall
func deptLabelCall(deptList *[]models.SysDept, dept dto.DeptLabel) dto.DeptLabel {
	list := *deptList
	var mi dto.DeptLabel
	min := make([]dto.DeptLabel, 0)
	for j := 0; j < len(list); j++ {
		if dept.Id != list[j].ParentId {
			continue
		}
		mi = dto.DeptLabel{Id: list[j].Id, Label: list[j].DeptName, Children: []dto.DeptLabel{}}
		ms := deptLabelCall(deptList, mi)
		min = append(min, ms)
	}
	dept.Children = min
	return dept
}

func (e *SysDept) deptPageCall(deptlist *[]models.SysDept, menu models.SysDept) models.SysDept {
	list := *deptlist
	min := make([]models.SysDept, 0)
	for j := 0; j < len(list); j++ {
		if menu.Id != list[j].ParentId {
			continue
		}
		list[j].IsFlag = true
		mi := models.SysDept{}
		mi.Id = list[j].Id
		mi.ParentId = list[j].ParentId
		mi.DeptPath = list[j].DeptPath
		mi.DeptName = list[j].DeptName
		mi.Sort = list[j].Sort
		mi.Leader = list[j].Leader
		mi.Phone = list[j].Phone
		mi.Email = list[j].Email
		mi.Status = list[j].Status
		mi.CreatedAt = list[j].CreatedAt
		mi.Children = []models.SysDept{}
		ms := e.deptPageCall(deptlist, mi)
		min = append(min, ms)
	}
	menu.Children = min
	return menu
}

// GetSysDeptList 获取组织数据
func (e *SysDept) getList(c *dto.SysDeptQueryReq) ([]models.SysDept, int, error) {
	var list []models.SysDept
	err := e.Orm.Model(&models.SysDept{}).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
		).Find(&list).Error
	if err != nil {
		return nil, lang.DataQueryLogCode, lang.MsgLogErrf(e.Log, e.Lang, lang.DataQueryCode, lang.DataQueryLogCode, err)
	}
	return list, lang.SuccessCode, nil
}
