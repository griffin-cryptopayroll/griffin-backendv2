package cache

import (
	"fmt"
	"github.com/gin-contrib/sessions/redis"
	"log"
	"os"
)

type CacheAccessKey struct {
	Address     string
	Port        string
	Password    string
	PrvKey      string
	SessionName string
}

type CacheAccess interface {
	New() CacheAccess
}

func (key *CacheAccessKey) New() *CacheAccessKey {
	key.Address = os.Getenv("CACHE_ADDR")
	key.Port = os.Getenv("CACHE_PORT")
	key.Password = os.Getenv("CACHE_PW")
	key.PrvKey = os.Getenv("CACHE_PK")
	key.SessionName = os.Getenv("griffinsession")
	return key
}

func CacheDatabase() redis.Store {
	sessionCfg := CacheAccessKey{}
	sessionCfg.New()

	store, err := redis.NewStore(
		1000,
		"tcp",
		fmt.Sprintf(
			"%s:%v",
			sessionCfg.Address,
			sessionCfg.Port,
		),
		sessionCfg.Password,
		[]byte(sessionCfg.PrvKey),
	)
	if err != nil {
		log.Panicln("failed to connect to session-cache database")
	}
	return store
}
