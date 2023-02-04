package startup

import (
	"griffin-dao/dao"
	"griffin-dao/util"
)

func ExecStartUp(client dao.GriffinWeb2Conn) {
	hasInit := isLogStartUp(client)

	if hasInit {
		// already has a start-up file
		util.PrintGreenStatus("Database already pre-setted up!")
		return
	}
	// Generate and log start up file
	util.PrintPurpleWarning("Initiate database with pre-set data")
	_ = empTypeStartUp(client)
	_ = sourceStartUp(client)
	_ = currencyStartUp(client)
	_ = employerStartUp(client)
	_ = logStartUp(client)
}
