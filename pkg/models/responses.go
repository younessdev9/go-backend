package models

type ResponseObject struct {
	Message string `json:"message"`
}

type CreatedCar struct {
	Id string `json:"id" bson:"_id,omitempty"`
}
