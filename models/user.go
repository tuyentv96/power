package model

import (
	"gopkg.in/mgo.v2/bson"
	"apis/models/mongo/db"
	"fmt"
	"crypto/sha256"
	"errors"
)

type User struct {
		Uid     string `json:"uid" bson:"uid" form:"uid"  `
		Pwd     string `json:"pwd,omitempty" bson:"pwd" form:"pwd"  `
}

type Object struct {
	ObjectId   string
	Score      int64
	PlayerName string
}


type Userpsmdevice struct {
	Lhome []struct {
		Hid  string `json:"hid"`
		Type int    `json:"type"`
	} `json:"lhome"`

	UID   string `json:"uid"`
	Uname string `json:"uname"`
}

func (u *User) FindByUid(uid string) (int,bool){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	if err := Db.C("users").Find(bson.M{"uid": uid}).One(&u); err != nil {
		print("Fail clgt")
		return 104,true
	}

	if u.Uid=="" {
		return 104,true
	}

	return 200,false

}

func (u *User)	CheckPwd(pwd string)  (int,bool){

	encryted_pwd:= fmt.Sprintf("%x",sha256.Sum256([]byte(pwd)))

	if u.Pwd!=encryted_pwd {
		return 410,true
	}

	return 200,false

}
func (u *User) ClearPass() {
	u.Pwd=""
}

func MGetDeviceByHid(uid string,hid string)  []Device{

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	user_data := LDevice{}
	result:= []Device{}


	if err := Db.C("users").Find(bson.M{"uid": uid}).One(&user_data); err != nil {
		print("Fail")
		return result

	}

	home_owner:=false

	for i:=0;i<len(user_data.Lhome);i++ {
		if user_data.Lhome[i].Hid == hid{
			home_owner=true
		}
	}

	if home_owner==false {
		print("User not owner home !!!")
		return result
	}

	//fmt.Print("BSSS",lh)
	if err1 := Db.C("devices").Find(bson.M{"hid": hid}).Sort("dname").All(&result); err1 != nil {
		print("Fail")
		return result

	}


	return result
}

func MGetHome(uid string)  MGetHomeData{

	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()
	result := MGetHomeData{}

	data:= LDevice{}


	if err := Db.C("users").Find(bson.M{"uid": uid}).One(&data); err != nil {
		print("Fail")
		return result

	}
	homes:= []Home{}
	home:= Home{}

	for i:=0;i<len(data.Lhome);i++ {
		if err := Db.C("homes").Find(bson.M{"hid": data.Lhome[i].Hid}).One(&home); err != nil {
			print("Fail")
		}
		homes=append(homes,home)

	}

	result.Total=len(data.Lhome)
	result.Homes=homes

	return result
}



func GetProfileByUid(uid string) (UserProfile,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	u:= UserProfile{}

	if err := Db.C("users").Find(bson.M{"uid": uid}).One(&u); err != nil {
		print("Fail clgt")
		return u,errors.New("Load fail")
	}

	println("Load success")
	return u,nil

}

func UpdateUserProfile(u UserProfile) (UserProfile,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	// Update
	colQuerier := bson.M{"uid": u.Uid}
	change := bson.M{"$set": bson.M{"uname": u.Uname,"address": u.Address, "phone": u.Phone,"email": u.Email}}
	err := Db.C("users").Update(colQuerier, change)
	if err != nil {
		return u,err
	}

	println("Load success")
	return u,nil

}