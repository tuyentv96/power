package model

import (
	"gopkg.in/mgo.v2/bson"
	"apis/models/mongo/db"

)

type Device struct {
	Did     string `json:"did" bson:"did"`
	Hid     string `json:"hid" bson:"hid"`
	Dname   string `json:"dname" bson:"dname"`
	Status  int `json:"status" bson:"status"`
	Type  int `json:"type" bson:"type"`
	Roomid  string `json:"roomid" bson:"roomid"`
	Roomname  string `json:"roomname" bson:"roomname"`
}

type LDevice struct {
	UID   string `json:"uid"`
	Uname string `json:"uname"`
	Lhome []Lhome `json:"homes"`
	Ldevice []Device `json:"devices"`
	Permission []Permission `json:"permission"`
}

type Permission struct {
	Hid  string `json:"hid"`
	Did string    `json:"did"`
}

type Lhome struct {
	Hid  string `json:"hid"`
	Type int    `json:"type"`
}

type HistoryDevice struct {
	Did     string `json:"did" bson:"did"`
	Hid     string `json:"hid" bson:"hid"`
	Dname   string `json:"dname" bson:"dname"`
	Status  int `json:"status" bson:"status"`
	Type  int `json:"type" bson:"type"`
	Uid  string `json:"uid" bson:"uid"`
	Time int64 `json:"time" bson:"time"`
}

func MGetHistoryHome(hid string,skip int,limit int)  (record []HistoryDevice,code int,error bool){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("history").Find(bson.M{"hid": hid}).Limit(limit).Skip(skip).Sort("-time").All(&record); err != nil {
		print("Fail")
		return record,400,true
	}
	print("Succc",record,"+++",hid,skip,limit)
	return record,200,false
}

func MGetHistoryHomeByTime(hid string,skip int,limit int,time_start int,time_end int)  (record []HistoryDevice,code int,error bool){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("history").Find(bson.M{"hid": hid,"time": bson.M{"$gte": time_start, "$lte": time_end} }).Limit(limit).Skip(skip).Sort("-time").All(&record); err != nil {
		print("Fail")
		return record,400,true
	}
	print("Succc",record,"+++",hid,skip,limit)
	return record,200,false
}

func MGetHistoryDevice(did string,skip int,limit int)  (record []HistoryDevice,code int,error bool){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("history").Find(bson.M{"did": did}).Limit(limit).Skip(skip).Sort("-time").All(&record); err != nil {
		print("Fail")
		return record,400,true
	}
	print("Succc",record,"+++",did,skip,limit)
	return record,200,false
}

func MGetHistoryDeviceByTime(did string,skip int,limit int,time_start int,time_end int)  (record []HistoryDevice,code int,error bool){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("history").Find(bson.M{"did": did,"time": bson.M{"$gte": time_start, "$lte": time_end} }).Limit(limit).Skip(skip).Sort("-time").All(&record); err != nil {
		print("Fail")
		return record,400,true
	}
	print("Succc",record,"+++",did,skip,limit)
	return record,200,false
}

func FindDeviceByID(did string)  (Device,string){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := Device{}

	if err := Db.C("devices").Find(bson.M{"did": did}).One(&result); err != nil {
		print("Fail")
		return result,"No device found"

	}

	return result,""

}

func MGetDevice(uid string)  LDevice{

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := LDevice{}


	if err := Db.C("users").Find(bson.M{"uid": uid}).One(&result); err != nil {
		print("Fail")
		return result

	}

	lh:= []bson.M{}

	for i:=0;i<len(result.Lhome);i++ {
		temp:= bson.M{"hid":result.Lhome[i].Hid}
		lh=append(lh,temp)

	}

	//fmt.Print("BSSS",lh)
	if err1 := Db.C("devices").Find(bson.M{"$or": lh}).Sort("hid","dname").All(&result.Ldevice); err1 != nil {
		print("Fail")
		return result

	}


	return result
}
