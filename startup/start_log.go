package startup

import (
	"context"
	"griffin-dao/dao"
	"griffin-dao/ent"
)

func isLogStartUp(db dao.GriffinWeb2Conn) bool {
	logger, err := dao.QueryTrLog("init", context.Background(), db.Conn)
	// Cannot find log or no logs
	if err != nil || len(logger) == 0 {
		return false
	}
	return true
}

func logStartUp(db dao.GriffinWeb2Conn) error {
	startUpLog := ent.Tr_log{
		TrType: "init",
	}
	err := dao.CreateTrLog(startUpLog, context.Background(), db.Conn)
	if err != nil {
		return err
	}
	return nil
}
