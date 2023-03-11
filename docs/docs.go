// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/client-id-client-secret/": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Get All client_id and client_secret for user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Get All client_id and client_secret for user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.ArrayClientRegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.BadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.UnauthorizedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "description": "login",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "type": "string",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "name": "username",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/auth/register-client": {
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "generate client_id and client_secret for request user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "register client",
                "parameters": [
                    {
                        "description": "register client",
                        "name": "register",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ClientRegiterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/schemas.ClientRegiterResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/schemas.ForbiddenResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/oauth/authorize/": {
            "post": {
                "description": "Login for Oauth",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Oauth"
                ],
                "summary": "Authorize Oauth Login",
                "parameters": [
                    {
                        "description": "login data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.OauthLoginJsonRequest"
                        }
                    }
                ],
                "responses": {
                    "302": {
                        "description": "Found"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.BadRequestResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/schemas.ForbiddenResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/oauth/token/": {
            "post": {
                "description": "Get Authorization Token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Oauth"
                ],
                "summary": "Oauth Token",
                "parameters": [
                    {
                        "description": "code data",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.Oauth2TokenJsonRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.BadRequestResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/schemas.ForbiddenResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Get All User",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get All User",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page_size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UserPaginateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.BadRequestResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schemas.UnauthorizedResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "OAuth2Password": []
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
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UserCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UserCreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.BadRequestResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Get detail user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get Detail User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UserDetailResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.BadRequestResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schemas.NotFoundResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Update User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Update User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schemas.UserUpdateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schemas.BadRequestResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schemas.NotFoundResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "OAuth2Password": []
                    }
                ],
                "description": "Delete user",
                "tags": [
                    "User"
                ],
                "summary": "Delete User",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/schemas.NotFoundResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/schemas.InternalServerErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.ArrayClientRegisterResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.ClientRegiterResponse"
                    }
                }
            }
        },
        "schemas.BadRequestResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.ClientRegiterRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schemas.ClientRegiterResponse": {
            "type": "object",
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "client_secret": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "schemas.ForbiddenResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.InternalServerErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "schemas.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "token_type": {
                    "type": "string"
                }
            }
        },
        "schemas.NotFoundResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.Oauth2TokenJsonRequest": {
            "type": "object",
            "required": [
                "client_id",
                "client_secret",
                "code",
                "grant_type",
                "redirect_uri"
            ],
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "client_secret": {
                    "type": "string"
                },
                "code": {
                    "type": "string"
                },
                "grant_type": {
                    "type": "string"
                },
                "redirect_uri": {
                    "type": "string"
                }
            }
        },
        "schemas.OauthLoginJsonRequest": {
            "type": "object",
            "required": [
                "client_id",
                "password",
                "redirect_uri",
                "response_type",
                "scope",
                "state",
                "username"
            ],
            "properties": {
                "client_id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "redirect_uri": {
                    "type": "string"
                },
                "response_type": {
                    "type": "string"
                },
                "scope": {
                    "type": "string"
                },
                "state": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.UnauthorizedResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "schemas.UserCreateRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_superuser": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.UserCreateResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_superuser": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.UserDetailResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_superuser": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.UserPaginateResponse": {
            "type": "object",
            "properties": {
                "counts": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "page_count": {
                    "type": "integer"
                },
                "page_size": {
                    "type": "integer"
                },
                "results": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schemas.UserDetailResponse"
                    }
                }
            }
        },
        "schemas.UserUpdateRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_superuser": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "schemas.UserUpdateResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "is_superuser": {
                    "type": "boolean"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "OAuth2Password": {
            "type": "oauth2",
            "flow": "password",
            "tokenUrl": "/auth/login"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Go Oauth2 Authorization Server",
	Description:      "Oauth2 Authorization server",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
