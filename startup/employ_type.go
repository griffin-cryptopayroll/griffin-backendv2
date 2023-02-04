package startup

import (
	"context"
	"griffin-dao/dao"
	"griffin-dao/util"
)

func empTypeStartUp(db dao.GriffinWeb2Conn) error {
	types := [][]string{
		{"permanent", "D"},
		{"permanent", "W"},
		{"permanent", "M"},
		{"freelance", "D"},
		{"freelance", "W"},
		{"freelance", "M"},
	}

	for _, r := range types {
		err := dao.CreateEmployType(r[0], r[1], context.Background(), db.Conn)
		if err != nil {
			util.PrintRedError(err.Error())
			return err
		}
	}
	return nil
}
