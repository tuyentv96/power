package model


type UpdatePasswordForm struct {
	Code string `json:"code" bson:"code" form:"code"`
	OldPwd string `json:"old_pwd" bson:"pwd" form:"old_pwd"`
	NewPwd string `json:"new_pwd" bson:"pwd" form:"new_pwd"`
}

type GetAllDeviceRsp struct {
	Rcode int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
	Status bool `json:"status"`
}

type GetHistoryRsp struct {
	Rcode int `json:"code"`
	Message string `json:"message"`
	Data HistoryInfo `json:"data"`
	Status bool `json:"status"`
}

type MGetDeviceByHidData struct {
	Total int `json:"total"`
	Devices []Device `json:"devices"`
}

type Home struct {
	Hid string `json:"hid"`
	Home_Name string `json:"home_name"`
	Address string `json:"address"`
	Phone string `json:"phone"`
}

type MGetHomeData struct {
	Total int `json:"total"`
	Homes []Home `json:"homes"`

}

type UserProfile struct {
	Uid     string `json:"uid" bson:"uid" form:"uid"`
	Uname     string `json:"uname" bson:"uname" form:"uname"`
	Phone     string `json:"phone" bson:"phone" form:"phone"`
	Address     string `json:"address" bson:"address" form:"address"`
	Email    string `json:"email" bson:"email" form:"email"`

}
