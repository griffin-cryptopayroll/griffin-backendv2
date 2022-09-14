package start

import "griffin-dao/ent"

const MySQL = "mysql"

type GriffinDataAccess struct {
	HostAddress string
	PortAddress string
	Username    string
	Password    string
	Name        string
}

type DataAccess interface {
	String() string
	Conn() GriffinWeb2Conn
}

type GriffinWeb2Conn struct {
	Conn *ent.Client
}
