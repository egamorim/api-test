FROM golang:1.10
ENV GOPATH=/go
WORKDIR /go
RUN mkdir -p src/github.com/egamorim/api-test
COPY . /go/src/github.com/egamorim/api-test
WORKDIR /go/src/github.com/egamorim/api-test
RUN go get -u github.com/gorilla/mux
RUN make build
FROM debian:jessie-slim
RUN apt-get update && apt-get install ca-certificates -y && rm -rf /var/lib/apt/lists/*
WORKDIR /root/
COPY --from=0 /go/src/github.com/egamorim/api-test .
CMD ["./api-test"]