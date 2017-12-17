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
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Dashboard"
  },
  "host": "localhost:3000",
  "paths": {
    "/destinations": {
      "get": {
        "tags": [
          "destinations",
          "Destinations"
        ],
        "summary": "Destinations",
        "operationId": "GetDestinations",
        "responses": {
          "200": {
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/destination"
              }
            }
          },
          "400": {
            "$ref": "#/responses/trait:global:400"
          },
          "404": {
            "$ref": "#/responses/trait:global:404"
          },
          "500": {
            "$ref": "#/responses/trait:global:500"
          }
        }
      }
    },
    "/destinations/{id}": {
      "get": {
        "tags": [
          "destinations",
          "Destinations"
        ],
        "summary": "Destination By ID",
        "operationId": "GetDestinationByID",
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/destination"
            }
          },
          "400": {
            "$ref": "#/responses/trait:global:400"
          },
          "404": {
            "$ref": "#/responses/trait:global:404"
          },
          "500": {
            "$ref": "#/responses/trait:global:500"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    },
    "/trips": {
      "get": {
        "tags": [
          "trip",
          "Trips"
        ],
        "summary": "Trips",
        "operationId": "GetTrips",
        "responses": {
          "200": {
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/trip"
              }
            }
          },
          "400": {
            "$ref": "#/responses/trait:global:400"
          },
          "404": {
            "$ref": "#/responses/trait:global:404"
          },
          "500": {
            "$ref": "#/responses/trait:global:500"
          }
        }
      }
    },
    "/trips/{id}": {
      "get": {
        "tags": [
          "trip",
          "Trips"
        ],
        "summary": "Trip By ID",
        "operationId": "GetTripByID",
        "responses": {
          "200": {
            "schema": {
              "$ref": "#/definitions/trip"
            },
            "examples": {
              "application/json": {
                "": "pariatur do dolore"
              }
            }
          },
          "400": {
            "$ref": "#/responses/trait:global:400"
          },
          "404": {
            "$ref": "#/responses/trait:global:404"
          },
          "500": {
            "$ref": "#/responses/trait:global:500"
          }
        }
      },
      "parameters": [
        {
          "type": "integer",
          "name": "id",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "destination": {
      "type": "object",
      "title": "Destination",
      "properties": {
        "budget": {
          "type": "integer"
        },
        "city": {
          "type": "string"
        },
        "country": {
          "type": "string"
        },
        "currency": {
          "type": "string"
        },
        "id": {
          "type": "integer"
        }
      }
    },
    "trip": {
      "type": "object",
      "title": "Trip",
      "properties": {
        "id": {
          "type": "integer"
        }
      }
    }
  },
  "responses": {
    "trait:global:400": {
      "schema": {
        "type": "string"
      }
    },
    "trait:global:404": {
      "schema": {
        "type": "string"
      }
    },
    "trait:global:500": {
      "schema": {
        "type": "string"
      }
    }
  }
}`))
}