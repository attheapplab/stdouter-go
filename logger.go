package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	kbody = "body"
	kidentifier = "identifier"
	kmethod = "method"
	kquery = "query"
	kresource = "resource"
)

const (
	kreset  = "\033[0m"
	kred    = "\033[31m"
	kgreen  = "\033[32m"
	kyellow = "\033[33m"
	kblue   = "\033[34m"
	kpurple = "\033[35m"
	kcyan   = "\033[36m"
	kgray   = "\033[37m"
	kwhite  = "\033[97m"
)

type logger struct {}

func New() *logger {
	return &logger{}
}

func colorMethod(method string) string {
	var color string
	switch method {
	case http.MethodDelete:
		color = kred
	case http.MethodGet:
		color = kgreen
	case http.MethodPatch:
		color = kblue
	case http.MethodPost:
		color = kyellow
	default:
		color = kwhite
	}
	coloredMethod := fmt.Sprintf("%s%s%s", color, method, kreset)
	return coloredMethod
}

func formatInputs(body map[string]string, query map[string]string) string {
	var inputs = make(map[string]string)
	for k, v := range body {
		inputs[k] = v
	}
	for k, v := range query {
		inputs[k] = v
	}
	serializedInputs, _ := json.Marshal(inputs)
	formattedInputs := string(serializedInputs)
	return formattedInputs
}

func formatPath(identifier string, resource string) string {
	var path string
	if identifier == "" {
		path = resource
	} else {
		path = resource + "/" + identifier
	}
	return path
}

func readBody(ctx context.Context) map[string]string {
	var bodyParams = make(map[string]string)
	if body, ok := ctx.Value(kbody).(map[string]string); ok {
		for k, v := range body {
			bodyParams[k] = v
		}
	}
	return bodyParams
}

func readIdentifier(ctx context.Context) string {
	identifier, ok := ctx.Value(kidentifier).(string)
	if !ok {
		return ""
	}
	return identifier
}

func readMethod(ctx context.Context) string {
	method, ok := ctx.Value(kmethod).(string)
	if !ok {
		return ""
	}
	return method
}

func readQuery(ctx context.Context) map[string]string {
	var queryParams = make(map[string]string)
	if query, ok := ctx.Value(kquery).(map[string]string); ok {
		for k, v := range query {
			queryParams[k] = v
		}
	}
	return queryParams
}

func readResource(ctx context.Context) string {
	resource, ok := ctx.Value(kresource).(string)
	if !ok {
		return ""
	}
	return resource
}

func log(body map[string]string, identifier string, method string, query map[string]string, resource string) {
	coloredMethod := colorMethod(method)
	formattedInputs := formatInputs(body, query)
	formattedPath := formatPath(identifier, resource)
	fmt.Printf("[%s] %s %s\n", coloredMethod, formattedPath, formattedInputs)
}

func (l *logger) Do(ctx context.Context, w http.ResponseWriter) context.Context {
	body := readBody(ctx)
	identifier := readIdentifier(ctx)
	method := readMethod(ctx)
	query := readQuery(ctx)
	resource := readResource(ctx)
	log(body, identifier, method, query, resource)
	return ctx
}
