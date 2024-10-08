package dao

import (
	"context"
	"fmt"
	"griffin-dao/ent"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func NewDao() (GriffinDataAccess, error) {
	g := GriffinDataAccess{
		HostAddress: os.Getenv("DATABASE_ADDR"),
		PortAddress: os.Getenv("DATABASE_PORT"),
		Username:    os.Getenv("USERNAME"),
		Password:    os.Getenv("PASSWORD"),
		Name:        os.Getenv("NAME"),
	}

	return g, nil
}

func (g GriffinDataAccess) String() string {
	engine := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", g.Username, g.Password, g.HostAddress, g.PortAddress, g.Name)
	return engine
}

func (g GriffinDataAccess) Conn() GriffinWeb2Conn {
	client, err := ent.Open(MySQL, g.String())
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	// Run the auto migration tool
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return GriffinWeb2Conn{
		Conn: client,
	}
}
