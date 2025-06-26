package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB() {

	clientOptions := options.Client().ApplyURI("mongodb://mongo:27017")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Erro ao criar cliente MongoDB: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Erro ao conectar ao MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Erro ao pingar o MongoDB: %v", err)
	}

	log.Println("✅ Conectado ao MongoDB com sucesso!")

	DB = client.Database("librarydb")
}
