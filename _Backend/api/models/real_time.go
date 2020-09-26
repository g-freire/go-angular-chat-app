package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type RealTime struct {
	ID              *primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Timestamp 		 float64   `json:"timestamp" bson:"timestamp"`
	TD_WEIGHT_COMPENSATOR_PRESSURE_STP_LOW_POSITION  SensorType   `json:"TD_WEIGHT_COMPENSATOR_PRESSURE_STP_LOW_POSITION" bson:"TD_WEIGHT_COMPENSATOR_PRESSURE_STP_LOW_POSITION"`
	TD_BLOWER_1_RUNNING SensorType   `json:"TD_BLOWER_1_RUNNING" bson:"TD_BLOWER_1_RUNNING"`
}

type SensorType struct{
	Value float64   `json:"value" bson:"value"`
	Uom   string   `json:"uom" bson:"uom"`
}
