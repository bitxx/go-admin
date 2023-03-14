package dto

type UserJoin struct {
	TrueName string `search:"type:contains;column:true_name;table:app_user" form:"trueName" comment:"真实姓名"`
	UserName string `search:"type:contains;column:user_name;table:app_user" form:"userName" comment:"昵称"`
	Mobile   string `search:"type:exact;column:mobile;table:app_user" form:"mobile" comment:"手机号"`
	Email    string `search:"type:exact;column:email;table:app_user" form:"email" comment:"邮箱"`
}

type LevelJoin struct {
	LevelName string `search:"type:contains;column:level_name;table:app_user_level" form:"name" comment:"等级名称"`
}
