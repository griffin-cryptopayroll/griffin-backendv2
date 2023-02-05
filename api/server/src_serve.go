package server

import (
	"griffin-dao/api/api_v0"
	"griffin-dao/api/common"
	"griffin-dao/api/login"
	"griffin-dao/dao"
	"griffin-dao/gsession"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "griffin-dao/docs"
)

type GriffinWS struct {
	Conn      *gin.Engine
	Database  dao.GriffinWeb2Conn
	SessionDB redis.Store
}

var griffinDb dao.GriffinWeb2Conn

// WebServerStartUp
// 1) Initiate Web2 Server Instance
// 2) Swag API documentation page
// 3) Attach CORS control middleware
// 4) Attach Login service + Session service
// 5) Initiate Griffin RDB Connection (MySQL)
func WebServerStartUp() GriffinWS {
	// 1)
	router := gin.Default()

	// Session Database 4)
	sessionDb := gsession.CacheDatabase()

	// Griffin Worker RDB 5)
	if conn, err := dao.NewDao(); err == nil {
		griffinDb = conn.Conn()
	}

	// 2), 3), 4)
	router.Use(common.CORSMiddleware())
	router.GET("/api/v0/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(sessions.Sessions(os.Getenv("CACHE_SESSION_NAME"), sessionDb))

	// Serve
	srv := GriffinWS{
		Conn:      router,
		Database:  griffinDb,
		SessionDB: sessionDb,
	}
	return srv
}

func (g GriffinWS) InitializeApiCommon() GriffinWS {
	g.Conn.GET("/ping", pingPong)
	g.Conn.GET("/version", version)
	return g
}

func (g GriffinWS) InitializeApiV0() GriffinWS {
	v0 := g.Conn.Group("/api/v0")
	// Pricing
	v0.GET("/price", api_v0.GetBinanceTrade)

	// Employ Type Web CRUD operation
	v0.POST("/employType", func(c *gin.Context) {
		api_v0.AddEmpType(c, g.Database)
	})
	v0.GET("/employType", func(c *gin.Context) {
		api_v0.GetEmpType(c, g.Database)
	})
	v0.DELETE("/employType", func(context *gin.Context) {
		api_v0.DelEmpType(context, g.Database)
	})

	// Employer Web CRUD operation
	v0.POST("/employer", func(c *gin.Context) {
		api_v0.AddEmployer(c, g.Database)
	})
	v0.GET("/employer", func(c *gin.Context) {
		api_v0.EmployerWithGid(c, g.Database)
	})
	v0.DELETE("/employer", func(c *gin.Context) {
		api_v0.DelEmployer(c, g.Database)
	})
	v0.PUT("/employer", func(c *gin.Context) {
		api_v0.UpdEmployer(c, g.Database)
	})

	// Employee Web CRUD operation
	v0.POST("/employee", func(c *gin.Context) {
		api_v0.GenerateEmployee(c, g.Database)
	})
	v0.GET("/employee/single", func(c *gin.Context) {
		api_v0.EmployeeSingle(c, g.Database)
	})
	v0.GET("/employee/multi", func(c *gin.Context) {
		api_v0.EmployeeMulti(c, g.Database)
	})
	v0.DELETE("/employee", func(c *gin.Context) {
		api_v0.RemoveEmployee(c, g.Database)
	})

	// Employee Payment Record - Align it with Web3 API calls
	v0.PUT("/payment/execute", func(c *gin.Context) {
		api_v0.PaymentExecute(c, g.Database)
	})
	v0.POST("/payment/oneoff", func(c *gin.Context) {
		api_v0.PaymentOneOff(c, g.Database)
	})
	v0.GET("/payment/employee", func(c *gin.Context) {
		api_v0.PaymentByEmployee(c, g.Database)
	})
	v0.GET("/payment/employer", func(c *gin.Context) {
		api_v0.PaymentByEmployer(c, g.Database)
	})
	v0.GET("/payment/future", func(c *gin.Context) {
		api_v0.PaymentFuture(c, g.Database)
	})
	v0.GET("/payment/past", func(c *gin.Context) {
		api_v0.PaymentPast(c, g.Database)
	})
	v0.GET("/payment/miss", func(c *gin.Context) {
		api_v0.PaymentMissed(c, g.Database)
	})
	return g
}

func (g GriffinWS) InitializeLoginV1() GriffinWS {
	v1 := g.Conn.Group("/api/v1")

	v1.GET("/nonce", api_login.SiweNonce)
	v1.POST("/verify", func(c *gin.Context) {
		api_login.SiweVerify(c, g.Database)
	})
	return g
}

func (g GriffinWS) SessionUsage() GriffinWS {
	u1 := g.Conn.Group("/api/test")
	u1.Use(common.SessionAuthMiddleware())
	u1.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "testing point",
		})
	})
	return g
}
