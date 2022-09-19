package gcrud

import "time"

type EmployeeJson struct {
	ID                int       `json:"id"`
	GriffinID         string    `json:"gid"`
	EmployerGriffinID string    `json:"employer_gid"`
	LastName          string    `json:"last_name"`
	FirstName         string    `json:"first_name"`
	Position          string    `json:"position"`
	Wallet            string    `json:"wallet"`
	Payroll           float64   `json:"payroll"`
	Currency          int       `json:"currency"` // get it from currency table
	PayDay            time.Time `json:"payday"`
	EmployType        int       `json:"employ_type"` // get it from employ type
	Email             string    `json:"email"`
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
