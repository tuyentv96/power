package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	model "power/models"
)

func GetRankingDevicePowerAll(c *gin.Context) {
	ret_val:= model.Response{Status:true,Rcode:200}
	ret_val.Message="Success"

	//token_str:=c.Request.Header.Get("auth")

	var hid string
	hid=c.Query("hid")
	println("DID:",hid)

	ldevice,err := model.GetRankingDevicePowerAll(hid)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		c.JSON(200, ret_val)
		return
	}
	print(ldevice)

	data:= model.PowerOutput{}

	data.Hid=hid
	data.Devices=ldevice

	ret_val.Data=data

	c.JSON(200, ret_val)
}