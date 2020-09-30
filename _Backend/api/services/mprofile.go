package services

import (
	"../models"
	"../repository"
	"go.mongodb.org/mongo-driver/bson"
)


func GetMatrixProfile(host string) models.MatrixProfile {
	m := repository.MongoGenericRepository{}
	m.GetClient(host)
	bsonMProfile := m.ReadLastN("QA_ODN2","MP_RESULTS", 5 , true)
	MatrixProfile := parseMatrixProfiletoModel(bsonMProfile)
	return MatrixProfile
}

func parseMatrixProfiletoModel(bsonMProfile []bson.M) models.MatrixProfile{
	var MProfile models.MatrixProfile
	bsonBytes, _ := bson.Marshal(bsonMProfile[0])
	bson.Unmarshal(bsonBytes, &MProfile)
	return MProfile
}