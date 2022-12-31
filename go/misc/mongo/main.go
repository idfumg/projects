package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	client *mongo.Client
}

func NewServer(c *mongo.Client) *Server {
	return &Server{
		client: c,
	}
}

func (s *Server) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	col1 := s.client.Database("catfact").Collection("facts")
	cursor, err := col1.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal(err)
		return
	}
	results := []bson.M{}
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(results)
}

type CatFactWorker struct {
	client *mongo.Client
}

func NewCatFactWorker(c *mongo.Client) *CatFactWorker {
	return &CatFactWorker{
		client: c,
	}
}

func (w *CatFactWorker) start() error {
	col1 := w.client.Database("catfact").Collection("facts")
	ticket := time.NewTicker(2 * time.Second)

	for {
		resp, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			return err
		}
		var fact bson.M // map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&fact); err != nil {
			return err
		}
		resp.Body.Close()
		fmt.Println(fact)

		_, err = col1.InsertOne(context.TODO(), fact)
		if err != nil {
			return err
		}

		<-ticket.C
	}
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	worker := NewCatFactWorker(client)
	go worker.start()

	server := NewServer(client)
	http.HandleFunc("/facts", server.handleGetAllFacts)
	http.ListenAndServe(":8080", nil)
}
