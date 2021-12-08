package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Schattenbrot/nos-api/config"
	"github.com/Schattenbrot/nos-api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	flag.IntVar(&config.Cfg.Port, "port", 4000, "Server port to listen on.")
	flag.StringVar(&config.Cfg.Env, "env", "dev", "Application environment (dev | prod)")
	flag.StringVar(&config.Cfg.DB.DSN, "dsn", "mongodb://nos-api-db:27017", "Mongodb dsn to connect to.")
	flag.Parse()

	const version = "1.0.0"

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	client, err := openDB(config.Cfg)
	if err != nil {
		logger.Fatal(err)
	}
	db := client.Database("nos-db")

	config.App = &config.Application{
		Version: version,
		Logger:  logger,
		Models:  models.NewModels(db),
	}

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Cfg.Port),
		Handler:      chiRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", config.Cfg.Port)

	err = serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

// openDB creates and returns a new client, or an error if it fails
func openDB(cfg config.Config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.DB.DSN))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}
