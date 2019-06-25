package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"cloud-opsme.ga/monitoring/agent/internal/app/monitor"
	"cloud-opsme.ga/monitoring/agent/internal/app/processes"
	"cloud-opsme.ga/monitoring/agent/internal/app/users"
	"cloud-opsme.ga/monitoring/agent/internal/app/webshell"
)

type ReturnError struct {
	StatusCode int
	Body       string
}

type RequestKillingProcess struct {
	Pid string
}

type RequestExecCommand struct {
	Command string
}

type RequestSudoer struct {
	Uid string
}

// Handler is the api based message for testing
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(`{"message": "hello dev"}`))
	}
}

//********* PROCESSES ***************** //

// HandlerProcessesListingRequest is an endpoint for the orchestrator
// It will call GetList() from processes/process package and return an json
func HandlerProcessesListingRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(processes.GetList()))
	}
}

// HandlerProcessesKillingRequest is an endpoint for the orchestrator
// It will call KillProcess() from processes/process package and return an json
func HandlerProcessesKillingRequest(w http.ResponseWriter, r *http.Request) {
	pid := r.FormValue("pid")

	if pid == "" {
		w.Header().Set("content-type", "application/json")
		r, _ := json.Marshal(ReturnError{
			StatusCode: 1,
			Body:       "empty request: need a pid",
		})
		w.Write(r)
	} else {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(processes.KillProcess(pid)))
	}
}

//********* USERS ***************** //

// HandlerUsersListingRequest is an endpoint for the orchestrator
// It will call GetList from users/users package and return a json
func HandlerUsersListingRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(users.GetList()))
	}
}

// HandlerUsersAddSudoer is an endpoint for the web-app
// It will call AddSudoer from users/users package and return a json
func HandlerUsersAddSudoer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()
		var req RequestSudoer
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Fatal(err)
		}

		if req.Uid == "" {
			w.Header().Set("content-type", "application/json")
			r, _ := json.Marshal(ReturnError{
				StatusCode: 1,
				Body:       "no user uid defined",
			})
			w.Write(r)
		} else {
			w.Header().Set("content-type", "application/json")
			w.Write(webshell.Command(req.Uid))
		}
	}
}

// HandlerUsersRemoveSudoer is an endpoint for the web-app
// It will call RemoveSudoer from users/users package and return a json
func HandlerUsersRemoveSudoer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()
		var req RequestSudoer
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(req.Uid)

		if req.Uid == "" {
			w.Header().Set("content-type", "application/json")
			r, _ := json.Marshal(ReturnError{
				StatusCode: 1,
				Body:       "no user uid defined",
			})
			w.Write(r)
		} else {
			w.Header().Set("content-type", "application/json")
			w.Write(webshell.Command(req.Uid))
		}
	}
}

//********* METRICS ***************** //

// HandlerMetricsRequest is an endpoint for the orchestrator
// It will call GetMetrics from monitor/metrics package and return a json
func HandlerMetricsRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		disks := []string{r.FormValue("disk")}
		w.Header().Set("content-type", "application/json")
		w.Write([]byte(monitor.GetMetrics(disks)))
	}
}

//********* WEB TERMINAL ***************** //

// HandlerWebShellCommandRequest is an endpoint for the webapp
// It will coll Command from webshell/webshell package and return a json
func HandlerWebShellCommandRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()
		var req RequestExecCommand
		err = json.Unmarshal(body, &req)
		if err != nil {
			log.Fatal(err)
		}

		if req.Command == "" {
			w.Header().Set("content-type", "application/json")
			r, _ := json.Marshal(ReturnError{
				StatusCode: 1,
				Body:       "empty line",
			})
			w.Write(r)
		} else {
			w.Header().Set("content-type", "application/json")
			w.Write(webshell.Command(req.Command))
		}
	}
}
