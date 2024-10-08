package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/NhyiraAmofaSekyi/go-webserver/internal/config"
	databaseCfg "github.com/NhyiraAmofaSekyi/go-webserver/internal/db"
	middleware "github.com/NhyiraAmofaSekyi/go-webserver/internal/middleware"
	v1 "github.com/NhyiraAmofaSekyi/go-webserver/internal/v1"
)

// type dbConfig struct {
// 	DB *database.Queries
// }

// var dbCFG *databaseCfg.DBConfig

func main() {

	env := flag.String("env", "development", "Define the application environment (development or production)")
	flag.Parse() // Parse the flags

	// Load environment variables from a .env file if in development environment
	if *env == "development" {
		log.Println("starting dev server")
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found, assuming production settings")
		}
	}
	// err := contracts.Test()
	// if err != nil {
	// 	log.Printf("client err %v", err)
	// }

	port := os.Getenv("PORT")
	host := os.Getenv("HOST")

	start := time.Now()

	router := http.NewServeMux()

	log.Println("server running on port:", port)

	dbCFG, err := databaseCfg.NewDBConfig(*env)
	if err != nil {
		log.Fatal(err)
	}

	// auth.NewAuth2()
	config.Initialize(dbCFG)

	v1 := v1.NewRouter(dbCFG)
	api := "/api/v1/"
	router.Handle(api, http.StripPrefix(strings.TrimRight(api, "/"), v1))
	router.Handle("/metrics", promhttp.Handler())

	stack := middleware.CreateStack(
		middleware.Logging,
		middleware.CorsWrapper,
		middleware.Monitoring,
	)

	server := &http.Server{
		Handler: stack(router),
		Addr:    ":" + port, // Listen address
		// Other configurations like ReadTimeout, WriteTimeout, etc.
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       120 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	// Create a channel to listen for interrupt signals.
	quit := make(chan os.Signal, 1)
	// Register the given channel to receive notifications of the specified signals.
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a goroutine.
	go func() {
		fmt.Println("Server goroutine starting...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %v\n", err)
		}
	}()
	healthEndpoint := "http://" + host + ":" + port + api + "healthz"
	fmt.Println(healthEndpoint)
	go func() {
		for {

			resp, err := http.Get(healthEndpoint)
			if err == nil && resp.StatusCode == http.StatusOK {
				log.Println("Server is ready.")
				elapsed := time.Since(start)
				log.Printf("Server ready in %s", elapsed)
				resp.Body.Close()
				break
			}
			if resp != nil {
				resp.Body.Close()
			}
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Block until a signal is received.
	<-quit
	fmt.Println("Shutting down server...")

	// context with a timeout for the shutdown process.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//gracefully shut down the server.
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("server shutdown failed: %v", err)
	}

	fmt.Println("Server gracefully stopped.")

}
