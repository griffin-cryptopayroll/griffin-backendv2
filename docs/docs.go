// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Sang Il Bae",
            "email": "baesangil0906@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/employType": {
            "get": {
                "description": "Employee type needs empMonth.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query employee type from the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "employee contract period in month. -1 if permanent",
                        "name": "empMonth",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gcrud.EmployerJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Employee type needs empType and empMonth.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Add employee type toe the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "employee type - whether it's permanent or not",
                        "name": "empType",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "employee contract period in month. -1 if permanent",
                        "name": "empMonth",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Employee type needs empMonth.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete employee type from the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "employee contract period in month. -1 if permanent",
                        "name": "empMonth",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            }
        },
        "/employee": {
            "post": {
                "description": "Worker's information needed.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Post employee to the database.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee's griffin id (in uuid form)",
                        "name": "gid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Last name",
                        "name": "last_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "First name",
                        "name": "first_name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Position ex) Backend engineer, Frontend engineer",
                        "name": "position",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employee's information. His or her payment wallet address",
                        "name": "wallet",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "number",
                        "description": "Payroll amount in float",
                        "name": "payroll",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "ID (integer) of the payroll currency",
                        "name": "currency",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Employee's information. Corp or organization's em",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employee's information. Corp Gid or Organization Gid",
                        "name": "employer_gid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employee's information. When does he or she starts work. In YYYYMMDD",
                        "name": "work_start",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Worker's information needed. His/Her Griffin ID (GID) and employer Griffin ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete employee from the database.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee's griffin id (in uuid form)",
                        "name": "gid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employee's information. Corp Gid or Organization Gid",
                        "name": "employer_gid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            }
        },
        "/employee/multi": {
            "get": {
                "description": "Worker's information needed.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query employee from the database.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee's information. Corp Gid or Organization Gid",
                        "name": "employer_gid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/gcrud.EmployeeJson"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            }
        },
        "/employee/multi/type": {
            "get": {
                "description": "Worker's information needed.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Post employee to the database.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee's information. Corp Gid or Organization Gid",
                        "name": "employer_gid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employee's information. Contract month. -1 if permanent",
                        "name": "contract_month",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/gcrud.EmployeeJson"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            }
        },
        "/employee/single": {
            "get": {
                "description": "Worker's information needed. Worker is singled out with their griffin id and his employer id.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query employee from the database.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee's griffin id (in uuid form)",
                        "name": "gid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employee's information. Corp Gid or Organization Gid",
                        "name": "employer_gid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gcrud.EmployeeJson"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            }
        },
        "/employer": {
            "get": {
                "description": "Employer is registered by google form.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Query employer from the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employer's griffin id (in uuid form)",
                        "name": "gid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Employer information is registered by google form.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Post employer to the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employer's griffin id (in uuid form)",
                        "name": "gid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employer's user id",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employer's user password.",
                        "name": "pw",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Employer information (corp or organization name)",
                        "name": "corp_name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Employer information (corp or organization email)",
                        "name": "corp_email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Employer's griffin id (in uuid form)",
                        "name": "wallet",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "-",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Delete employer from the database",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employer's griffin id (in uuid form)",
                        "name": "gid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            }
        },
        "/ping": {
            "get": {
                "description": "Check 1) server is alive 2) database is alive.\nDatabase ping using internal sql method in golang",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Server, DB ping test",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            }
        },
        "/price": {
            "get": {
                "description": "ETH, MATIC data from binance.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Get all the price information that Griffin serves",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/price.PriceInformation"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "env file's parameter is GRIFFIN_WS_VERSION",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Read version file from environment file.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/v1.CommonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gcrud.EmployeeJson": {
            "type": "object",
            "required": [
                "email",
                "employ_type",
                "employer_gid",
                "first_name",
                "last_name",
                "payday",
                "position",
                "wallet",
                "work_start"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "currency": {
                    "description": "get it from currency table",
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "employ_type": {
                    "description": "get it from employ type",
                    "type": "integer"
                },
                "employer_gid": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "gid": {
                    "type": "string"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "last_name": {
                    "type": "string"
                },
                "payday": {
                    "type": "string"
                },
                "payroll": {
                    "type": "number"
                },
                "position": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                },
                "wallet": {
                    "type": "string"
                },
                "work_start": {
                    "type": "string",
                    "example": "20220701"
                }
            }
        },
        "gcrud.EmployerJson": {
            "type": "object",
            "properties": {
                "corp_email": {
                    "type": "string"
                },
                "corp_name": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "created_by": {
                    "type": "string"
                },
                "gid": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "updated_by": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                },
                "wallet": {
                    "type": "string"
                }
            }
        },
        "price.PriceInformation": {
            "type": "object",
            "properties": {
                "ethusdt": {
                    "type": "number"
                },
                "maticusdt": {
                    "type": "number"
                },
                "usdcusdt": {
                    "type": "number"
                }
            }
        },
        "v1.CommonResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "database (create / delete) (successful / failed)"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "Document 1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Griffin Web Server API Documentation",
	Description:      "Griffin webserver that serves, employee .",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
