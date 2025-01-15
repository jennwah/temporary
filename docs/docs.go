// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/friends": {
            "post": {
                "description": "Creates a new user friend pair record in the database.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "Create a new user friend pair",
                "parameters": [
                    {
                        "type": "string",
                        "name": "user_id",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "Create user friend pair payload",
                        "name": "createRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/friend.CreateReq"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/friends/nearby": {
            "get": {
                "description": "Retrieve nearby friends from database, within radius meter specified",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "friends"
                ],
                "summary": "Get nearby friends with radius meter specified",
                "parameters": [
                    {
                        "type": "number",
                        "default": 5000,
                        "description": "Radius in meters (default: 5000)",
                        "name": "radius_meter",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "User ID (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/friend.GetNearbyResp"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Creates a new user record in the database.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create a new user",
                "parameters": [
                    {
                        "description": "Create user payload",
                        "name": "createRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/user.CreateResp"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    }
                }
            }
        },
        "/users/{id}": {
            "get": {
                "description": "Retrieve an existing user by UUID from database.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.GetResp"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an existing user record from the database.",
                "tags": [
                    "users"
                ],
                "summary": "Delete an existing user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    }
                }
            },
            "patch": {
                "description": "Update an existing user record in the database.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Update an existing user",
                "parameters": [
                    {
                        "description": "Update user payload",
                        "name": "patchRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.PatchRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User ID (UUID)",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/controller.ErrorResponseModel"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ErrorResponseModel": {
            "type": "object",
            "properties": {
                "error_code": {
                    "type": "string"
                },
                "error_message": {
                    "type": "string"
                }
            }
        },
        "friend.CreateReq": {
            "type": "object",
            "required": [
                "friend_id"
            ],
            "properties": {
                "friend_id": {
                    "type": "string"
                }
            }
        },
        "friend.Friend": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "friend.GetNearbyResp": {
            "type": "object",
            "required": [
                "friends"
            ],
            "properties": {
                "friends": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/friend.Friend"
                    }
                }
            }
        },
        "user.CreateRequest": {
            "type": "object",
            "required": [
                "address",
                "date_of_birth",
                "latitude",
                "longitude",
                "name"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "user.CreateResp": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "user.GetResp": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.PatchRequest": {
            "type": "object",
            "required": [
                "address",
                "date_of_birth",
                "latitude",
                "longitude",
                "name"
            ],
            "properties": {
                "address": {
                    "type": "string"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "latitude": {
                    "type": "number"
                },
                "longitude": {
                    "type": "number"
                },
                "name": {
                    "type": "string",
                    "minLength": 3
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
