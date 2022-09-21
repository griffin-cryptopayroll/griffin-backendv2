package api

import (
	"griffin-dao/gcrud"

	"github.com/gin-gonic/gin"
)

type GriffinWS struct {
	Conn     *gin.Engine
	Database gcrud.GriffinWeb2Conn
}

func (g GriffinWS) StartService() GriffinWS {
	// Initiate Web2 Server Instance
	c := gin.Default()
	c.Use(CORSMiddleware())
	g.Conn = c
	// Initiate Griffin RDB Conn
	var griffinDb gcrud.GriffinWeb2Conn
	if conn, err := gcrud.NewDao(); err == nil {
		griffinDb = conn.Conn()
	}
	g.Database = griffinDb
	return g
}

func (g GriffinWS) PingTest() GriffinWS {
	g.Conn.GET("/ping", pingPong)
	return g
}

func (g GriffinWS) Version() GriffinWS {
	g.Conn.GET("/version", version)
	return g
}

func (g GriffinWS) AddEmployer() GriffinWS {
	g.Conn.POST("/employer", func(c *gin.Context) {
		addEmployer(c, g.Database)
	})
	return g
}

func (g GriffinWS) DeleteEmployer() GriffinWS {
	g.Conn.DELETE("/employer", func(c *gin.Context) {
		delEmployer(c, g.Database)
	})
	return g
}

func (g GriffinWS) UpdateEmployer() GriffinWS {
	g.Conn.PATCH("/employer", func(c *gin.Context) {
		updEmployer(c, g.Database)
	})
	return g
}

func (g GriffinWS) AddEmployee() GriffinWS {
	g.Conn.POST("/employee", func(c *gin.Context) {
		addEmployee(c, g.Database)
	})
	return g
}

func (g GriffinWS) DeleteEmployee() GriffinWS {
	g.Conn.DELETE("/employee", func(c *gin.Context) {
		delEmployee(c, g.Database)
	})
	return g
}

//
//func (g GriffinWS) GetEmployee() GriffinWS {
//	g.Conn.GET("/employee", func(c *gin.Context) {
//		getEmployee(c, g.Database)
//	})
//	return g
//}
//

//
//func (g GriffinWS) GetPrice() GriffinWS {
//	g.Conn.GET("/price", getBinanceTrade)
//	return g
//}
//
//func (g GriffinWS) AddPaymentRecord() GriffinWS {
//	g.Conn.POST("/payment", func(c *gin.Context) {
//		postPayment(c, g.Database)
//	})
//	return g
//}
//
//func (g GriffinWS) GetPaymentRecord() GriffinWS {
//	g.Conn.GET("/payment", func(c *gin.Context) {
//		getPayment(c, g.Database)
//	})
//	return g
//}
//
//func (g GriffinWS) GetPaymentRecordMonth() GriffinWS {
//	g.Conn.GET("/paymentMonth", func(c *gin.Context) {
//		getPaymentMonthly(c, g.Database)
//	})
//	return g
//}
//
//func (g GriffinWS) Login() GriffinWS {
//	g.Conn.GET("/login", func(c *gin.Context) {
//		loginEmployer(c, g.Database)
//	})
//	return g
//}
