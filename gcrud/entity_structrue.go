package gcrud

import "time"

type EmployeeJson struct {
	ID                int       `json:"id" example:"1"`
	GriffinID         string    `json:"gid"`
	EmployerGriffinID string    `json:"employer_gid"`
	Name              string    `json:"name" binding:"required"`
	Position          string    `json:"position" binding:"required"`
	Wallet            string    `json:"wallet" binding:"required"`
	Payroll           float64   `json:"payroll"`
	Currency          int       `json:"currency"` // foreign key currency table
	PayDay            time.Time `json:"payday" binding:"required"`
	EmployType        string    `json:"employ_type" binding:"required" example:"permanent"` // foreign key employ type
	Email             string    `json:"email" binding:"required,email" example:"example@ex.com"`
	WorkStart         string    `json:"work_start" binding:"required" example:"20220701"`
	WorkEnd           string    `json:"work_end" binding:"required" example:"20221231"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         string    `json:"created_by"`
	UpdatedAt         time.Time `json:"updated_at"`
	UpdatedBy         string    `json:"updated_by"`
}

type EmployerJson struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	GriffinID string    `json:"gid"`
	CorpName  string    `json:"corp_name"`
	CorpEmail string    `json:"corp_email"`
	Wallet    string    `json:"wallet"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

type PaymentHistory struct {
	ID                int       `json:"id"`
	EmployeeGriffinID string    `json:"employee_gid"`
	CreatedAt         time.Time `json:"created_at"`
	CreatedBy         string    `json:"created_by"`
}
