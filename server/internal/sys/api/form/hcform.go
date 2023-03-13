package form

type UserName struct {
	Username string `json:"username" binding:"required"`
}

// 登录表单
type HcLoginForm struct {
	Username string `json:"username" binding:"required"`
	Code     string `json:"code" binding:"required"`
}
