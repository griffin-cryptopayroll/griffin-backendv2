package dbstartup

import (
	"context"
	"griffin-dao/gcrud"
)

func empTypeStartUp(db gcrud.GriffinWeb2Conn) {
	types := [][]string{
		[]string{"permanent", "D"},
		[]string{"permanent", "W"},
		[]string{"permanent", "M"},
		[]string{"freelance", "D"},
		[]string{"freelance", "W"},
		[]string{"freelance", "M"},
	}

	for _, r := range types {
		gcrud.CreateEmployType(r[0], r[1], context.Background(), db.Conn)
	}
}
