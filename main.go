package main

import (
	"github.com/nunutech40/my-app-withgolang/common/config"
	"github.com/nunutech40/my-app-withgolang/handlers"
	"github.com/nunutech40/my-app-withgolang/handlers/auth"
	"log"
	"net/http"
)

func main() {
	db := config.ConnectDb() // connect to mysql + db
	defer db.Close()

	// Injection
	// Inject auth with db
	authHandler := auth.NewHandler(db)

	// Routing here
	http.HandleFunc("/", handlers.HelloWorld)
	// auth routing
	http.HandleFunc("/register", authHandler.Register)

	// starting server
	log.Println("Starting server on port 8080...")
	// listening server
	err := http.ListenAndServe(":8080", nil) // get error dari listen tcp ports 8080, jika error nil, berarti server tidak error
	if err != nil {                          // jika error tidak nil, berarti servernya error
		log.Fatal("Listen and server: ", err)
	}
}
