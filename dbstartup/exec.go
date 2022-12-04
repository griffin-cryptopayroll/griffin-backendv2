package dbstartup

import (
	"griffin-dao/gcrud"
	"griffin-dao/service"
)

func ExecStartUp(client gcrud.GriffinWeb2Conn) {
	hasInit := isLogStartUp(client)

	if hasInit {
		// already has a start-up file
		service.PrintGreenStatus("Database already pre-setted up!")
		return
	}
	// Generate and log start up file
	service.PrintPurpleWarning("Initiate database with pre-set data")
	_ = empTypeStartUp(client)
	_ = sourceStartUp(client)
	_ = currencyStartUp(client)
	_ = employerStartUp(client)
	_ = logStartUp(client)
}
