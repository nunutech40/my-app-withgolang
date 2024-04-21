package main

import (
	"fmt"
	"log"
	"net/http"
)

// only testing hello world for connection
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {

	http.HandleFunc("/", helloWorld)

	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil) // get error dari listen tcp ports 8080, jika error nil, berarti server tidak error
	if err != nil {                          // jika error tidak nil, berarti servernya error
		log.Fatal("Listen and server: ", err)
	}
}
