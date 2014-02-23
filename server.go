package main

import (
	// "database/sql"
	"encoding/json"
	// "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"runtime"
)

var (
	router = mux.NewRouter()
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores

	router.HandleFunc("/", GetIndex)
	router.HandleFunc("/api/about", GetAbout)
	router.HandleFunc("/api/persons", GetAllPersons)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./angapp/app")))

	http.Handle("/", router)

	http.ListenAndServe(":8080", nil)
}

type Person struct {
	id         uint8
	first_name string
	last_name  string
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

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:@/gogular")
	// TODO : load all people from MySQL db and return as JSON
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	d := db.DB()
}
