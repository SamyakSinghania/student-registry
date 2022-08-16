package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
	"log"
)

type mongoStore struct {
	ctx               context.Context
	client            *mongo.Client
	collection        *mongo.Collection
	ticketsCollection *mongo.Collection
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			log.Fatal(err)
		}
	}(logger)
	zap.ReplaceGlobals(logger)

	record := new(mongoStore)
	record.ctx = context.TODO()
	opts := options.Client().ApplyURI("mongodb+srv://database1:mongodbpwd@cluster0.pq79yox.mongodb.net/?retryWrites=true&w=majority")

	record.client, err = mongo.Connect(record.ctx, opts)
	if err != nil {
		panic(err)
	}

	defer record.client.Disconnect(record.ctx)

	fmt.Printf("%T\n", record.client)

	testDB := record.client.Database("test")
	fmt.Printf("%T\n", testDB)

	record.collection = testDB.Collection("example")

	fmt.Printf("%T\n", record.collection)
	e := echo.New()
	e.GET("/", record.homeHandler)
	e.POST("/add", record.addHandler)
	e.POST("/delete", record.deleteHandler)
	e.POST("/filter", record.filterHandler)
	e.POST("/edit", record.editHandler)
	e.GET("/all", record.allHandler)
	e.Logger.Fatal(e.Start(":4000"))

}
