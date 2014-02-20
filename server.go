// 2014.02.14

package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runtime"
	"time"
)

var (
	Debug = true
	Port  = "8080"

	Domain = flag.String("domain", "", "Domain or interface to listen on")
)

// func init() {
// 	flag.BoolVar(&Debug, "debug", Debug, "Enable debugging (verbose output)")
// 	flag.StringVar(&Port, "port", Port, "HTTP listen port")

// 	// MUST run this to parse the above CLI flags
// 	flag.Parse()
// }

var (
	router = mux.NewRouter()
)

func init() {
	// Tasks
	router.HandleFunc("/", GetIndex).Methods("GET")

	http.Handle("/", router)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("/Users/smoochy/source/gogular/angapp/app"))))
}

func main() {
	// Use all CPU cores
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Start HTTP server
	server := SimpleHTTPServer(router, *Domain+":"+Port)
	log.Printf("HTTP server trying to listen on %v...\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP listen failed: %v\n", err)
	}
}

func SimpleHTTPServer(handler http.Handler, host string) *http.Server {
	server := http.Server{
		Addr:           host,
		Handler:        handler,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	return &server
}

func writeError(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	http.Error(w, err.Error(), statusCode)
}

//
// HTTP Handler functions
//

func GetIndex(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to ballz!")
}
