package api

import (
	"mayfly-go/internal/sys/api/form"
	"mayfly-go/internal/sys/domain/entity"
	"mayfly-go/pkg/biz"
	"mayfly-go/pkg/ginx"
	"mayfly-go/pkg/global"
	"mayfly-go/pkg/req"
)

func (a *Account) SendCode(rc *req.Ctx) {
	Form := &form.UserName{}
	ginx.BindJsonAndValid(rc.GinCtx, Form)

	hcuser := &entity.Hcuser{Username: Form.Username}
	err := global.HcDb.Select("username, is_active, unickname, uuid").Where(hcuser).First(hcuser).Error
	biz.ErrIsNil(err, "账号不存在")
	biz.IsTrue(hcuser.IsEnable(), "该账号不可用")
	global.Log.Info(hcuser)
}

func (a *Account) HcLogin(rc *req.Ctx) {

}
