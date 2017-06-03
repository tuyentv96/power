package model

import (

	"apis/models/mongo/db"
	"time"
	"gopkg.in/mgo.v2/bson"
	"errors"
)

type TimerDevice struct {
	Hid     string `json:"hid" bson:"hid"`
	Uid     string `json:"uid" bson:"uid"`
	Timer_Name string `json:"timer_name" bson:"timer_name" form:"timer_name"`
	Device     []Device `json:"device" bson:"device"`
	Time int64 `json:"time" bson:"time"`
	Execute bool `json:"execute" bson:"execute"`
	TimerId string `json:"timer_id" bson:"timer_id"`
	Create_At int64 `json:"create_at" bson:"create_at"`
}

type CronDevice struct {
	Hid     string `json:"hid" bson:"hid"`
	Uid     string `json:"uid" bson:"uid"`
	Cron_Name string `json:"cron_name" bson:"cron_name" form:"cron_name"`
	Device     []Device `json:"device" bson:"device"`
	Time string `json:"time" bson:"time"`
	Date string `json:"date" bson:"date"`
	Enable bool `json:"enable" bson:"enable"`
	TimerId string `json:"timer_id" bson:"timer_id"`
	Create_At int64 `json:"create_at" bson:"create_at"`
}

func FindListDeviceByID(data []DeviceInfo)  ([]Device){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := []Device{}
	device:= Device{}

	for _,dev:= range data {
		if err := Db.C("devices").Find(bson.M{"did": dev.Did}).One(&device); err == nil {
			result=append(result,device)
		}
	}

	return result

}

func CheckTimerExisted(uid string,timer_name string,time int)  (bool) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	//result:= TimerDevice{}

	if count,err := Db.C("timer").Find(bson.M{"uid":uid,"$or":[]interface{}{
		bson.M{"time": time},
		bson.M{"timer_name": timer_name},
	}}).Count(); err != nil {
		print("No recode found")
		return true
	} else {
		print("Record:",count)
		if count>0 {
			return true
		}

	}
	return false


}

func CheckCronExisted(uid string,cron_name string,time string)  (bool) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	//result:= TimerDevice{}

	if count,err := Db.C("cron").Find(bson.M{"uid":uid,"$or":[]interface{}{
		bson.M{"time": time},
		bson.M{"timer_name": cron_name},
	}}).Count(); err != nil {
		print("No recode found")
		return true
	} else {
		print("Record:",count)
		if count>0 {
			return true
		}

	}
	return false


}

func SaveTimerDevice(dev []Device,timer_name string,hid string,uid string,time_exp int64)  TimerDevice{
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result:=TimerDevice{Timer_Name:timer_name,Device: dev,Hid:hid,Uid:uid,Time:time_exp,Execute:false,TimerId:bson.NewObjectId().Hex(),Create_At:time.Now().Unix()}

	if err := Db.C("timer").Insert(result); err != nil {
		print("Fail")

	}

	return result
}

func GetTimerDeviceByDid(uid string,did string)  ([]TimerDevice,bool) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result:= []TimerDevice{}

	if err := Db.C("timer").Find(bson.M{"uid":uid, "device": bson.M{"$elemMatch":bson.M{"did":did}}}).Sort("-time").All(&result); err != nil {
		print("Fail clgt")
		return result,true
	}

	return result,false
}

func GetTimerDeviceByUid(uid string)  ([]TimerDevice,bool) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result:= []TimerDevice{}

	if err := Db.C("timer").Find(bson.M{"uid": uid}).Sort("-time").All(&result); err != nil {
		print("Fail clgt")
		return result,true
	}

	return result,false
}

func CheckTimerExcuteByTimerID(timer_id string)  (bool,error) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result:= TimerDevice{}

	if err := Db.C("timer").Find(bson.M{"timer_id": timer_id}).One(&result); err != nil {
		print("Fail clgt")
		return false,errors.New("timer_id not found")
	}

	if result.Execute==false {
		return false,nil
	}

	return true,nil
}


func DeleteTimerDevice(timer_id string)  bool {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("timer").Remove(bson.M{"timer_id": timer_id}); err != nil {
		print("Faillllllllllllllllllllllllllllllllllllllllll")
		return true
	}
	print("Delete timer completed")
	return false
}

func UpdateStatusTimer(timer_id string)  bool{
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	//Update status
	colQuerier := bson.M{"timer_id": timer_id}
	change := bson.M{"$set": bson.M{"execute": true}}

	if err := Db.C("timer").Update(colQuerier,change); err != nil {
		print("Faillllllllllllllllllllllllllllllllllllllllll")
		return true
	}
	print("Updatetimer completed")
	return false
}

func SaveCronDevice(dev []Device,cron_name string,hid string,uid string,_time string,date string)  CronDevice{
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result:=CronDevice{Cron_Name:cron_name,Device: dev,Hid:hid,Uid:uid,Time: _time,Enable: true,TimerId:bson.NewObjectId().Hex(),Create_At:time.Now().Unix()}

	if err := Db.C("cron").Insert(result); err != nil {
		print("Fail")

	}

	return result
}

func GetCronByHid(uid string,hid string)  ([]CronDevice,bool) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result:= []CronDevice{}

	if err := Db.C("cron").Find(bson.M{"uid":uid,"hid": hid}).Sort("-create_at").All(&result); err != nil {
		print("Fail clgt")
		return result,true
	}

	return result,false
}

func StopCronTimer(timer_id string)  error {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	colQuerier := bson.M{"timer_id": timer_id}
	change := bson.M{"$set": bson.M{"enable": false}}

	if err := Db.C("cron").Update(colQuerier,change); err != nil {
		print("Faillllllllllllllllllllllllllllllllllllllllll")
		return errors.New("Timer id not found")
	}

	print("Delete timer completed")
	return nil
}

func FindCronByTimerID(timer_id string)  (CronDevice,error) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result:= CronDevice{}

	if err := Db.C("cron").Find(bson.M{"timer_id": timer_id}).One(&result); err != nil {
		print("Fail clgt")
		return result,errors.New("timer_id not found")
	}

	return result,nil
}

func CheckCronEnableByTimerID(timer_id string)  (bool,error) {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result:= CronDevice{}

	if err := Db.C("cron").Find(bson.M{"timer_id": timer_id}).One(&result); err != nil {
		print("Fail clgt")
		return false,errors.New("timer_id not found")
	}

	if result.Enable==false {
		return false,nil
	}

	return true,nil
}

func DeleteCronTimer(uid string,timer_id string)  bool {
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("cron").Remove(bson.M{"uid": uid,"timer_id": timer_id}); err != nil {
		print("Faillllllllllllllllllllllllllllllllllllllllll")
		return true
	}
	print("Delete timer completed")
	return false
}

func StartCronTimer(timer_id string)  bool{
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	//Update status
	colQuerier := bson.M{"timer_id": timer_id}
	change := bson.M{"$set": bson.M{"enable": true}}

	if err := Db.C("cron").Update(colQuerier,change); err != nil {
		print("Faillllllllllllllllllllllllllllllllllllllllll")
		return true
	}
	print("Updatetimer completed")
	return false
}