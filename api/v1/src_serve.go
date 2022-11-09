package v1

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

type GriffinWSS interface {
	// Landing page operations
	StartService() GriffinWS
	PingTest() GriffinWS
	Version() GriffinWS
	Login() GriffinWS
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
	GetMetrics() GriffinWS
}

func (g GriffinWS) StartService() GriffinWS {
	// Initiate Web2 Server Instance
	//  Swag API documentation page
	//  Attach CORS control middleware
	//  Attach JWT Token Service
	c := gin.Default()
	c.Use(common.CORSMiddleware())

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

func (g GriffinWS) StartServiceWithAuth() GriffinWS {
	g.Conn.Use(common.JwtAuthMiddleware())
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

func (g GriffinWS) JWTTest() GriffinWS {
	g.Conn.GET("/jwtping", pingPong)
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

// Login
// @Summary Login into griffin payroll service
// @Description Matches Username with Password. If Username does not exists, 500 error.
// @Accept  json
// @Produce  json
// @Param username query string true "Employer's username (in email form)"
// @Param password query string true "Employer's password"
// @Router /login [get]
// @Success 200 {object} CommonResponseToken
// @Failure 400 {object} CommonResponse
// @Failure 403 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) Login() GriffinWS {
	g.Conn.GET("/login", func(c *gin.Context) {
		login(c, g.Database)
	})
	return g
}

// AddEmployType
// @Summary Add Employ type.
// @Description Whether the employer is full-time worker(fulltime) or contract worker(contract)
// @Accept  json
// @Produce  json
// @Param isPerma query string true "Enumerator (fulltime or contract)"
// @Param payFreq query string true "Single formed frequency, such as D(Day), W(Week), M(Month), Y(Year)"
// @Router /employType [post]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 403 {object} CommonResponse
// @Failrue 500 {object} CommonResponse
func (g GriffinWS) AddEmployType() GriffinWS {
	g.Conn.POST("/employType", func(c *gin.Context) {
		addEmpType(c, g.Database)
	})
	return g
}

// DeleteEmployType
// @Summary Delete Employ type.
// @Description Not yet implemented
// @Accept json
// @Produce json
// @Router /employType [delete]
// @Failure 403 {object} CommonResponse
func (g GriffinWS) DeleteEmployType() GriffinWS {
	g.Conn.DELETE("/employType", func(c *gin.Context) {
		delEmpType(c, g.Database)
	})
	return g
}

// GetEmployType
// @Summary Get Employ type.
// @Description Whether the employer is full-time worker(fulltime) or contract worker(contract)
// @Accept  json
// @Produce  json
// @Param isPerma query string true "Enumerator (fulltime or contract)"
// @Param payFreq query string true "Single formed frequency, such as D(Day), W(Week), M(Month), Y(Year)"
// @Router /employType [get]
// @Success 200 {object} ent.EMPLOY_TYPE
// @Failure 400 {object} CommonResponse
// @Failure 403 {object} CommonResponse
// @Failrue 500 {object} CommonResponse
func (g GriffinWS) GetEmployType() GriffinWS {
	g.Conn.GET("/employType", func(c *gin.Context) {
		getEmpType(c, g.Database)
	})
	return g
}

// GetEmployer
// @Summary Query employer from the database
// @Description Employer Griffin ID is in UUID form. Login gives you access to UUID.
// @Accept  json
// @Produce  json
// @Param gid query string true "Employer's griffin id (in uuid form)"
// @Router /employer [get]
// @Success 200 {object} gcrud.EmployerJson
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) GetEmployer() GriffinWS {
	g.Conn.GET("/employer", func(c *gin.Context) {
		getEmployerwGid(c, g.Database)
	})
	return g
}

// AddEmployer
// @Summary Post employer to the database
// @Description Employer information is registered by google form.
// @Accept  json
// @Produce  json
// @Param id query string false "Employer's user id"
// @Param pw query string true "Employer's user password."
// @Param corp_name query string false "Employer information (corp or organization name)"
// @Param corp_email query string true "Employer information (corp or organization email)"
// @Param wallet query string false "Employer's griffin id (in uuid form)"
// @Router /employer [post]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) AddEmployer() GriffinWS {
	g.Conn.POST("/employer", func(c *gin.Context) {
		addEmployer(c, g.Database)
	})
	return g
}

