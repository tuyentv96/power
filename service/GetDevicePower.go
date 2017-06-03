package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	model "power/models"
)

func GetDevicePower(c *gin.Context) {
	ret_val:= model.Response{Status:true,Rcode:200}
	ret_val.Message="Success"

	token_str:=c.Request.Header.Get("auth")

	var uid,hid string

	uid= GetUidByToken(token_str)
	println("token:",token_str)
	hid=c.Query("hid")



	ldevice,err := model.GetDevicePower(uid,hid)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		c.JSON(200, ret_val)
		return
	}
	print(ldevice)

	data:= model.GetDevicePowerOutput{}

	data.Uid=uid
	data.Hid=hid
	data.Devices=ldevice

	ret_val.Data=data

	c.JSON(200, ret_val)
}