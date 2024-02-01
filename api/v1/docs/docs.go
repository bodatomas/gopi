// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Tomáš Boďa",
            "url": "https://github.com/dvandyy",
            "email": "tomasboda.dev@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "security": [
                    {
                        "JWT_TOKEN": []
                    }
                ],
                "description": "Retun a hello message if everything is ok",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Welcome"
                ],
                "summary": "Show a hello message",
                "responses": {
                    "200": {
                        "description": "Return 'Hello from gopi!'",
                        "schema": {
                            "$ref": "#/definitions/models.HelloResponse"
                        }
                    }
                }
            }
        },
        "/boards/new": {
            "post": {
                "security": [
                    {
                        "JWT_TOKEN": []
                    }
                ],
                "description": "Create a new board in database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Boards"
                ],
                "summary": "Create new board",
                "parameters": [
                    {
                        "description": "Create board with Title",
                        "name": "CreateBoardRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateBoardRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateBoardResponse"
                        }
                    }
                }
            }
        },
        "/boards/{uid}": {
            "get": {
                "security": [
                    {
                        "JWT_TOKEN": []
                    }
                ],
                "description": "Return board with unique id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Boards"
                ],
                "summary": "Get board with UUID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Board unique ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Board"
                        }
                    }
                }
            }
        },
        "/team/new": {
            "post": {
                "security": [
                    {
                        "JWT_TOKEN": []
                    }
                ],
                "description": "Create a new team in database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Teams"
                ],
                "summary": "Create new team",
                "parameters": [
                    {
                        "description": "Create new team with workspace_id and name",
                        "name": "CreateTeamRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateTeamRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateTeamResponse"
                        }
                    }
                }
            }
        },
        "/users/login": {
            "post": {
                "description": "Authenticate user with jwt token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login user with Email and Password.",
                        "name": "LoginRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.LoginResponse"
                        }
                    }
                }
            }
        },
        "/users/register": {
            "post": {
                "description": "Create new user in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "Create user with Email and Password.",
                        "name": "RegisterRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.RegisterResponse"
                        }
                    }
                }
            }
        },
        "/users/{uid}": {
            "get": {
                "security": [
                    {
                        "JWT_TOKEN": []
                    }
                ],
                "description": "Return user with unique id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get user with UUID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User unique ID",
                        "name": "uid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    }
                }
            }
        },
        "/workspace/new": {
            "post": {
                "security": [
                    {
                        "JWT_TOKEN": []
                    }
                ],
                "description": "Create a new workspace in database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Workspaces"
                ],
                "summary": "Create new workspace",
                "parameters": [
                    {
                        "description": "Create new workspace with owner_id and name",
                        "name": "CreateWorkspaceRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CreateWorkspaceRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CreateWorkspaceResponse"
                        }
                    }
                }
            }
        },
        "/workspace/user": {
            "get": {
                "security": [
                    {
                        "JWT_TOKEN": []
                    }
                ],
                "description": "Return all workspaces for certain user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Workspaces"
                ],
                "summary": "Get workspaces for user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.UserWorkspacesResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Board": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-01-22 17:03:50.283466+00"
                },
                "description": {
                    "type": "string",
                    "example": "Board description"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Title"
                },
                "unique_id": {
                    "type": "string",
                    "example": "b-1706453613063fa2eb4"
                }
            }
        },
        "models.CreateBoardRequest": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "My description"
                },
                "team_id": {
                    "type": "string",
                    "example": "t-1706453613063fa2eb4"
                },
                "title": {
                    "type": "string",
                    "example": "My title"
                }
            }
        },
        "models.CreateBoardResponse": {
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
        "models.CreateTeamRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "My team"
                },
                "workspace_id": {
                    "type": "string",
                    "example": "w-1706453613063fa2eb4"
                }
            }
        },
        "models.CreateTeamResponse": {
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
        "models.CreateWorkspaceRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "My workspace"
                },
                "owner": {
                    "type": "string",
                    "example": "u-1706453613063fa2eb4"
                }
            }
        },
        "models.CreateWorkspaceResponse": {
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
        "models.HelloResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.LoginRequest": {
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
        "models.LoginResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "models.RegisterRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "password": {
                    "type": "string",
                    "example": "Password123\u0026"
                }
            }
        },
        "models.RegisterResponse": {
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
        "models.User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2024-01-22 17:03:50.283466+00"
                },
                "email": {
                    "type": "string",
                    "example": "email@email.com"
                },
                "first_name": {
                    "type": "string",
                    "example": "FirstName"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "last_name": {
                    "type": "string",
                    "example": "LastName"
                },
                "password": {
                    "type": "string",
                    "example": "hash"
                },
                "role": {
                    "type": "string",
                    "example": "Tester"
                },
                "unique_id": {
                    "type": "string",
                    "example": "u-1706453613063fa2eb4"
                }
            }
        },
        "models.UserWorkspacesResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "integer"
                },
                "workspaces": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.WorkspaceResponse"
                    }
                }
            }
        },
        "models.WorkspaceResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "My workspace"
                },
                "unique_id": {
                    "type": "string",
                    "example": "w-1706453613063fa2eb4"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT_TOKEN": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:4000",
	BasePath:         "/api/v1/",
	Schemes:          []string{},
	Title:            "Gopi API",
	Description:      "REST api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
