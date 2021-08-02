package model

type BaseParams struct {
	// 默认传参
	CustomerId     int64  `json:"customerId" validate:"required"`
	OpenId         string `json:"openId" validate:"required"`
	DeviceId       string `json:"deviceId" validate:"required"`
	HostAppName    string `json:"hostAppName" validate:"-"`
	HostAppVersion string `json:"hostAppVersion" validate:"-"`
	Platform       string `json:"platform" validate:"-"`
	GameVersion    string `json:"gameVersion" validate:"-"`
	TopLevel       int64  `json:"topLevel" validate:"-"`
	Diamond        int64  `json:"diamond" validate:"-"`

	// 自定义参数
	AccessToken       string          `json:"access_token" form:"access_token" validate:"-"`
	RegisterTime      int64           `json:"registerTime" validate:"-"`
	InitialAppVersion string          `json:"initialAppVersion" validate:"-"`
	Rid               int64           `json:"rid" validate:"-"`
	ClientIp          string          `json:"clientIp" validate:"-"`
	Tags              map[int64]int64 `json:"tags" validate:"-"`
}
