package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type MatrixProfile struct {
	ID             *primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Timestamp 		 int64 `json:"timestamp" bson:"timestamp"`
	MatrixProfile   string   `json:"matrix_profile" bson:"matrix_profile"`
	GlobalTime 		 primitive.DateTime`json:"global_time" bson:"global_time"`
	TrainVersion 	 string   `json:"train_version" bson:"train_version"`
	TrainWindow     string   `json:"train_window" bson:"train_window"`
}



