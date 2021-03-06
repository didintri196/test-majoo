// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/product": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show all data Product Mercant with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get list data Product In Mercant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some Fiter",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "name|asc",
                        "name": "sort",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "5",
                        "name": "per_page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ListProduct"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "500": {
                        "description": "Something Wrong on Server",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Create Product",
                "parameters": [
                    {
                        "description": "Create Product",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Product"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    },
                    "400": {
                        "description": "Invalid Form",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "500": {
                        "description": "Something Wrong on Server",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    }
                }
            }
        },
        "/product/view/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show by id Product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get By Id Product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductData"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update Product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Update Product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update Product",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.Product"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    },
                    "401": {
                        "description": "Error Not Logged In",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "404": {
                        "description": "ID not found",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    },
                    "406": {
                        "description": "Invalid Form",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete Product by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Delete Product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    },
                    "401": {
                        "description": "Error Not Logged In",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "404": {
                        "description": "ID not found",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    },
                    "406": {
                        "description": "Invalid Form",
                        "schema": {
                            "$ref": "#/definitions/responses.ProductDefault"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show all data User Mercant with pagination",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get list data User Mercant",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some Fiter",
                        "name": "filter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "name|asc",
                        "name": "sort",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "1",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "5",
                        "name": "per_page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.ListUser"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "500": {
                        "description": "Something Wrong on Server",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Create User",
                "parameters": [
                    {
                        "description": "Create User",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    },
                    "400": {
                        "description": "Invalid Form",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "500": {
                        "description": "Something Wrong on Server",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "description": "Login User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Login User",
                "parameters": [
                    {
                        "description": "Login User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthOK"
                        }
                    },
                    "400": {
                        "description": "Invalid Form",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    }
                }
            }
        },
        "/user/session": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show Session User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get By Session User",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserData"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    }
                }
            }
        },
        "/user/view/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Show by id User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get By Id User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserData"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update User by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update User",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserUpdate"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    },
                    "401": {
                        "description": "Error Not Logged In",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "404": {
                        "description": "ID not found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    },
                    "406": {
                        "description": "Invalid Form",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete User by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete User by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Some ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    },
                    "401": {
                        "description": "Error Not Logged In",
                        "schema": {
                            "$ref": "#/definitions/responses.AuthError"
                        }
                    },
                    "404": {
                        "description": "ID not found",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    },
                    "406": {
                        "description": "Invalid Form",
                        "schema": {
                            "$ref": "#/definitions/responses.UserDefault"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "forms.Product": {
            "type": "object",
            "properties": {
                "foto": {
                    "type": "string"
                },
                "harga_beli": {
                    "type": "integer"
                },
                "harga_jual": {
                    "type": "integer"
                },
                "id_barang": {
                    "type": "integer"
                },
                "id_mercant": {
                    "type": "integer"
                },
                "id_supplier": {
                    "type": "integer"
                },
                "is_aktif": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                },
                "sku": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "forms.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id_mercant": {
                    "type": "integer"
                },
                "id_user": {
                    "type": "integer"
                },
                "last_login": {
                    "type": "integer"
                },
                "nama": {
                    "type": "string"
                },
                "no_telp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "request.Product": {
            "type": "object",
            "properties": {
                "foto": {
                    "type": "string"
                },
                "harga_beli": {
                    "type": "integer"
                },
                "harga_jual": {
                    "type": "integer"
                },
                "id_mercant": {
                    "type": "integer"
                },
                "id_supplier": {
                    "type": "integer"
                },
                "is_aktif": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                },
                "sku": {
                    "type": "string"
                },
                "stock": {
                    "type": "integer"
                }
            }
        },
        "request.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id_mercant": {
                    "type": "integer"
                },
                "nama": {
                    "type": "string"
                },
                "no_telp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "request.UserLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "request.UserUpdate": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "nama": {
                    "type": "string"
                },
                "no_telp": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "responses.AuthError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.AuthOK": {
            "type": "object",
            "properties": {
                "expired": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/forms.User"
                }
            }
        },
        "responses.ListProduct": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/forms.Product"
                    }
                },
                "last_page": {
                    "type": "integer"
                },
                "next_page": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "prev_page": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "responses.ListUser": {
            "type": "object",
            "properties": {
                "current_page": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/forms.User"
                    }
                },
                "last_page": {
                    "type": "integer"
                },
                "next_page": {
                    "type": "integer"
                },
                "per_page": {
                    "type": "integer"
                },
                "prev_page": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "responses.ProductData": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/forms.Product"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.ProductDefault": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.UserData": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/forms.User"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "responses.UserDefault": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "/v1",
	Schemes:     []string{},
	Title:       "API MINI-POS",
	Description: "Api untuk service MINI-POS",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
