// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2019-08-16 18:20:43.6822963 +0800 CST m=+0.080979601

package docs

import (
	"bytes"
	"encoding/json"

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
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/adjudication": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "adjudication"
                ],
                "summary": "get adjudication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "{integer}",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/adjudication/history": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "adjudication"
                ],
                "summary": "get adjudication",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "{integer}",
                        "description": "pageSize",
                        "name": "pageSize",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        },
        "/api/v1/node": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "node"
                ],
                "summary": "get nodes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "{integer}",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Node"
                            }
                        }
                    }
                }
            },
            "put": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "node"
                ],
                "summary": "modify node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "in",
                        "name": "in",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.NodePutIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "node"
                ],
                "summary": "add node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "in",
                        "name": "in",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.NodeIn"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "node"
                ],
                "summary": "delete node",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "{integer}",
                        "description": "id",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "boolean"
                        }
                    }
                }
            }
        },
        "/api/wechat": {
            "get": {
                "tags": [
                    "wechat"
                ],
                "summary": "config for wechat's settings",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "text/xml"
                ],
                "tags": [
                    "wechat"
                ],
                "summary": "Listening WechatEvent",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.Result"
                        }
                    }
                }
            }
        },
        "/api/wxseession": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "session"
                ],
                "summary": "get WxSession",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/models.TokenOut"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Node": {
            "type": "object",
            "properties": {
                "created": {
                    "type": "string"
                },
                "deleted": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nodeType": {
                    "type": "integer"
                },
                "parent": {
                    "type": "integer"
                },
                "remark": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                },
                "updated": {
                    "type": "string"
                },
                "weight": {
                    "type": "number"
                },
                "wxAccountId": {
                    "type": "integer"
                }
            }
        },
        "models.NodeIn": {
            "type": "object",
            "required": [
                "name",
                "nodeType"
            ],
            "properties": {
                "name": {
                    "type": "string"
                },
                "nodeType": {
                    "description": "1: dir; 2: opt;",
                    "type": "integer"
                },
                "parent": {
                    "type": "integer"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "models.NodePutIn": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "parent": {
                    "type": "integer"
                },
                "weight": {
                    "type": "number"
                }
            }
        },
        "models.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string"
                },
                "serverTime": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "models.TokenOut": {
            "type": "object",
            "properties": {
                "expireat": {
                    "type": "string"
                },
                "expirein": {
                    "type": "integer"
                },
                "token": {
                    "type": "string"
                }
            }
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
var SwaggerInfo = swaggerInfo{ Schemes: []string{}}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface {}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
