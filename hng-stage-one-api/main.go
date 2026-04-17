package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// Helper to write JSON response
func writeJSON(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

// Handlers
func rootHandler(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, map[string]string{
		"message": "API is running",
	})
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{
		"message": "healthy",
	})
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{
		"name":   "Olatayo Lateef Damilola",
		"email":  "lateefolatayo@gmail.com",
		"github": "https://github.com/devLateef",
	})
}

func notFoundHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	json.NewEncoder(w).Encode(map[string]string{
		"error": "route not found",
	})
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/me", meHandler)

	// Root handler with strict check
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			notFoundHandler(w, r)
			return
		}
		rootHandler(w, r)
	})

	server := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	log.Println("Server running on port 8000")
	log.Fatal(server.ListenAndServe())
}
