package v0

import (
	"griffin-dao/api/common"
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

var griffinDb gcrud.GriffinWeb2Conn

// WebServerStartUp
// 1) Initiate Web2 Server Instance
// 2) Swag API documentation page
// 3) Attach CORS control middleware
// 4) Attach JWT Token Service
// 5) Initiate Griffin RDB Connection
func WebServerStartUp() GriffinWS {
	router := gin.Default()
	router.Use(common.CORSMiddleware())
	router.GET("/api/v0/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if conn, err := gcrud.NewDao(); err == nil {
		griffinDb = conn.Conn()
	}
	srv := GriffinWS{
		Conn:     router,
		Database: griffinDb,
	}
	return srv
}

func (g GriffinWS) StartServiceAuth() GriffinWS {
	g.Conn.Use(common.JwtAuthMiddleware())
	g.Conn.GET("/jwtping", pingPong)
	return g
}

func (g GriffinWS) InitializeApiCommon() GriffinWS {
	v := g.Conn.Group("/api/v0")
	v.GET("/ping", pingPong)
	v.GET("/version", version)
	v.GET("/login", func(c *gin.Context) {
		login(c, g.Database)
	})
	v.GET("/price", getBinanceTrade)
	return g
}

func (g GriffinWS) InitializeApiV0() GriffinWS {
	v0 := g.Conn.Group("/api/v0")

	// Employ Type Web CRUD operation
	v0.POST("/employType", func(c *gin.Context) {
		addEmpType(c, g.Database)
	})
	v0.GET("/employType", func(c *gin.Context) {
		getEmpType(c, g.Database)
	})
	v0.DELETE("/employType", func(context *gin.Context) {
		delEmpType(context, g.Database)
	})

	// Employer Web CRUD operation
	v0.POST("/employer", func(c *gin.Context) {
		addEmployer(c, g.Database)
	})
	v0.GET("/employer", func(c *gin.Context) {
		EmployerWithGid(c, g.Database)
	})
	v0.DELETE("/employer", func(c *gin.Context) {
		delEmployer(c, g.Database)
	})
	v0.PUT("/employer", func(c *gin.Context) {
		updEmployer(c, g.Database)
	})

	// Employee Web CRUD operation
	v0.POST("/employee", func(c *gin.Context) {
		GenerateEmployee(c, g.Database)
	})
	v0.GET("/employee/single", func(c *gin.Context) {
		EmployeeSingle(c, g.Database)
	})
	v0.GET("/employee/multi", func(c *gin.Context) {
		EmployeeMulti(c, g.Database)
	})
	v0.DELETE("/employee", func(c *gin.Context) {
		RemoveEmployee(c, g.Database)
	})

	// Employee Payment Record - Align it with Web3 API calls
	v0.PUT("/payment/execute", func(c *gin.Context) {
		PaymentExecute(c, g.Database)
	})
	v0.POST("/payment/oneoff", func(c *gin.Context) {
		PaymentOneOff(c, g.Database)
	})
	v0.GET("/payment/employee", func(c *gin.Context) {
		PaymentByEmployee(c, g.Database)
	})
	v0.GET("/payment/employer", func(c *gin.Context) {
		PaymentByEmployer(c, g.Database)
	})
	return g
}
