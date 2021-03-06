package claw

import (
	"fmt"
	"net/http"
	"reflect"
)

// Mutate generate a valid handler with a provided http.HandlerFunc
func mutate(h http.HandlerFunc) MiddleWare {
	return func(next http.Handler) http.Handler {
		return ClawFunc(func(rw http.ResponseWriter, req *http.Request) {
			h(rw, req)
			next.ServeHTTP(rw, req)
		})
	}
}

// Get the interface type and transform to MiddleWare type. If valid append to the Middleware stack
func toMiddleware(m []interface{}) []MiddleWare {
	var stack []MiddleWare
	if len(m) > 0 {
		for _, f := range m {
			switch v := f.(type) {
			case func(http.ResponseWriter, *http.Request):
				stack = append(stack, mutate(http.HandlerFunc(v)))
			case func(http.Handler) http.Handler:
				stack = append(stack, v)
			default:
				fmt.Println("[x] [", reflect.TypeOf(v), "] is not a valid MiddleWare Type.")
			}
		}
	}
	return stack
}
