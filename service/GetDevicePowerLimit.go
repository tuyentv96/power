package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	model "power/models"
	"strconv"
)

func GetDevicePowerLimit(c *gin.Context) {
	var did,uid string
	var date_limit,date_skip int64
	ret_val:= model.Response{Status:true,Rcode:200}

	did= c.Query("did")
	date_skip,_=strconv.ParseInt(c.Query("date_skip"),0,64)
	date_limit,_=strconv.ParseInt(c.Query("date_limit"),0,64)
	token_str:=c.Request.Header.Get("auth")
	uid= GetUidByToken(token_str)

	dev,err1:= model.FindDeviceByID(did)
	if err1!=""{
		ret_val.Rcode=202
		ret_val.Message="No device found"
		c.JSON(200, ret_val)
	}

	ldevice,err := model.GetDevicePowerLimit(did,date_skip,date_limit)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		c.JSON(200, ret_val)
	}
	print(ldevice)

	data:= model.GetDevicePowerOneOutput{}

	data.Uid=uid
	data.Hid=dev.Hid
	data.Did=dev.Did
	data.Dname=dev.Dname
	data.Devices=ldevice
	data.Type=dev.Type
	data.Status=dev.Status
	data.Total=len(ldevice)

	ret_val.Data=data


	c.JSON(200, ret_val)
}