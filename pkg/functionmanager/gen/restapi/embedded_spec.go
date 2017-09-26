///////////////////////////////////////////////////////////////////////
// Copyright (C) 2016 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "This is the API server for the serverless function manager service.\n",
    "title": "Function Manager",
    "contact": {
      "email": "kstepniewski@vmware.com"
    },
    "version": "1.0.0"
  },
  "basePath": "/v1/function",
  "paths": {
    "/": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "List all existing functions",
        "operationId": "getFunctions",
        "parameters": [
          {
            "type": "string",
            "description": "Function state",
            "name": "state",
            "in": "query"
          },
          {
            "type": "array",
            "items": {
              "type": "string"
            },
            "collectionFormat": "multi",
            "description": "Filter on function tags",
            "name": "tags",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Function"
              }
            }
          },
          "default": {
            "description": "Custom error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Add a few function",
        "operationId": "addFunction",
        "parameters": [
          {
            "description": "function object",
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/Function"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "Function accepted for creation",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          },
          "405": {
            "description": "Invalid input"
          }
        }
      }
    },
    "/{functionName}": {
      "get": {
        "description": "Returns a single function",
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Find function by Name",
        "operationId": "getFunctionByName",
        "responses": {
          "200": {
            "description": "Successful operation",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid Name supplied"
          },
          "404": {
            "description": "Function not found"
          }
        }
      },
      "delete": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Deletes a function",
        "operationId": "deleteFunctionByName",
        "responses": {
          "204": {
            "description": "Successful deletion"
          },
          "400": {
            "description": "Invalid Name supplied"
          },
          "404": {
            "description": "Function not found"
          }
        }
      },
      "patch": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Store"
        ],
        "summary": "Updates a function",
        "operationId": "updateFunctionByName",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful update",
            "schema": {
              "$ref": "#/definitions/Function"
            }
          },
          "400": {
            "description": "Invalid input"
          },
          "404": {
            "description": "Function not found"
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to work on",
          "name": "functionName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{functionName}/runs": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Get function runs that are being executed",
        "operationId": "getRuns",
        "responses": {
          "200": {
            "description": "List of function runs",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Run"
              }
            }
          },
          "404": {
            "description": "Function not found"
          }
        }
      },
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Run a function",
        "operationId": "runFunction",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/Run"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Successful execution (blocking call)",
            "schema": {
              "$ref": "#/definitions/Run"
            }
          },
          "202": {
            "description": "Execution started (non-blocking call)",
            "schema": {
              "$ref": "#/definitions/Run"
            }
          },
          "404": {
            "description": "Function not found"
          },
          "500": {
            "description": "Execution failed (blocking call)"
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to run",
          "name": "functionName",
          "in": "path",
          "required": true
        }
      ]
    },
    "/{functionName}/runs/{runName}": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "Runner"
        ],
        "summary": "Get function run by its name",
        "operationId": "getRunByName",
        "responses": {
          "200": {
            "description": "Function Run",
            "schema": {
              "$ref": "#/definitions/Run"
            }
          },
          "404": {
            "description": "Function or Run not found"
          }
        }
      },
      "parameters": [
        {
          "pattern": "^[\\w\\d\\-]+$",
          "type": "string",
          "description": "Name of function to retrieve a run for",
          "name": "functionName",
          "in": "path",
          "required": true
        },
        {
          "type": "string",
          "format": "uuid",
          "description": "name of run to retrieve",
          "name": "runName",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "Function": {
      "type": "object",
      "required": [
        "name",
        "code",
        "image"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "createdTime": {
          "type": "integer"
        },
        "id": {
          "type": "string",
          "format": "uuid"
        },
        "image": {
          "type": "string"
        },
        "modifiedTime": {
          "type": "integer"
        },
        "name": {
          "type": "string",
          "pattern": "^[\\w\\d\\-]+$"
        },
        "schema": {
          "$ref": "#/definitions/Schema"
        },
        "state": {
          "$ref": "#/definitions/State"
        },
        "tags": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/Tag"
          }
        }
      }
    },
    "Run": {
      "type": "object",
      "properties": {
        "blocking": {
          "type": "boolean"
        },
        "executedTime": {
          "type": "integer",
          "readOnly": true
        },
        "finishedTime": {
          "type": "integer",
          "readOnly": true
        },
        "input": {
          "type": "object"
        },
        "name": {
          "type": "string",
          "format": "uuid",
          "readOnly": true
        },
        "output": {
          "type": "object",
          "readOnly": true
        },
        "secrets": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "Schema": {
      "type": "object",
      "properties": {
        "in": {
          "type": "object"
        },
        "out": {
          "type": "object"
        }
      }
    },
    "State": {
      "type": "string",
      "enum": [
        "INITIALIZED",
        "CREATING",
        "READY",
        "ERROR",
        "DELETED"
      ]
    },
    "Tag": {
      "type": "object",
      "properties": {
        "key": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Crud operations on functions",
      "name": "Store"
    },
    {
      "description": "Execution operations on functions",
      "name": "Runner"
    }
  ]
}`))
}
