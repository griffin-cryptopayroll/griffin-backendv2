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
                "summary": "Query employee type to the database",
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
                "summary": "Delete employee type to the database",
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
