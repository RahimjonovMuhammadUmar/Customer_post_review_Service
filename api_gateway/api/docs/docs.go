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
        "termsOfService": "2 term exam",
        "contact": {
            "name": "Muhammad Umar",
            "url": "https://t.me/muhammad_ummar",
            "email": "rahimzanovmuhammadumar@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/admin/login/{username}/{password}": {
            "get": {
                "description": "Logins admin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Login admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AdminResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/customer": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "this api updates customer by id in database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "Update customer api",
                "parameters": [
                    {
                        "description": "Customer",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customer.CustomerWithoutPost"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/v1/customer/search": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "this api searches customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "search customer api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Order:DescOrAsc",
                        "name": "orderBy",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Field:Value",
                        "name": "fieldValue",
                        "in": "query",
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
                        "description": "Bad Request"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/v1/customer/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "this api finds existing customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "get customer api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "this api deletes customer from database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "Delete customer api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/login/{email}/{password}": {
            "get": {
                "description": "Logins customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login customer",
                "parameters": [
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/customer.CustomerWithoutPost"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/moderator/login/{username}/{password}": {
            "get": {
                "description": "Logins moderator",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Moderator"
                ],
                "summary": "Login moderator",
                "parameters": [
                    {
                        "type": "string",
                        "description": "username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "password",
                        "name": "password",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.AdminResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/post": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "update post api",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "Updates post by id",
                "parameters": [
                    {
                        "description": "Update post by id",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.PostWithoutReview"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.PostWithoutReview"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "this api creates new post",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "create post api",
                "parameters": [
                    {
                        "description": "Post",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/post.PostRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/v1/post/customers_posts/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get posts of customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "Gets post by customers id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "customer_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.Posts"
                        }
                    }
                }
            }
        },
        "/v1/post/delete_customers_posts/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete Post by Customer Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "Delete customers posts",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "customer_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/post/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get Post infos with id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "Get post with customer information",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "post_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/post.PostWithCustomerInfo"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete Post and it's reviews by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "post"
                ],
                "summary": "Delete post from database",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "post_id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/v1/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register for authentication",
                "parameters": [
                    {
                        "description": "user data",
                        "name": "userData",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.CustomerRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Message sended to your email succesfully"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/register/{code}/{email}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Verify for authentication",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "email",
                        "name": "email",
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
                            "$ref": "#/definitions/models.Error"
                        }
                    }
                }
            }
        },
        "/v1/review": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "this api creates new review",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "review"
                ],
                "summary": "create review api",
                "parameters": [
                    {
                        "description": "Review",
                        "name": "review",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/review.ReviewRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "type": "json"
                        }
                    }
                }
            }
        },
        "/v1/review/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "this api gets review from database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "review"
                ],
                "summary": "Get review api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "this api deletes review from database",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "review"
                ],
                "summary": "Delete review api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/v1/review_by_custID/{id}": {
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "this api deletes review by customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "review"
                ],
                "summary": "delete review by cust api",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
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
                    }
                }
            }
        }
    },
    "definitions": {
        "customer.Address": {
            "type": "object",
            "properties": {
                "house_number": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "customer.CustomerWithoutPost": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/customer.Address"
                    }
                },
                "bio": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "models.AddressRequest": {
            "type": "object",
            "properties": {
                "house_number": {
                    "type": "integer"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "models.AdminResponse": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "models.CustomerRegister": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.AddressRequest"
                    }
                },
                "bio": {
                    "description": "` + "`" + `json:\"username\"` + "`" + `",
                    "type": "string"
                },
                "email": {
                    "description": "` + "`" + `json:\"email\"` + "`" + `",
                    "type": "string"
                },
                "firstName": {
                    "description": "` + "`" + `json:\"first_name\"` + "`" + `",
                    "type": "string"
                },
                "lastName": {
                    "description": "` + "`" + `json:\"last_name\"` + "`" + `",
                    "type": "string"
                },
                "password": {
                    "description": "` + "`" + `json:\"password\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "error": {}
            }
        },
        "post.Address": {
            "type": "object",
            "properties": {
                "house_number": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "street": {
                    "type": "string"
                }
            }
        },
        "post.Customer": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Address"
                    }
                },
                "bio": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "post.Media": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "post.MediaRequest": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "post.Post": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "medias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Media"
                    }
                },
                "name": {
                    "type": "string"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Review"
                    }
                }
            }
        },
        "post.PostRequest": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "medias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.MediaRequest"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "post.PostWithCustomerInfo": {
            "type": "object",
            "properties": {
                "customer": {
                    "$ref": "#/definitions/post.Customer"
                },
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "medias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Media"
                    }
                },
                "name": {
                    "type": "string"
                },
                "reviews": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Review"
                    }
                }
            }
        },
        "post.PostWithoutReview": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "medias": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Media"
                    }
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "post.Posts": {
            "type": "object",
            "properties": {
                "posts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/post.Post"
                    }
                }
            }
        },
        "post.Review": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "post_id": {
                    "type": "integer"
                },
                "review": {
                    "type": "integer"
                }
            }
        },
        "review.ReviewRequest": {
            "type": "object",
            "properties": {
                "customer_id": {
                    "type": "integer"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "post_id": {
                    "type": "integer"
                },
                "review": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "54.95.64.84:8800",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Api's of all services",
	Description:      "This web app is running on AWS EC2 instance",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
