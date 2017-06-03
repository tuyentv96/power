package model

// LoginInfo definiton.
type LoginForm struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Pwd string `json:"pwd" bson:"pwd" form:"pwd"`
}
type LoginRsp struct {
	Rcode int `json:"code"`
	Message string `json:"message"`
	UserInfo UserInfo `json:"data"`
	Status bool `json:"status"`
}

type UserInfo struct {
	Uid string `json:"uid"`
	Token string `json:"token"`
	//UName string `json:"uname"`
}
type GetCodeForm struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
}

type VerifyCodeForm struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Code string `json:"code" bson:"code" form:"code"`
}

type Message struct {
	Rcode int `json:"code"`
	Message string `json:"message"`
	Status bool `json:"status"`
}

