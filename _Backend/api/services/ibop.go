package services

import (
	"../models"
	"../repository"
	"go.mongodb.org/mongo-driver/bson"

)


func GetIbopInfo(host string) models.IBOPInfo{
	m := repository.MongoGenericRepository{}
	m.GetClient(host)
	bsonIbopInfo:= m.ReadLastN("QA_ODN2","IBOP_INFO", 1)
	IbopInfo := parseBSONtoModel(bsonIbopInfo)
	return IbopInfo
}

func parseBSONtoModel(bsonIbopInfo []bson.M) models.IBOPInfo {
	var IbopInfo models.IBOPInfo
	bsonBytes, _ := bson.Marshal(bsonIbopInfo[0])
	bson.Unmarshal(bsonBytes, &IbopInfo)
	return IbopInfo
}