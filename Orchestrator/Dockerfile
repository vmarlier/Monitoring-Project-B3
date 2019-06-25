FROM golang:1.12-stretch
MAINTAINER <valentin.marlier@gmail.com>

WORKDIR /opt/orchestrator

COPY . .
COPY go.mod .

RUN mkdir configs

# allow golang to use modules
RUN export GO111MODULE=on 

RUN go build -o cmd/monitoring-orchestrator cmd/main.go

CMD cmd/monitoring-orchestrator