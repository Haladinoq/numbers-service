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
            "name": "Numbers Service Support"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/reservation": {
            "get": {
                "description": "get reservation numbers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation V1"
                ],
                "summary": "get reservation numbers",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.NumbersResponse"
                        }
                    },
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            },
            "post": {
                "description": "create numbers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reservation V1"
                ],
                "summary": "service for reservation numbers",
                "parameters": [
                    {
                        "description": "The reservation data",
                        "name": "numbers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.NumbersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "401": {
                        "description": "Unauthorized Request"
                    },
                    "404": {
                        "description": "Not Found"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.NumbersRequest": {
            "type": "object",
            "properties": {
                "client": {
                    "type": "string"
                },
                "number": {
                    "type": "integer"
                }
            }
        },
        "model.NumbersResponse": {
            "type": "object",
            "properties": {
                "CreatedAt": {
                    "type": "integer"
                },
                "UpdatedAt": {
                    "type": "integer"
                },
                "client": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "number": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Numbers Service",
	Description:      "This is a Numbers Service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
