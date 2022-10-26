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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/customer": {
            "put": {
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
            },
            "post": {
                "description": "this api creates new customer",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "create customer api",
                "parameters": [
                    {
                        "description": "Customer",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customer.CustomerRequest"
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
        "/v1/customer/{id}": {
            "get": {
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
        "/v1/post": {
            "put": {
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
        "/v1/register/{code}": {
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
        "customer.AddressRequest": {
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
        "customer.CustomerRequest": {
            "type": "object",
            "properties": {
                "addresses": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/customer.AddressRequest"
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
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "customer.CustomerWithoutPost": {
            "type": "object",
            "properties": {
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
                }
            }
        },
        "models.Error": {
            "type": "object",
            "properties": {
                "error": {
                    "description": "` + "`" + `json:\"error\"` + "`" + `",
                    "type": "string"
                }
            }
        },
        "models.CustomerRegister": {
            "type": "object",
            "properties": {
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
                },
                "username": {
                    "description": "` + "`" + `json:\"username\"` + "`" + `",
                    "type": "string"
                }
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
