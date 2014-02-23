package main

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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

	CreateDBTables()
	// uncomment and run to create Jim
	// CreatePerson()

	fmt.Println("Server is running and listening at localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type Person struct {
	Id        int64
	FirstName string `sql:"size:50"`
	LastName  string `sql:"size:50"`
}

func CreateDBTables() {
	db, err := gorm.Open("mysql", "root:@/gogular")
	if err != nil {
		panic(err)
	}
	db.CreateTable(Person{})
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

func CreatePerson() {
	person := Person{
		FirstName: "Jim",
		LastName:  "McGaw",
	}

	db, err := gorm.Open("mysql", "root:@/gogular")
	// TODO : load all people from MySQL db and return as JSON
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	db.Save(&person)
}

func GetPerson() {
	db, err := gorm.Open("mysql", "root:@/gogular")
	// TODO : load all people from MySQL db and return as JSON
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	d := db.DB()
	d.Ping()

	// person := db.First(&persons) // grab first record
}

func GetAllPersons(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "root:@/gogular")
	// TODO : load all people from MySQL db and return as JSON
	if err != nil {
		panic(fmt.Sprintf("Got error when connect database, the error is '%v'", err))
	}

	d := db.DB()
	d.Ping()

	// db.Find(&persons)

	obj := map[string]string{}
	obj["people"] = "insert list of people here"
	someJson, _ := json.Marshal(obj)
	w.Header().Set("Content-Type", "application/json")
	w.Write(someJson)
}
