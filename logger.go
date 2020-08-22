package logger

import (
	"context"
	"fmt"
	"net/http"
)

const (
	klocator = "locator"
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

func New() *logger {
	return &logger{}
}

func readLocator(ctx context.Context) string {
	locator, ok := ctx.Value(klocator).(string)
	if !ok {
		return ""
	}
	return locator
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

func colorMethod(method string) string {
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
	return fmt.Sprintf("%s%s%s", color, method, kreset)
}

func formatPath(locator string, resource string) string {
	var path string
	if locator == "" {
		path = resource
	} else {
		path = resource + "/" + locator
	}
	return path
}

func log(locator string, method string, resource string) {
	coloredMethod := colorMethod(method)
	formattedPath := formatPath(locator, resource)
	fmt.Printf("[%s] %s\n", coloredMethod, formattedPath)
}

func (l *logger) Do(ctx context.Context, w http.ResponseWriter) context.Context {
	locator := readLocator(ctx)
	method := readMethod(ctx)
	resource := readResource(ctx)
	log(locator, method, resource)
	return ctx
}
