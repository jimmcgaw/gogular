package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"runtime"
	"encoding/json"
)

var (
	router = mux.NewRouter()
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())  // Use all CPU cores

	router.HandleFunc("/", GetIndex)
	router.HandleFunc("/api/about", GetAbout)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./angapp/app")))

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "angapp/app/index.html")	
}

func GetAbout(w http.ResponseWriter, r *http.Request) {
	obj := map[string]string{}
	obj["aboutus"] = "Here is some about us data that is coming from an API call."
	someJson, _ := json.Marshal(obj)
	w.Header().Set("Content-Type", "application/json")
	w.Write(someJson)
}
