package models

type ReturnCarRequest struct {
	KilometersDriven int `json:"kilometersDriven" validate:"required"`
}
