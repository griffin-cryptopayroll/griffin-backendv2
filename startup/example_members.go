package startup

import (
	"griffin-dao/dao"
	"griffin-dao/ent"
)

func employerStartUp(db dao.GriffinWeb2Conn) error {
	myEmployer := ent.EMPLOYER{
		Password:  "123",
		CorpName:  "Griffin DAO",
		CorpEmail: "griffindao@griffin.xyz",
		Wallet:    "0x0it0is0my0wallet0",
	}
	err := dao.CreateEmployer(myEmployer, ctx, db.Conn)
	if err != nil {
		return err
	}
	return nil
}
