package main

import (
	"net/http"
	"os"

	"cloud-opsme.ga/monitoring/agent/internal/app/api"
)

func main() {
	apiTest(":" + os.Args[1])
}

/* Testing
func processTest() {
	fmt.Println(processes.GetList())
}

func metricsTest() {
	disks := []string{"/"}
	fmt.Println(string(monitor.GetMetrics(disks)))
}
*/

func apiTest(port string) {
	http.HandleFunc("/", api.Handler)
	http.HandleFunc("/metrics", api.HandlerMetricsRequest)
	http.HandleFunc("/users", api.HandlerUsersListingRequest)
	http.HandleFunc("/users/list", api.HandlerUsersListingRequest)
	http.HandleFunc("/users/sudoers/add", api.HandlerUsersAddSudoer)
	http.HandleFunc("/users/sudoers/remove", api.HandlerUsersRemoveSudoer)
	http.HandleFunc("/process", api.HandlerProcessesListingRequest)
	http.HandleFunc("/process/list", api.HandlerProcessesListingRequest)
	http.HandleFunc("/process/kill", api.HandlerProcessesKillingRequest)
	http.HandleFunc("/terminal", api.HandlerWebShellCommandRequest)
	http.ListenAndServe(port, nil)
}
