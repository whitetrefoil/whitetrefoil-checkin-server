package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"whitetrefoil.com/checkin/server"
)

func main() {

	id := flag.String("id", "", "4SQ's Client ID")
	sec := flag.String("sec", "", "4SQ's Client Secret")
	red := flag.String("red", "", "Redirect URI")
	port := flag.Int("port", 3333, "Port to listen")

	flag.Parse()

	if *id == "" || *sec == "" || *red == "" {
		log.Println("Usage: checkin -port=<server_port=3333> -id=<client_id> -sec=<secret_id> -red=<redirect_uri>")
		os.Exit(255)
	}

	srv := server.NewServer(*port, *id, *sec, *red)

	exited := make(chan error)
	sigint := make(chan os.Signal)
	go func() {
		<-sigint
		timeoutCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		exited <- srv.Shutdown(timeoutCtx)
	}()
	signal.Notify(sigint, os.Interrupt)

	srvErr := srv.ListenAndServe()
	if srvErr != nil && srvErr != http.ErrServerClosed {
		log.Printf(srvErr.Error())
	}

	err := <-exited
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Exit...")
}
