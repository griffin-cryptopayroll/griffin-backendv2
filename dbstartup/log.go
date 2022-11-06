package dbstartup

import (
	"context"
	"griffin-dao/ent"
	"griffin-dao/gcrud"
)

func isLogStartUp(db gcrud.GriffinWeb2Conn) bool {
	logger, err := gcrud.QueryTrLog("init", context.Background(), db.Conn)
	// Cannot find log or no logs
	if err != nil || len(logger) == 0 {
		return false
	}
	return true
}

func logStartUp(db gcrud.GriffinWeb2Conn) {
	startUpLog := ent.Tr_log{
		TrType: "init",
	}
	gcrud.CreateTrLog(startUpLog, context.Background(), db.Conn)
}
