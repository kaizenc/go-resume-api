package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var jobs_db = []Job{
	{
		Company:         "Howl",
		Position:        "Platform Engineer",
		StartDateString: "20210501",
		EndDateString:   "20230901",
	},
	{
		Company:         "Howl",
		Position:        "Junior Software Engineer",
		StartDateString: "20200501",
		EndDateString:   "20210501",
	},
	{
		Company:         "Narrativ",
		Position:        "Software Engineer Intern",
		StartDateString: "20190101",
		EndDateString:   "20190801",
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
	err := json.NewEncoder(w).Encode(jobs_db)
	if err != nil {
		log.Fatalln("There was an error returning list of all jobs")
	}
}

func JobSingleCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	company := vars["company"]
	var jobs []Job
	jobFound := false
	for _, _job := range jobs_db {
		if _job.Company == company {
			jobs = append(jobs, _job)
			jobFound = true
		}
	}
	if !jobFound {
		// todo: needs to be a better way to check for found job
		// also needs to return an error code
		fmt.Println("Job not found")
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jobs)
	if err != nil {
		log.Fatalln("There was an error returning your job")
	}
}

func JobSingleDate(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	startdate := vars["startdate"]
	var job Job
	jobFound := false
	for _, _job := range jobs_db {
		if _job.StartDateString == startdate {
			job = _job
			jobFound = true
			break
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
