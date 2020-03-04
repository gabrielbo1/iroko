package api

import (
	"net/http"
)

//Route - Defines structure for mapping the routes that will be mapped by gorilla.
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - Array with all mapped APIs.
type Routes []Route

var routes = Routes{
	Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/login",
		HandlerFunc: Login,
	},
	Route{
		Name:        "Refresh",
		Method:      "POST",
		Pattern:     "/refresh",
		HandlerFunc: Refresh,
	},
	Route{
		Name:        "Test",
		Method:      "GET",
		Pattern:     "/test",
		HandlerFunc: Test,
	},
}
