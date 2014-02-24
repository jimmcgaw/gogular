package main

import (
	"fmt"
	"github.com/ziutek/mymysql/mysql"
	_ "github.com/ziutek/mymysql/native"
	// "os"
	// "github.com/ziutek/mymysql/thrsafe"
)

type Person struct {
	id   int
	name string
}

func main() {
	db := mysql.New("tcp", "", "127.0.0.1:3306", "testuser", "TestPasswd9", "test")

	err := db.Connect()
	if err != nil {
		panic(err)
	}

	rows, res, err := db.Query("SELECT * FROM people")
	if err != nil {
		panic(err)
	}

	for _, row := range rows {
		for _, col := range row {
			if col == nil {
				// col has NULL value
			} else {
				// Do something with text in col (type []byte)
			}
		}
		// You can get specific value from a row
		// id := row.Int(0)
		// name := string(row[1].([]byte))

		id := row.Int(res.Map("id"))
		name := row.Str(res.Map("name"))

		// You can use it directly if conversion isn't needed
		fmt.Println(id)
		fmt.Println(name)

		// You can get converted value
		// number := row.Int(0) // Zero value
		// str := row.Str(1)    // First value

		// // You may get values by column name
		// first := res.Map("id")
		// second := res.Map("name")
		// val1, val2 := row.Str(first), row.Str(second)

		// myInt := res.Map("name")
		// name := string(myInt)
		// os.Stdout.Write(name)
	}
}
