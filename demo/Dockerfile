FROM golang:1.19 as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o /demo

FROM ubuntu:18.04 
WORKDIR /
COPY --from=builder /demo /demo
EXPOSE 8080
ENV GIN_MODE=release
ENTRYPOINT ["/demo"]

