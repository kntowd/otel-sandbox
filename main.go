package main

import (
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Hello, World!"))
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)

	// Start the server on port 8080
	http.ListenAndServe(":8080", nil)
}
