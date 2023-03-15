package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"github.com/e-berman/baseball_api/handlers"
)

func main() {
	l := log.New(os.Stdout, "baseball-api", log.LstdFlags)
	ph := handlers.NewPlayer(l)

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	s := &http.Server{
		Addr: ":4242",
		Handler: sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	
	sig := <- sigChan
	l.Println("Server has been gracefully shutdown.", sig)
	
	tc, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	s.Shutdown(tc)
}
