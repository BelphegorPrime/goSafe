package main
import (
	"net/http"
	"fmt"
	"runtime"
)

func index_func(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}
func save_func(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}
func get_all_function(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "hello world, I'm running on %s with an %s CPU ", runtime.GOOS, runtime.GOARCH)
}