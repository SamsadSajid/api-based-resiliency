package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/resilience-poc/pay-bill-resilience-service/config"
	"github.com/resilience-poc/pay-bill-resilience-service/router"
	"github.com/resilience-poc/pay-bill-resilience-service/subscriber"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	g := gin.New()
	g.Use(gin.Recovery())
	g.Use(config.InitRedisPool())
	r := router.New(g)

	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", "8000"),
		Handler:        r,
	}

	go func() {
		log.Println("Listening on port 8000")
		if err := server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				log.Println("Closing http server...")
				return
			}
			os.Exit(1)
		}
	}()

	go subscriber.TTLListener()

	signalChan := make(chan os.Signal, 2)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)

	sig := <-signalChan
	log.Println("Received signal: ", sig.String(), " shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Avoid context leak
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("could not gracefully shutdown")
		os.Exit(1)
	}
}
