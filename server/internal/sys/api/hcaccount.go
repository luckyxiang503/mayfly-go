package api

import (
	"fmt"
	"mayfly-go/internal/sys/api/form"
	"mayfly-go/internal/sys/domain/entity"
	"mayfly-go/pkg/biz"
	"mayfly-go/pkg/ginx"
	"mayfly-go/pkg/global"
	"mayfly-go/pkg/rediscli"
	"mayfly-go/pkg/req"
)

func (a *Account) SendCode(rc *req.Ctx) {
	Form := &form.UserName{}
	ginx.BindJsonAndValid(rc.GinCtx, Form)

	hcuser := &entity.Hcuser{Username: Form.Username}
	err := global.HcDb.Select("username, is_active, unickname, uuid").Where(hcuser).First(hcuser).Error
	biz.ErrIsNil(err, "账号不存在")
	biz.IsTrue(hcuser.IsEnable(), "该账号不可用")
	redisKey := fmt.Sprintf("user:%s:%d", hcuser.Username, hcuser.Uuid)
	code := "ABCDEF"

	// 发送验证码到企业微信
	biz.ErrIsNil(_sendCodeToWeChat(code), "发送验证码失败")

	// 保存验证码到redis
	biz.NotNil(rediscli.GetCli(), "redis连接不可用")
	rediscli.Set(redisKey, code, 120*1000*1000*1000)
}

func (a *Account) HcLogin(rc *req.Ctx) {

}

func _sendCodeToWeChat(code string) error {
	global.Log.Info(code)
	return nil
}
