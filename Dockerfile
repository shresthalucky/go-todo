# Build stage
FROM golang:1.18.10-alpine3.17 AS build-stage

WORKDIR /go/src/github.com/shresthalucky/go-todo
COPY . ./

RUN go mod download
RUN go build -o app main.go

# EXPOSE 8080
# CMD ["/go/src/github.com/shresthalucky/go-todo/app"]

# Run stage
FROM alpine:3.17

WORKDIR /
COPY --from=build-stage /go/src/github.com/shresthalucky/go-todo/app .
COPY --from=build-stage /go/src/github.com/shresthalucky/go-todo/.env .

EXPOSE 8080

ENTRYPOINT [ "/app" ]
