package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RealTime struct {
	ID              	*primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Timestamp 		 	float64   `json:"timestamp" bson:"timestamp"`
	TDA_MOTOR_TEMP_1   SensorType   `json:"TDA_MOTOR_TEMP_1" bson:"TDA_MOTOR_TEMP_1"`
	TDA_MOTOR_TEMP_2   SensorType   `json:"TDA_MOTOR_TEMP_2" bson:"TDA_MOTOR_TEMP_2"`
	TDB_MOTOR_TEMP_1   SensorType   `json:"TDB_MOTOR_TEMP_1" bson:"TDB_MOTOR_TEMP_1"`
	TDB_MOTOR_TEMP_2   SensorType   `json:"TDB_MOTOR_TEMP_2" bson:"TDB_MOTOR_TEMP_2"`
	TDA_TORQUE      	SensorType   `json:"TDA_TORQUE" bson:"TDA_TORQUE"`
	TDB_TORQUE 			SensorType   `json:"TDB_TORQUE" bson:"TDB_TORQUE"`
	TDA_SPEED          SensorType   `json:"TDA_SPEED" bson:"TDA_SPEED"`
	TDB_SPEED 			SensorType   `json:"TDB_SPEED" bson:"TDB_SPEED"`
	TDA_POWER  			SensorType   `json:"TDA_POWER" bson:"TDA_POWER"`
	TDB_POWER 			SensorType   `json:"TDB_POWER" bson:"TDB_POWER"`

}

type SensorType struct{
	Value float64   `json:"value" bson:"value"`
	Uom   string   `json:"uom" bson:"uom"`
}