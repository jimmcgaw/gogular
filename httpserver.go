// httpserver.go
package main

import (
	"flag"
	"net/http"
	"runtime"
)

var port = flag.String("port", "8080", "Define what TCP port to bind to")
var root = flag.String("root", ".", "Define the root filesystem path")
var projectRoot = flag.String("projectRoot", "angapp/app", "Define the project root filesystem path")

func main() {
	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*projectRoot)))
}
