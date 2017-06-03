package model

import (
	"gopkg.in/mgo.v2/bson"
	"apis/models/mongo/db"

	"errors"
	"time"
	"fmt"
)

type DevicePower struct {
	Did     string `json:"did" bson:"did"`
	Hid     string `json:"hid" bson:"hid"`
	Dname   string `json:"dname" bson:"dname"`
	Power  int `json:"power" bson:"power"`
	Type  int `json:"type" bson:"type"`
	Date time.Time `json:"date " bson:"date"`
	Time  int64 `json:"time" bson:"time"`
}

type DevicePowerOne struct {
	Power  int `json:"power" bson:"power"`
	Date time.Time `json:"date " bson:"date"`
	Time  int64 `json:"time" bson:"time"`
}

func GetDevicePower(uid string,hid string)  ([]DevicePower,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	year, month, day := time.Now().Date()
	time_now:= time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	result:= []DevicePower{}

	if err := Db.C("power").Find(bson.M{"hid": hid,"date": time_now}).Sort("dname").All(&result); err != nil {
		print("Fail")
		return result,errors.New("No home_id found")

	}

	return result,nil
}

func GetDevicePowerLimit(did string,date_skip int64,date_limit int64)  ([]DevicePowerOne,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	var date_end int64=time.Now().Local().Unix()
	date_end-= (24 * 60 * 60)*date_skip
	y,m,d:=time.Unix(date_end,0).Date()
	time_end:= time.Date(y, m, d, 0, 0, 0, 0, time.Local)

	var next_day int64=time.Now().Local().Unix()
	next_day-= (24 * 60 * 60)*(date_limit+date_skip)
	y1,m1,d1:=time.Unix(next_day,0).Date()
	time_from:= time.Date(y1, m1, d1, 0, 0, 0, 0, time.Local)

	result:= []DevicePowerOne{}

	if err := Db.C("power").Find(bson.M{"did": did,
		"date": bson.M{
		"$gt": time_from,
		"$lte": time_end,
		},
	}).Sort("dname").All(&result); err != nil {
		print("Fail")
		return result,errors.New("No home_id found")

	}

	return result,nil
}

func GetDevicePowerByTime(did string,date_start int64,date_end int64)  ([]DevicePowerOne,error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	result:= []DevicePowerOne{}
	println("DAte:",date_start,date_end)

	if err := Db.C("power").Find(bson.M{"did": did,
		"time": bson.M{
			"$gte": date_start,
			"$lte": date_end,
		},
	}).Sort("dname").All(&result); err != nil {
		print("Fail")
		return result,errors.New("No home_id found")

	}

	return result,nil
}


func GetHomePowerInYearBYMonth(hid string,time_now int64)  (ret []bson.M, err_r error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()


	y,_,_:=time.Unix(time_now,0).Local().Date()
	time_end:= time.Date(y, 12, 31, 23, 59, 59, 0, time.Local)
	time_from:= time.Date(y, 1, 1, 0, 0, 0, 0, time.Local)

	fmt.Printf("%+v",time_end,time_from,"\n")

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"time": bson.M{
				"$gte": time_from.Unix(),
				"$lte": time_end.Unix(),
			},
				"hid": hid,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"month": bson.M{ "$month": "$date" },
				},
				"power_total": bson.M{ "$sum": "$power" },
			},
		},
		bson.M{ "$sort": bson.M{ "_id.month": 1 } },
	}
	resp := []bson.M{}

	err:= Db.C("power").Pipe(pipeline).All(&resp)

	if err != nil {
		//handle error
		println("Query fai;")
	}
	//fmt.Printf("%+v",resp)
	return resp,nil
}

func GetHomePowerInMonthByDate(hid string,time_now int64)  (ret []bson.M, err_r error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	y,m,_:=time.Unix(time_now,0).Local().Date()
	time_end:= time.Date(y, m, 31, 23, 59, 59, 0, time.Local)
	time_from:= time.Date(y, m, 1, 0, 0, 0, 0, time.Local)

	fmt.Printf("%+v",time_end,time_from,m,"\n")

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"time": bson.M{
				"$gte": time_from.Unix(),
				"$lte": time_end.Unix(),
			},
				"hid": hid,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"date": bson.M{ "$dayOfMonth": "$date" },
				},
				"power_total": bson.M{ "$sum": "$power" },
			},
		},
		bson.M{ "$sort": bson.M{ "_id.date": 1 } },
	}
	resp := []bson.M{}

	err:= Db.C("power").Pipe(pipeline).All(&resp)

	if err != nil {
		//handle error
		println("Query fai;")
	}
	//fmt.Printf("%+v",resp)
	return resp,nil
}

