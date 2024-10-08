package api_base

import (
	"griffin-dao/ent"
	"griffin-dao/util"
)

type CommonResponse struct {
	Status  bool   `json:"status,omitempty" example:"true"`
	Message string `json:"message,omitempty" example:"database (create / delete) (successful / failed)"`
}

type CommonResponseData[T any] struct {
	Status        bool   `json:"status,omitempty" example:"true"`
	Message       string `json:"message,omitempty" example:"database(create/delete)"`
	DataContainer T      `json:"data"`
}

type LoginResponse struct {
	Status  bool   `json:"status,omitempty" example:"true"`
	Message string `json:"message,omitempty" example:"login successful"`
	Gid     string `json:"gid,omitempty" example:"some griffin id. gid && sessionid is required to request"`
}

type PaymentType struct {
	Scheduled []*ent.PAYMENT `json:"scheduled"`
	Executed  []*ent.PAYMENT `json:"executed"`
	OneOff    []*ent.PAYMENT `json:"oneoff"`
}

type PaymentTypeV2 struct {
	Scheduled []util.Payment `json:"scheduled"`
	Executed  []util.Payment `json:"executed"`
	OneOff    []util.Payment `json:"oneoff"`
}

type PaymentTime struct {
	Future []*ent.PAYMENT `json:"future,omitempty"`
	Past   []*ent.PAYMENT `json:"past,omitempty"`
	Missed []*ent.PAYMENT `json:"missed,omitempty"`
}

type PaymentTimeV2 struct {
	Future []util.Payment `json:"future,omitempty"`
	Past   []util.Payment `json:"past,omitempty"`
	Missed []util.Payment `json:"missed,omitempty"`
}
