package entity

type Hcuser struct {
	Username  string `json:"username"`
	Unickname string `json:"unickname"`
	IsActive  int8   `json:"is_active"`
	Uuid      int64  `json:"uuid"`
}

func (a *Hcuser) TableName() string {
	return "auth_user"
}

func (a *Hcuser) IsEnable() bool {
	return a.IsActive == AccountEnableStatus
}
