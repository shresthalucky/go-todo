# Build stage
FROM golang:1.21.4-alpine3.18 AS build-stage

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 go build -o loggerApp ./cmd/api

RUN chmod +x /app/loggerApp

# Run stage
FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=build-stage /app/loggerApp /app
COPY --from=build-stage /app/.env /app

EXPOSE 8080

ENTRYPOINT [ "/app/loggerApp" ]