func GetDevicePowerInMonthByDate(did string,time_now int64)  (ret []bson.M, err_r error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	y,m,_:=time.Unix(time_now,0).Local().Date()
	time_end:= time.Date(y, m, 31, 23, 59, 59, 0, time.Local)
	time_from:= time.Date(y, m, 1, 0, 0, 0, 0, time.Local)

	fmt.Printf("%+v",time_end,time_from,m,did,"\n")

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"time": bson.M{
				"$gte": time_from.Unix(),
				"$lte": time_end.Unix(),
			},
				"did": did,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"date": bson.M{ "$dayOfMonth": "$date" },
				},
				"power_total": bson.M{ "$sum": "$power" },
			},
		},
		bson.M{ "$sort": bson.M{ "_id.date": 1 } },
	}
	resp := []bson.M{}

	err:= Db.C("power").Pipe(pipeline).All(&resp)

	if err != nil {
		//handle error
		println("Query fai;")
	}
	//fmt.Printf("%+v",resp)
	return resp,nil
}

func GetDevicePowerInYearBYMonth(did string,time_now int64)  (ret []bson.M, err_r error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()


	y,_,_:=time.Unix(time_now,0).Local().Date()
	time_end:= time.Date(y, 12, 31, 23, 59, 59, 0, time.Local)
	time_from:= time.Date(y, 1, 1, 0, 0, 0, 0, time.Local)

	fmt.Printf("%+v",time_end,time_from,"\n")

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"time": bson.M{
				"$gte": time_from.Unix(),
				"$lte": time_end.Unix(),
			},
				"did": did,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"month": bson.M{ "$month": "$date" },
				},
				"power_total": bson.M{ "$sum": "$power" },
			},
		},
		bson.M{ "$sort": bson.M{ "_id.month": 1 } },
	}
	resp := []bson.M{}

	err:= Db.C("power").Pipe(pipeline).All(&resp)

	if err != nil {
		//handle error
		println("Query fai;")
	}
	//fmt.Printf("%+v",resp)
	return resp,nil
}


func GetRankingDevicePowerInMonth(hid string,time_now int64)  (ret []bson.M, err_r error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()


	y,m,_:=time.Unix(time_now,0).Local().Date()
	time_end:= time.Date(y, m, 31, 23, 59, 59, 0, time.Local)
	time_from:= time.Date(y, m, 1, 0, 0, 0, 0, time.Local)

	fmt.Printf("%+v",time_end,time_from,"\n")

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"time": bson.M{
				"$gte": time_from.Unix(),
				"$lte": time_end.Unix(),
			},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":  bson.M{
					"did": "$did",
					"hid": "$hid",
					"dname": "$dname",
				},
				"power_total": bson.M{ "$sum": "$power" },
			},
		},
		bson.M{ "$sort": bson.M{ "power_total": -1 } },
	}
	resp := []bson.M{}

	err:= Db.C("power").Pipe(pipeline).All(&resp)

	if err != nil {
		//handle error
		println("Query fai;")
	}
	//fmt.Printf("%+v",resp)
	return resp,nil
}

func GetRankingDevicePowerInYear(hid string,time_now int64)  (ret []bson.M, err_r error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()


	y,_,_:=time.Unix(time_now,0).Local().Date()
	time_end:= time.Date(y, 12, 31, 23, 59, 59, 0, time.Local)
	time_from:= time.Date(y, 1, 1, 0, 0, 0, 0, time.Local)

	fmt.Printf("%+v",time_end,time_from,"\n")

	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{"time": bson.M{
				"$gte": time_from.Unix(),
				"$lte": time_end.Unix(),
			},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id":  bson.M{
					"did": "$did",
					"hid": "$hid",
					"dname": "$dname",
				},
				"power_total": bson.M{ "$sum": "$power" },
			},
		},
		bson.M{ "$sort": bson.M{ "power_total": -1 } },
	}
	resp := []bson.M{}

	err:= Db.C("power").Pipe(pipeline).All(&resp)

	if err != nil {
		//handle error
		println("Query fai;")
	}
	//fmt.Printf("%+v",resp)
	return resp,nil
}


func GetRankingDevicePowerAll(hid string)  (ret []bson.M, err_r error){
	Db := db.MgoDb{}
	Db.Init()
	defer Db.Close()

	pipeline := []bson.M{
		bson.M{
			"$group": bson.M{
				"_id":  bson.M{
					"did": "$did",
					"hid": "$hid",
					"dname": "$dname",
				},
				"power_total": bson.M{ "$sum": "$power" },
			},
		},
		bson.M{ "$sort": bson.M{ "power_total": -1 } },
	}
	resp := []bson.M{}

	err:= Db.C("power").Pipe(pipeline).All(&resp)

	if err != nil {
		//handle error
		println("Query fai;")
	}
	//fmt.Printf("%+v",resp)
	return resp,nil
}



