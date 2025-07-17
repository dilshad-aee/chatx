package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	initDB()

	// Root route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintln(w, "chatx backend is running 🏃‍♂️🏃‍♂️")
	})

	// Login route with CORS
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for all requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			fmt.Println("OPTIONS request received for /login")
			return
		}

		// Call your actual login handler
		handleLogin(w, r)
	})

	// Signup route with CORS
	http.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers for all requests
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			fmt.Println("OPTIONS request received for /signup")
			return
		}

		// Call your actual signup handler
		handleSignup(w, r)
	})

	fmt.Println("server is running at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("server failed to start " + err.Error())
	}

}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
