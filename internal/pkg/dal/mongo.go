package dal

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var (
	m *mongo.Client
)

func MustInitMongo() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	m, err = mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("MONGODB_DSN")).SetMaxPoolSize(20))
	if err != nil {
		panic(fmt.Errorf("mongo connect err:%w", err))
	}

	err = m.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("mongodb initialized!")
}

func GetMongo() *mongo.Client {
	return m
}

func MongoDB() string {
	return viper.GetString("MONGODB_DATABASE")
}
