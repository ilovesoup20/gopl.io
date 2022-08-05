package main

import (
	"fmt"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) foo() {
	fmt.Printf("%p\t%p\t%v\n", db, &db, db)
	db["hats"] = 10
	fmt.Printf("%p\t%p\t%v\n", db, &db, db)
}

func main() {
	db := database{"shoes": 50, "socks": 5}
	fmt.Printf("%p\t%p\t%v\n", db, &db, db)
	db.foo()
	fmt.Printf("%p\t%p\t%v\n", db, &db, db)
	// log.Fatal(http.ListenAndServe("localhost:8000", db))

}
