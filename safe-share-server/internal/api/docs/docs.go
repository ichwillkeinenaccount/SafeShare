// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag/v2"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "components": {},
    "info": {"contact":{"name":"SafeShare GitHub","url":"https://github.com/ichwillkeinenaccount/SafeShare"},"description":"{{escape .Description}}","license":{"name":"All Rights Reserved","url":"https://en.wikipedia.org/wiki/All_rights_reserved"},"termsOfService":"https://github.com/ichwillkeinenaccount/SafeShare","title":"{{.Title}}","version":"{{.Version}}"},
    "externalDocs": {"description":"","url":""},
    "paths": {"/api/v1/text/":{"get":{"description":"get all texts","responses":{"200":{"content":{"application/json":{"schema":{"type":"string"}}},"description":"OK"},"400":{"content":{"application/json":{"schema":{"type":"string"}}},"description":"Bad Request"},"404":{"content":{"application/json":{"schema":{"type":"string"}}},"description":"Not Found"},"500":{"content":{"application/json":{"schema":{"type":"string"}}},"description":"Internal Server Error"}},"summary":"Show all texts","tags":["text"]}}},
    "openapi": "3.1.0"
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Title:            "SafeShare API",
	Description:      "This is the API for SafeShare, a secure text sharing service.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
