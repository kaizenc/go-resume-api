package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// _router := mux.NewRouter()
	// router := _router.Host("www.yourdomain.com").Headers("Connection", "Keep-Alive")
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health_check", HealthCheck).Methods("GET")
	r.HandleFunc("/api/v1/job", JobsAll).Methods("GET")
	r.HandleFunc("/api/v1/job/company/{company}", JobSingleCompany).Methods("GET")
	r.HandleFunc("/api/v1/job/{startdate}", JobSingleDate).Methods("GET")
	// r.HandleFunc("/api/v1/job/range", HealthCheck).Methods("POST")

	srv := &http.Server{
		Handler: r,
		Addr:    ":8000",
		// good practice: timeout enforcement
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
