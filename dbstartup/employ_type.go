package dbstartup

import (
	"context"
	"griffin-dao/gcrud"
)

func empTypeStartUp(db gcrud.GriffinWeb2Conn) {
	types := [][]string{
		{"permanent", "D"},
		{"permanent", "W"},
		{"permanent", "M"},
		{"freelance", "D"},
		{"freelance", "W"},
		{"freelance", "M"},
	}

	for _, r := range types {
		gcrud.CreateEmployType(r[0], r[1], context.Background(), db.Conn)
	}
}
