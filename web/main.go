package main

import (
	"context"
	"expvar"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pamelag/go-web/web/authoring"
	"github.com/pamelag/go-web/web/postgres"
	"github.com/pamelag/go-web/web/rde"
	"github.com/pamelag/go-web/web/server"
)

func main() {

	// Setup repository
	var (
		projects rde.ProjectRepository
	)

	// create connection pool
	pool, err := createConnPool()
	if err != nil {
		panic(err)
	}
	projects = postgres.NewProjectRepository(pool)

	var ds authoring.Service
	ds = authoring.NewService(projects)
	ds = authoring.NewLoggingService(ds)
	ds = authoring.NewInstrumentingService(
		expvar.NewInt("addProject"),
		expvar.NewInt("addFeature"),
		expvar.NewInt("addWireframe"),
		expvar.NewInt("updateWireframeTitle"),
		ds)

	srv := server.New(ds)

	httpServer := &http.Server{Addr: httpAddr,
		Handler:      srv,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second}

	// registers a function to call on Shutdown
	httpServer.RegisterOnShutdown(func() {
		log.Println("Call shutdown hooks")
	})

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGINT)
		signal.Notify(sigint, syscall.SIGTERM)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := httpServer.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed

}
