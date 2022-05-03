package main

import (
	"assigment3/controllers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Update)
	http.Handle("/static", http.FileServer(http.Dir("assets")))
	address := "localhost:8083"
	log.Println("Server started at: ", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err)
	}
}
