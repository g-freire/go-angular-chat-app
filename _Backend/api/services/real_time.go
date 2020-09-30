package services

import (
	"../models"
	"../repository"
	"go.mongodb.org/mongo-driver/bson"
)


func GetRealTimeInfo(host string) models.RealTime{
	m := repository.MongoGenericRepository{}
	m.GetClient(host)
	bsonIbopInfo:= m.ReadLastN("QA_ODN2","REAL_TIME", 1, true)
	RealTime := parseRealTimeBSONtoModel(bsonIbopInfo)
	return RealTime
}

func parseRealTimeBSONtoModel(bsonIbopInfo []bson.M) models.RealTime {
	var RealTime models.RealTime
	bsonBytes, _ := bson.Marshal(bsonIbopInfo[0])
	bson.Unmarshal(bsonBytes, &RealTime)
	return RealTime
}
