package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {


	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)  // Send HTTP 200 OK status
    w.Write([]byte("OK"))         // Send response body "OK"
})


	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"status":  "ok",
			"service": "1",
		})
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, map[string]string{
			"message": "Hello from Service 1",
		})
	})


	log.Println("congrats! Service 1 successfully started on port 8001...")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatalf("oops!Server starting failed: %v", err)
	}
}

func jsonResponse(w http.ResponseWriter, data map[string]string) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
