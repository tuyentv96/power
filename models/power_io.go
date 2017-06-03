package model

import "gopkg.in/mgo.v2/bson"

type GetDevicePowerOutput struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Did string `json:"did" bson:"did" form:"did"`
	Dname   string `json:"dname" bson:"dname"`
	Status  int `json:"status" bson:"status"`
	Type  int `json:"type" bson:"type"`
	Total int `json:"total" bson:"total" form:"total"`
	Devices []DevicePower `json:"power_data" bson:"power_data" form:"power_data"`
}

type GetDevicePowerOneOutput struct {
	Uid string `json:"uid" bson:"uid" form:"uid"`
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Did string `json:"did" bson:"did" form:"did"`
	Dname   string `json:"dname" bson:"dname"`
	Status  int `json:"status" bson:"status"`
	Type  int `json:"type" bson:"type"`
	Total int `json:"total" bson:"total" form:"total"`
	Devices []DevicePowerOne `json:"power_data" bson:"power_data" form:"power_data"`
}

type PowerOutput struct {
	Uid string `json:"uid,omitempty" bson:"uid" form:"uid"`
	Hid string `json:"hid,omitempty" bson:"hid" form:"hid"`
	Did string `json:"did,omitempty" bson:"did" form:"did"`
	Dname   string `json:"dname,omitempty" bson:"dname"`
	Status  int `json:"status,omitempty" bson:"status"`
	Type  int `json:"type,omitempty" bson:"type"`
	Total int `json:"total,omitempty" bson:"total" form:"total"`
	Devices []bson.M `json:"power_data,omitempty" bson:"power_data" form:"power_data"`
}

type PowerInput struct {
	Uid string `json:"uid,omitempty" bson:"uid" form:"uid"`
	Hid string `json:"hid,omitempty" bson:"hid" form:"hid"`
	Did string `json:"did,omitempty" bson:"did" form:"did"`
	Time int64 `json:"time,omitempty" bson:"time" form:"time"`
	Date_Start int64  `json:"date_start,omitempty" bson:"date_start" form:"date_start"`
	Date_End int64  `json:"date_end,omitempty" bson:"date_end" form:"date_end"`
	Date_Limit int64  `json:"date_limit,omitempty" bson:"date_limit" form:"date_limit"`
	Date_Skip int64  `json:"date_skip,omitempty" bson:"date_skip" form:"date_skip"`
}



type GetRankingDevicePowerOutput struct {
	Hid string `json:"hid" bson:"hid" form:"hid"`
	Total int `json:"total" bson:"total" form:"total"`
	Devices []bson.M `json:"power_data" bson:"power_data" form:"power_data"`
}

