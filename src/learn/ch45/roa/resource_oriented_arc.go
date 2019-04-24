package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var employeeDB map[string]*Employee

func init() {
	employeeDB = map[string]*Employee{}
	employeeDB["Mike"] = &Employee{"e-1", "Mike", 32}
	employeeDB["Rose"] = &Employee{"e-2", "Rose", 33}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	qName := ps.ByName("name")
	var (
		ok       bool
		info     *Employee
		infoJson []byte
		err      error
	)

	if info, ok = employeeDB[qName]; !ok {
		w.Write([]byte("{\"error\":\"Not Found\"}"))
		return
	}

	if infoJson, err = json.Marshal(info); err != nil {
		w.Write([]byte(fmt.Sprintf("{\"error\":,\"%s\"}", err)))
		return
	}

	w.Write(infoJson)
}

// func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
// }

func main() {
	router := httprouter.New()
	// router.GET(path,handle) puth delete
	router.GET("/", Index)
	// router.GET("/hello/:name", Hello)
	router.GET("/employee/:name", GetEmployeeByName)

	log.Fatal(http.ListenAndServe(":8081", router))
}
