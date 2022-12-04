package dbstartup

import (
	"context"
	"griffin-dao/gcrud"
	"griffin-dao/service"
)

func empTypeStartUp(db gcrud.GriffinWeb2Conn) error {
	types := [][]string{
		{"permanent", "D"},
		{"permanent", "W"},
		{"permanent", "M"},
		{"freelance", "D"},
		{"freelance", "W"},
		{"freelance", "M"},
	}

	for _, r := range types {
		err := gcrud.CreateEmployType(r[0], r[1], context.Background(), db.Conn)
		if err != nil {
			service.PrintRedError(err.Error())
			return err
		}
	}
	return nil
}
