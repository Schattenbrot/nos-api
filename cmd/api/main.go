package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Schattenbrot/nos-api/models"
	"github.com/rs/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config
	logger *log.Logger
	models models.Models
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen on.")
	flag.StringVar(&cfg.env, "env", "dev", "Application environment (dev | prod)")
	flag.StringVar(&cfg.db.dsn, "dsn", "mongodb://nos-api-db:27017", "Mongodb dsn to connect to.")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	client, err := openDB(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	db := client.Database("nos-db")

	app := &application{
		config: cfg,
		logger: logger,
		models: models.NewModels(db),
	}

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "UPDATE"},
	})
	handler := c.Handler(app.routes())

	serve := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      handler,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Println("Starting server on port", cfg.port)

	err = serve.ListenAndServe()
	if err != nil {
		log.Println(err)
	}
}

// openDB creates and returns a new client, or an error if it fails
func openDB(cfg config) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.db.dsn))
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
