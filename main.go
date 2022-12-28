package main

import (
	"awesomeProject/Router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API + MongoDB Example")
	log.Fatal(http.ListenAndServe(":4000", Router.Router()))
}
