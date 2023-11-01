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
            "name": "Alex Chinaev",
            "url": "https://vk.com/l.chinaev",
            "email": "ax.chinaev@yandex.ru"
        },
        "license": {
            "name": "AS IS (NO WARRANTY)"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/auth/login": {
            "post": {
                "description": "create user session and put it into cookie",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "login user",
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/logout": {
            "post": {
                "description": "delete current session and nullify cookie",
                "tags": [
                    "auth"
                ],
                "summary": "logout user",
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "add new user to db and return it id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "register user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/films/cast/{id}": {
            "get": {
                "description": "Get a list of films based on the cast name.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cast"
                ],
                "summary": "Get cast page",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The Films of the Cast you want to retrieve.",
                        "name": "cast",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/films/genre/{genre}": {
            "get": {
                "description": "Get a list of films based on the specified genre.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Films"
                ],
                "summary": "Get films by genre",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The genre of the Films you want to retrieve.",
                        "name": "genre",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/films/{id}": {
            "get": {
                "description": "Get content for Film page",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Films"
                ],
                "summary": "Get Film data by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Id film you want to get.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/profile/update": {
            "post": {
                "description": "update user data in db and return it",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "update profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/profile_http.Result"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "body": {
                                            "type": "object",
                                            "properties": {
                                                "user": {
                                                    "$ref": "#/definitions/domain.User"
                                                }
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/api/v1/profile/{id}": {
            "get": {
                "description": "Get user data by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "profile"
                ],
                "summary": "Get user by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "The user id you want to retrieve.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "json"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "imageData": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "imagePath": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "profile_http.Result": {
            "type": "object",
            "properties": {
                "body": {},
                "err": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1",
	BasePath:         "/",
	Schemes:          []string{"http"},
	Title:            "Netfilx API",
	Description:      "API of the nelfix project by holi",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
