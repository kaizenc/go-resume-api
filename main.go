package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	// "fmt"
	"time"
	"log"
	"net/http"
)

func healthcheckHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// name := mux.Vars(request)["name"]
	err := json.NewEncoder(writer).Encode(map[string]bool{"ok": true})
	if err != nil {
		log.Fatalln("There was an error returning health check")
	}
}

func main() {
    // _router := mux.NewRouter()
	// router := _router.Host("www.yourdomain.com").Headers("Connection", "Keep-Alive")
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/health_check", healthcheckHandler).Methods("GET")
	
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("hello")
	})

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// good practice: timeout enforcement
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}