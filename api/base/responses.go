package api_base

import "griffin-dao/ent"

type CommonResponse struct {
	Status  bool   `json:"status,omitempty" example:"true"`
	Message string `json:"message,omitempty" example:"database (create / delete) (successful / failed)"`
}

type PaymentType struct {
	Scheduled []*ent.PAYMENT `json:"scheduled"`
	Executed  []*ent.PAYMENT `json:"executed"`
	OneOff    []*ent.PAYMENT `json:"oneoff"`
}

type PaymentTime struct {
	Future []*ent.PAYMENT `json:"future,omitempty"`
	Past   []*ent.PAYMENT `json:"past,omitempty"`
	Missed []*ent.PAYMENT `json:"missed,omitempty"`
}
