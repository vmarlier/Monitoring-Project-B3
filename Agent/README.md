# Agent

The Agent need to be installed on the host you want to monitor

The main package is in `cmd/`

The private libraries are in `internal/`

You have to build librairies each time you want to update the code:
`$ go build internal/app/monitor/metrics.go`

Same for the main package
`$ go build cmd/main.go`

To run the entire application
`$ go build cmd/main.go`
`$ ./cmd/main`

To test a run quickly
`$ go run cmd/main.go`

To develop on a library without having to compile all the project you can
replace the package name of the file you editing by `main`.
Don't forget to pull it back when you commit (the package must be the name of the folder your file is in)

To create the docker image
`$ docker built -t agent .`

To run the docker container
`$ docker run -d -p 8001:8001 agent`

You now can use docker compose to run the agent easily
`$ docker-compose up -d`
