package main

import (
	"log"
	"net/http"

	"github.com/nunutech40/my-app-withgolang/common/config"
	"github.com/nunutech40/my-app-withgolang/handlers"
	"github.com/nunutech40/my-app-withgolang/handlers/auth"
)

func main() {
	db := config.ConnectDb() // connect to mysql + db
	defer db.Close()

	// Injection
	// Inject auth with db
	handlers := handlers.NewHandler(db)

	// auth routing
	// register
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		auth.Register(handlers, w, r)
	})
	// login
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.Login(handlers, w, r)
	})

	// starting server
	log.Println("Starting server on port 8080...")
	// listening server
	err := http.ListenAndServe(":8080", nil) // get error dari listen tcp ports 8080, jika error nil, berarti server tidak error
	if err != nil {                          // jika error tidak nil, berarti servernya error
		log.Fatal("Listen and server: ", err)
	}
}
