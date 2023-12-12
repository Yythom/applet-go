package bootstrap

import (
	"context"
	"fmt"
	"log"
	"test/mongo"
	"time"
)

func NewMongoDatabase(env *Env) mongo.Client {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbHost := env.DBHost
	dbPort := env.DBPort
	dbUser := env.DBUser
	dbPass := env.DBPass

	var connectionString string

	if dbUser != "" && dbPass != "" {
		connectionString = fmt.Sprintf("mongodb://%s:%s@%s:%s", dbUser, dbPass, dbHost, dbPort)
	} else {
		connectionString = fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	}

	log.Println(connectionString)

	client, err := mongo.NewClient(connectionString)
	if err != nil {
		log.Fatalf("connect error is %e/n", err)
	}

	err = client.Ping(context.TODO())
	if err != nil {
		log.Fatalf("ping error is %e/n", err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
