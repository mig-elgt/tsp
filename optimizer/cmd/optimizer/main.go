package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/mig-elgt/tsp/optimizer/handler"
	"github.com/mig-elgt/tsp/optimizer/table"
	"github.com/mig-elgt/tsp/optimizer/vns"
	log "github.com/sirupsen/logrus"
)

func main() {
	port := flag.Int("p", 8080, "service port")
	tableSvcHost := flag.String("table", "table:8080", "table service addr")
	vnsSvcHost := flag.String("vns", "vns:8080", "vns service addr")
	flag.Parse()

	tableSvc, err := table.New(*tableSvcHost)
	if err != nil {
		log.Fatalf("could not create table service: %v", err)
	}
	vnsSvc, err := vns.New(*vnsSvcHost)
	if err != nil {
		log.Fatalf("could not create vns service: %v", err)
	}

	h := handler.New(tableSvc, vnsSvc)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", *port),
		Handler: h,
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	log.Infof("Server running at localhost: %v", *port)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