// DeleteEmployer
// @Summary Delete employer to the database
// @Description Deleting the employer will delete all the employee's related to that employer
// @Accept  json
// @Produce  json
// @Param gid query string true "Employer's griffin id (in uuid form)"
// @Router /employer [delete]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) DeleteEmployer() GriffinWS {
	g.Conn.DELETE("/employer", func(c *gin.Context) {
		delEmployer(c, g.Database)
	})
	return g
}

// UpdateEmployer
// @Summary Delete Employer.
// @Description Not yet implemented
// @Accept json
// @Produce json
// @Router /employer [patch]
// @Failure 403 {object} CommonResponse
func (g GriffinWS) UpdateEmployer() GriffinWS {
	g.Conn.PATCH("/employer", func(c *gin.Context) {
		updEmployer(c, g.Database)
	})
	return g
}

// AddEmployee
// @Summary Post employee to the database.
// @Description Worker's information needed.
// @Accept  json
// @Produce  json
// @Param name query string true "Full name, since crypto lovers don't use their original name"
// @Param employ_type query string true "permanent or freelance"
// @Param pay_freq query string true "D, W, or M"
// @Param position query string false "Position ex: Backend engineer, Frontend engineer"
// @Param wallet query string true "Employee's information. His or her payment wallet address"
// @Param payroll query float32 true "Payroll amount in float"
// @Param currency query int true "ID (integer) of the payroll currency"
// @Param email query string true "Employee's information. Corp or organization's em"
// @Param payday query time.Time true "Employee's information. Payday information"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Param work_start query string true "Employee's information. When does he or she starts work. In YYYYMMDD"
// @Param work_end query string false "Employee's information. When does he or she ends work. In YYYYMMDD"
// @Router /employee [post]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) AddEmployee() GriffinWS {
	g.Conn.POST("/employee", func(c *gin.Context) {
		addEmployee(c, g.Database)
	})
	return g
}

// DeleteEmployee
// @Summary Delete employee from the database.
// @Description Worker's information needed. His/Her Griffin ID (GID) and employer Griffin ID.
// @Accept  json
// @Produce  json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /employee [delete]
// @Success 200 {object} CommonResponse
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) DeleteEmployee() GriffinWS {
	g.Conn.DELETE("/employee", func(c *gin.Context) {
		delEmployee(c, g.Database)
	})
	return g
}

// GetEmployeeSingle
// @Summary Query employee from the database.
// @Description Worker's information needed. Worker is singled out with their griffin id and his employer id.
// @Accept  json
// @Produce  json
// @Param gid query string true "Employee's griffin id (in uuid form)"
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /employee/single [get]
// @Success 200 {object} gcrud.EmployeeJson
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) GetEmployeeSingle() GriffinWS {
	g.Conn.GET("/employee/single", func(c *gin.Context) {
		getEmployeeSingle(c, g.Database)
	})
	return g
}

// GetEmployeeMulti
// @Summary Query employee from the database.
// @Description Worker's information needed.
// @Accept  json
// @Produce  json
// @Param employer_gid query string true "Employee's information. Corp Gid or Organization Gid"
// @Router /employee/multi [get]
// @Success 200 {object} []gcrud.EmployeeJson
// @Failure 400 {object} CommonResponse
// @Failure 500 {object} CommonResponse
func (g GriffinWS) GetEmployeeMulti() GriffinWS {
	g.Conn.GET("/employee/multi", func(c *gin.Context) {
		getEmployeeMulti(c, g.Database)
	})
	return g
}

// GetPrice
// @Summary Get all the price information that Griffin serves
// @Description ETH, MATIC data from binance.
// @Accept  json
// @Produce  json
// @Router /price [get]
// @Success 200 {object} price.PriceInformation
// @Failure 400 {object} CommonResponse
func (g GriffinWS) GetPrice() GriffinWS {
	g.Conn.GET("/price", getBinanceTrade)
	return g
}

func (g GriffinWS) GetMetrics() GriffinWS {
	g.Conn.GET("/metrics/month", func(c *gin.Context) {

	})
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
