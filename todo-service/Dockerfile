FROM golang:1.21.4-alpine3.18 AS build-stage

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod download
RUN CGO_ENABLED=0 go build -o todoApp ./cmd/api

RUN chmod +x /app/todoApp

# Run stage
FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=build-stage /app/todoApp /app
COPY --from=build-stage /app/.env /app

EXPOSE 8081

ENTRYPOINT [ "/app/todoApp" ]


