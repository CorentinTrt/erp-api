// This is the main entry
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	A "service-users-data/_v1"
	DB "service-users-data/internals/database"
	U "service-users-data/internals/utils"
)

func main() {
	U.Log("Service started", nil)

	err := DB.InitDatabase()
	if err != nil {
		panic(err)
	} else {
		U.Log("Connected to the DB", nil)
	}

	l := log.New(os.Stdout, "service-user-data - ", log.LstdFlags)

	ro := A.InitRouter()

	s := &http.Server{
		Addr:         ":9090",
		Handler:      ro,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
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

	// Wait for an available signal
	// Then print the message into the channel
	sig := <-sigChan
	U.Log("Recieved terminated, gracefully shutdown", sig)

	ctxTimeOut, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(ctxTimeOut)
}
