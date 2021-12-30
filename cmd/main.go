package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	port := os.Getenv("PORT")

	router := http.NewServeMux()

	mg, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongo://localhost:27017/test"))

	if err != nil {
		fmt.Println(err)
	}

	router.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

		cx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)

		defer cancel()

		list, err := mg.ListDatabaseNames(cx, bson.M{}, options.ListDatabases().SetNameOnly(true))
		rw.WriteHeader(200)

		fmt.Println(err)

		json.NewEncoder(rw).Encode(list)

	})

	http.ListenAndServe(":"+port, router)
}
