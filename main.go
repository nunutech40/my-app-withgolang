package main

import (
	"github.com/nunutech40/my-app-withgolang/handlers"
	"log"
	"net/http"
)

func main() {

	// Routing here
	http.HandleFunc("/", handlers.HelloWorld)

	// starting server
	log.Println("Starting server on port 8080...")
	// listening server
	err := http.ListenAndServe(":8080", nil) // get error dari listen tcp ports 8080, jika error nil, berarti server tidak error
	if err != nil {                          // jika error tidak nil, berarti servernya error
		log.Fatal("Listen and server: ", err)
	}
}
