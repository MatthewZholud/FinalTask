FROM golang:latest

WORKDIR app/

COPY TimeTracker .

CMD ["go","run","main.go"]