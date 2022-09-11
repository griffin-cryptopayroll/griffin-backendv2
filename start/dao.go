package start

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

const MySQL = "mysql"

type GriffinDataAccess struct {
	HostAddress string
	PortAddress string
	Username    string
	Password    string
	Name        string
}

func init() {
	option := flag.String("state", "local", "local state or dev state")
	flag.Parse()

	var fileName string
	switch {
	case *option == "local":
		fmt.Println("develop in local env")
		fileName = "./.env.local"
	case *option == "dev":
		fmt.Println("deploy env")
		fileName = "./.env.production"
	default:
		log.Panicln("state your correct dev state")
	}
	godotenv.Load(fileName)
}

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
