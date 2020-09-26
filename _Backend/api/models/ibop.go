package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type IBOPInfo struct {
	ID           *primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	IbopId 		    string   `json:"ibop_id" bson:"ibop_id"`
	CreationTime    primitive.DateTime   `json:"creation_time" bson:"creation_time"`
	CreationTimeISO string   `json:"creation_time_iso" bson:"creation_time_iso"`
	TimeUI 			string   `json:"time_ui" bson:"time_ui"`
	StartDate 		string   `json:"start_date" bson:"start_date"`
	EndDate 		string   `json:"end_date" bson:"end_date"`
	IsActive        bool     `json:"is_active" bson:"is_active"`
	CurrentState    string   `json:"current_state" bson:"current_state"`
	TotalCycles     string   `json:"total_cycles" bson:"total_cycles"`
	OpenCycles 	    string   `json:"open_cycles" bson:"open_cycles"`
	CloseCycles     string   `json:"close_cycles" bson:"close_cycles"`
}


