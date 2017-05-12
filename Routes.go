package main
import (
	"net/http"
)
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route
var routes = Routes{

	Route{
		"index",
		"GET",
		"/",
		index_func,
	},
	Route{
		"save",
		"POST",
		"/save",
		save_func,
	},
	Route{
		"get",
		"POST",
		"/get",
		get_func,
	},
	Route{
		"getAll",
		"GET",
		"/getAll",
		get_all_function,
	},
}