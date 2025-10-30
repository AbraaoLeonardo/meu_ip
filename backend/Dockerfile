FROM golang:tip-alpine3.22 AS builder
WORKDIR /APP
COPY . /APP
RUN go build -o bin/app


FROM alpine:3.22.2
WORKDIR /APP
EXPOSE 8080
COPY --from=builder /APP/bin/app .
ENTRYPOINT ["./app"]
