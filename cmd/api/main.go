package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/egamorim/api-test/cmd/api/handlers"
	"github.com/egamorim/api-test/cmd/api/routers"
)

var port string

func init() {
	port = "8000"
}

func main() {
	h := &handlers.Handler{}
	routes := routers.Router(h)

	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      routes,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}

	log.Printf("Api test is running at port %s\n", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("can't initialize Server", err)
		os.Exit(1)
	}
}
