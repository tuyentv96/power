package service

import (
	"gopkg.in/gin-gonic/gin.v1"
	model "power/models"
	"strconv"
)

func GetDevicePowerInYear(c *gin.Context) {
	ret_val:= model.Response{Status:true,Rcode:200}
	ret_val.Message="Success"

	//token_str:=c.Request.Header.Get("auth")

	var did string
	var time int64

	time,_=strconv.ParseInt(c.Query("time"),0,64)
	did=c.Query("did")
	println("DID:",did)

	ldevice,err := model.GetDevicePowerInYearBYMonth(did,time)

	if err!=nil{
		ret_val.Rcode=201
		ret_val.Message=err.Error()
		c.JSON(200, ret_val)
		return
	}
	print(ldevice)

	data:= model.PowerOutput{}

	data.Did=did
	data.Devices=ldevice

	ret_val.Data=data

	c.JSON(200, ret_val)
}