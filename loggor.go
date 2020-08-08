package loggor

import (
	"context"
	"fmt"
	"net/http"
)

const (
	kid = "id"
	kmethod = "method"
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

func NewLogger() *logger {
	return &logger{}
}

func readId(ctx context.Context) string {
	id, ok := ctx.Value(kid).(string)
	if !ok {
		return ""
	}
	return id
}

func readMethod(ctx context.Context) string {
	method, ok := ctx.Value(kmethod).(string)
	if !ok {
		return ""
	}
	return method
}

func readResource(ctx context.Context) string {
	resource, ok := ctx.Value(kresource).(string)
	if !ok {
		return ""
	}
	return resource
}

func log(id string, method string, resource string) {
	var color string
	switch method {
	case http.MethodDelete:
		color = kred
	case http.MethodGet:
		color = kgreen
	case http.MethodPost:
		color = kyellow
	default:
		color = kwhite
	}
	coloredMethod := fmt.Sprintf("%s%s%s", color, method, kreset)
	fmt.Printf("[%s] %s {%s}\n", coloredMethod, resource, id)
}

func (l *logger) Do(ctx context.Context, w http.ResponseWriter) context.Context {
	id := readId(ctx)
	method := readMethod(ctx)
	resource := readResource(ctx)
	log(id, method, resource)
	return ctx
}
