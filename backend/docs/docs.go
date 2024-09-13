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
        "/customer/customer": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaCustomer"
                ],
                "summary": "获取单一客户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "主键",
                        "name": "ID",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "createdAt",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "客户名",
                        "name": "customerName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "客户手机号",
                        "name": "customerPhoneData",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "管理角色ID",
                        "name": "sysUserAuthorityID",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "管理ID",
                        "name": "sysUserId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "updatedAt",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新客户信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaCustomer"
                ],
                "summary": "更新客户信息",
                "parameters": [
                    {
                        "description": "客户ID, 客户信息",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/example.ExaCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "更新客户信息",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaCustomer"
                ],
                "summary": "创建客户",
                "parameters": [
                    {
                        "description": "客户用户名, 客户手机号码",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/example.ExaCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "创建客户",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
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
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaCustomer"
                ],
                "summary": "删除客户",
                "parameters": [
                    {
                        "description": "客户ID",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/example.ExaCustomer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "删除客户",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/customer/customerList": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "ExaCustomer"
                ],
                "summary": "分页获取权限客户列表",
                "responses": {}
            }
        },
        "/init/checkdb": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "CheckDB"
                ],
                "summary": "查询用户数据库存在",
                "responses": {
                    "200": {
                        "description": "查询用户数据库存在",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object",
                                            "additionalProperties": true
                                        },
                                        "msg": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/init/initdb": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "InitDB"
                ],
                "summary": "初始化数据库",
                "parameters": [
                    {
                        "description": "初始化数据库参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.InitDB"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "初始化用户数据库",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/response.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "example.ExaCustomer": {
            "type": "object",
            "properties": {
                "ID": {
                    "description": "主键",
                    "type": "integer"
                },
                "createdAt": {
                    "type": "string"
                },
                "customerName": {
                    "description": "客户名",
                    "type": "string"
                },
                "customerPhoneData": {
                    "description": "客户手机号",
                    "type": "string"
                },
                "sysUser": {
                    "description": "管理详情",
                    "allOf": [
                        {
                            "$ref": "#/definitions/system.SysUser"
                        }
                    ]
                },
                "sysUserAuthorityID": {
                    "description": "管理角色ID",
                    "type": "integer"
                },
                "sysUserId": {
                    "description": "管理ID",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "request.InitDB": {
            "type": "object",
            "required": [
                "dbName"
            ],
            "properties": {
                "dbName": {
                    "description": "数据库名",
                    "type": "string"
                },
                "dbPath": {
                    "description": "sqlite数据库文件路径",
                    "type": "string"
                },
                "dbType": {
                    "type": "string"
                },
                "host": {
                    "description": "服务器地址",
                    "type": "string"
                },
                "password": {
                    "description": "数据库密码",
                    "type": "string"
                },
                "port": {
                    "description": "数据库连接端口",
                    "type": "string"
                },
                "userName": {
                    "description": "数据库用户名",
                    "type": "string"
                }
            }
        },
        "response.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "system.SysAuthority": {
            "type": "object",
            "properties": {
                "authorityId": {
                    "type": "integer"
                },
                "authorityName": {
                    "description": "角色名",
                    "type": "string"
                },
                "children": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "createdAt": {
                    "description": "创建时间",
                    "type": "string"
                },
                "dataAuthorityId": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "defaultRouter": {
                    "description": "默认菜单(默认dashboard)",
                    "type": "string"
                },
                "deletedAt": {
                    "description": "时间字段可以为空*time.Time",
                    "type": "string"
                },
                "parentId": {
                    "description": "父角色ID",
                    "type": "integer"
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                }
            }
        },
        "system.SysUser": {
            "type": "object",
            "properties": {
                "ID": {
                    "description": "主键",
                    "type": "integer"
                },
                "activeColor": {
                    "description": "活跃颜色",
                    "type": "string"
                },
                "authorities": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/system.SysAuthority"
                    }
                },
                "authority": {
                    "$ref": "#/definitions/system.SysAuthority"
                },
                "authorityId": {
                    "type": "integer"
                },
                "baseColor": {
                    "description": "基础颜色",
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "enable": {
                    "type": "integer"
                },
                "headerImg": {
                    "description": "用户头像",
                    "type": "string"
                },
                "nickName": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "sideMode": {
                    "description": "用户侧边主题",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userName": {
                    "description": "用户登录名",
                    "type": "string"
                },
                "uuid": {
                    "description": "用户UUID",
                    "type": "string"
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
