package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var jobs = []Job{
	{
		Company:         "Howl",
		Position:        "Platform Engineer",
		StartDateString: "2020-05-01",
		EndDateString:   "2023-09-01",
	},
	{
		Company:         "Narrativ",
		Position:        "Software Engineer Intern",
		StartDateString: "2019-01-01",
		EndDateString:   "2019-08-01",
	},
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	if err != nil {
		log.Fatalln("There was an error returning health check")
	}
}

func JobsAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jobs)
	if err != nil {
		log.Fatalln("There was an error returning list of all jobs")
	}
}

func JobSingle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	company := vars["company"]
	var job Job
	jobFound := false
	for _, _job := range jobs {
		if _job.Company == company {
			job = _job
			jobFound = true
		}
	}
	if !jobFound {
		// todo: needs to be a better way to check for found job
		// also needs to return an error code
		fmt.Printf("Job not found")
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(job)
	if err != nil {
		log.Fatalln("There was an error returning your job")
	}
}
