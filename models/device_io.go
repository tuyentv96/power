package model

type HistoryForm struct {
	Hid string `json:"hid" form:"hid" valid:"Required"`
	Limit string `json:"limit" form:"limit" valid:"Required"`
	Skip string `json:"skip" form:"skip" valid:"Required"`
}

type HistoryInfo struct {
	Total     int   `json:"total"`
	Ldevice     []HistoryDevice   `json:"devices"`
}
