// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/buy": {
            "post": {
                "description": "Buy value",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exchange"
                ],
                "summary": "Buy",
                "parameters": [
                    {
                        "description": "balance for buy",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.BuyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.BuyResponse"
                        }
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/sell": {
            "post": {
                "description": "Sell value",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exchange"
                ],
                "summary": "Sell",
                "parameters": [
                    {
                        "description": "balance for sell",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.SellRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.SellResponse"
                        }
                    },
                    "400": {
                        "description": "invalid request",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal error",
                        "schema": {
                            "$ref": "#/definitions/dto.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.BuyRequest": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "integer"
                }
            }
        },
        "dto.BuyResponse": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "integer"
                }
            }
        },
        "dto.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "dto.SellRequest": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "integer"
                }
            }
        },
        "dto.SellResponse": {
            "type": "object",
            "properties": {
                "value": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "SideEcho API",
	Description:      "This is a sample Exchange server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}