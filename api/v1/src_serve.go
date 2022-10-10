package v1

import (
	"griffin-dao/gcrud"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "griffin-dao/docs"
)

type GriffinWS struct {
	Conn     *gin.Engine
	Database gcrud.GriffinWeb2Conn
}

type GriffinWSS interface {
	// Landing page operations
	StartService() GriffinWS
	PingTest() GriffinWS
	Version() GriffinWS
	// Employ Type
	AddEmployType() GriffinWS
	DeleteEmployType() GriffinWS
	GetEmployType() GriffinWS
	// Employer CRUD Op
	GetEmployer() GriffinWS
	AddEmployer() GriffinWS
	DeleteEmployer() GriffinWS
	UpdateEmployer() GriffinWS
	// Employee CRD Op
	AddEmployee() GriffinWS
	DeleteEmployee() GriffinWS
	GetEmployeeSingle() GriffinWS
	GetEmployeeMulti() GriffinWS
	// Price and payment history
	GetPrice() GriffinWS
}

func (g GriffinWS) StartService() GriffinWS {
	// Initiate Web2 Server Instance
	//  Swag API documentation page
	//  Attach CORS control middleware
	c := gin.Default()
	c.Use(CORSMiddleware())
	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.Conn = c
	// Initiate Griffin RDB Conn
	var griffinDb gcrud.GriffinWeb2Conn
	if conn, err := gcrud.NewDao(); err == nil {
		griffinDb = conn.Conn()
	}

	g.Database = griffinDb
	return g
}

// PingTest
// @Summary Server, DB ping test
// @Description Check 1) server is alive 2) database is alive.
// @Description Database ping using internal sql method in golang
// @Accept  json
// @Produce  json
// @Router /ping [get]
// @Success 200 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) PingTest() GriffinWS {
	g.Conn.GET("/ping", pingPong)
	return g
}

// Version
// @Summary Read version file from environment file.
// @Description env file's parameter is GRIFFIN_WS_VERSION
// @Accept  json
// @Produce  json
// @Router /version [get]
// @Success 200 {object} CommonResponse
func (g GriffinWS) Version() GriffinWS {
	g.Conn.GET("/version", version)
	return g
}

// AddEmployType
// @Summary Add employee type toe the database
// @Description Employee type needs empType and empMonth.
// @Description e
// @Accept  json
// @Produce  json
// @Param empType query string true "employee type - whether it's permanent or not"
// @Param empMonth query string true "employee contract period in month. -1 if permanent"
// @Router /employType [get]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) AddEmployType() GriffinWS {
	g.Conn.POST("/employType", func(c *gin.Context) {
		addEmpType(c, g.Database)
	})
	return g
}

// DeleteEmployType
func (g GriffinWS) DeleteEmployType() GriffinWS {
	g.Conn.DELETE("/employType", func(c *gin.Context) {
		delEmpType(c, g.Database)
	})
	return g
}

func (g GriffinWS) GetEmployType() GriffinWS {
	g.Conn.GET("/employType", func(c *gin.Context) {
		getEmpType(c, g.Database)
	})
	return g
}

func (g GriffinWS) GetEmployer() GriffinWS {
	g.Conn.GET("/employer", func(c *gin.Context) {
		getEmployer(c, g.Database)
	})
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

func (g GriffinWS) GetEmployeeSingle() GriffinWS {
	g.Conn.GET("/employee/single", func(c *gin.Context) {
		getEmployeeSingle(c, g.Database)
	})
	return g
}

func (g GriffinWS) GetEmployeeMulti() GriffinWS {
	g.Conn.GET("/employee/multi", func(c *gin.Context) {
		getEmployeeMulti(c, g.Database)
	})
	return g
}

func (g GriffinWS) GetEmployeeByType() GriffinWS {
	g.Conn.GET("/employee/multi/type", func(c *gin.Context) {
		getEmployeeMultiWYType(c, g.Database)
	})
	return g
}

func (g GriffinWS) GetPrice() GriffinWS {
	g.Conn.GET("/price", getBinanceTrade)
	return g
}

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
