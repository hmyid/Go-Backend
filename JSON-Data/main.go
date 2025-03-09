package main 

import (
	"JSON-Data/routers"
	"net/http"
	"fmt"
)

func DecodeAndEncode(res http.ResponseWriter, req *http.Request) {
	routers.RegistRoute()
	fmt.Fprintln(res, "Route executed!")
}

func main () {
	http.HandleFunc("/", DecodeAndEncode)

	fmt.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}