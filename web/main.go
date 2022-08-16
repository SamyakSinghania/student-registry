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

//Create a dictionary to store name, roll number, branch and user id for all the students.
type student struct {
	present bool
	Name    string
	Branch  string
	UserID  string
}
type mongoStore struct {
	ctx               context.Context
	client            *mongo.Client
	collection        *mongo.Collection
	ticketsCollection *mongo.Collection
}

// Receive add, delete, update, and get request from client
//func addStudent(c echo.Context) error {
//	// Read JSON input
//	var input student
//	if err := c.Bind(&input); err != nil {
//		return err
//	}
//	// Add to database
//	return c.JSON(http.StatusOK, input)
//}

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
	//defer exampleCollection.Drop(ctx)

	fmt.Printf("%T\n", record.collection)
	e := echo.New()
	e.GET("/", record.homeHandler)
	e.POST("/add", record.addHandler)
	e.POST("/delete", record.deleteHandler)
	e.POST("/filter/:parameter", record.filterHandler)
	e.POST("/edit", record.editHandler)
	//e.POST("/edit/:rollno", record.editHandler1)
	//e.POST("/edit/:parameter", record.updateHandler)
	e.Logger.Fatal(e.Start(":4000"))
	//e.Start(":8080")

}
