package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (record *mongoStore) homeHandler(c echo.Context) error {
	zap.S().Info("Executing homeHandler")
	return c.String(http.StatusOK, "Hello, World!")
}
func (record *mongoStore) addHandler(c echo.Context) error {
	zap.S().Info("Executing addHandler")
	name := c.FormValue("name")
	rollnoString := c.FormValue("rollno")
	rollno, err := strconv.Atoi(rollnoString)
	if err != nil {
		zap.S().Error(err)
		fmt.Println(rollnoString)
		//return c.String(http.StatusBadRequest,
		//	http.StatusText(http.StatusBadRequest))
	}
	branch := c.FormValue("branch")
	userid := c.FormValue("userid")
	ch := record.collection.FindOne(record.ctx, bson.M{"roll_no": rollno})
	var exampleResult bson.M
	ch.Decode(&exampleResult)
	if exampleResult != nil {
		return c.String(http.StatusOK, "Student already exists")
	}
	example := bson.D{
		{"name", name},
		{"roll_no", rollno},
		{"branch", branch},
		{"user_id", userid},
	}
	r, err := record.collection.InsertOne(record.ctx, example)
	if err != nil {
		panic(err)
	}
	fmt.Println("Inserted ID:", r.InsertedID)
	return c.JSONPretty(http.StatusOK, example, "  ")
}
func (record *mongoStore) deleteHandler(c echo.Context) error {
	zap.S().Info("Executing deleteHandler")
	rollnoString := c.FormValue("rollno")
	rollno, err := strconv.Atoi(rollnoString)
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest))
	}
	entry := record.collection.FindOne(record.ctx, bson.M{"roll_no": rollno})
	var exampleResult bson.M
	entry.Decode(&exampleResult)
	if exampleResult == nil {
		return c.String(http.StatusOK, "Student does not exist")
	}
	ch, err := record.collection.DeleteOne(record.ctx, bson.M{"roll_no": rollno})
	if err != nil {
		panic(err)
	}

	fmt.Println("Number of items deleted:", ch.DeletedCount)
	return c.String(http.StatusOK, fmt.Sprintf("Deleted record of Roll Number %d", rollno))
}
func (record *mongoStore) filterHandler(c echo.Context) error {
	zap.S().Info("Executing filterHandler")
	parameter := c.FormValue("parameter")
	value := c.FormValue("value")
	var filter interface{}
	if parameter == "userid" {
		parameter = "user_id"
	}
	if parameter == "rollno" {
		parameter = "roll_no"
		valueInt, _ := strconv.Atoi(value)
		filter = bson.D{{parameter, valueInt}}
	} else {
		filter = bson.D{{parameter, value}}
	}
	ch, err := record.collection.Find(record.ctx, filter)
	if err != nil {
		panic(err)
	}
	var lists []bson.M
	if err := ch.All(record.ctx, &lists); err != nil {
		panic(err)
	}
	return c.JSONPretty(http.StatusOK, lists, "  ")
	//return c.JSON(http.StatusOK, lists)
}
func (record *mongoStore) editHandler(c echo.Context) error {
	zap.S().Info("Executing updateHandler")
	rollnoString := c.FormValue("rollno")
	rollno, err := strconv.Atoi(rollnoString)
	if err != nil {
		zap.S().Error(err)
		return c.String(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest))
	}
	entry := record.collection.FindOne(record.ctx, bson.M{"roll_no": rollno})
	var exampleResult bson.M
	entry.Decode(&exampleResult)
	if exampleResult == nil {
		return c.String(http.StatusOK, "Student does not exist")
	}
	id := exampleResult["_id"]
	parameter := c.FormValue("parameter")
	value := c.FormValue("value")
	var filter interface{}
	if parameter == "userid" {
		parameter = "user_id"
	}
	if parameter == "rollno" {
		parameter = "roll_no"
		valueInt, _ := strconv.Atoi(value)
		entry = record.collection.FindOne(record.ctx, bson.M{"roll_no": valueInt})
		var exampleResult bson.M
		entry.Decode(&exampleResult)
		if exampleResult != nil {
			return c.String(http.StatusOK, "Use different Roll No as it already exists")
		}
		filter = bson.M{parameter: valueInt}
	} else {
		filter = bson.M{parameter: value}
	}
	ch, err := record.collection.UpdateOne(record.ctx, bson.M{"roll_no": rollno}, bson.D{{"$set", filter}})
	if err != nil {
		panic(err)
	}
	entry = record.collection.FindOne(record.ctx, bson.M{"_id": id})
	fmt.Println("Number of items updated:", ch.ModifiedCount)
	var result bson.M
	entry.Decode(&result)
	return c.JSONPretty(http.StatusOK, result, "  ")
	//return c.JSON(http.StatusOK, result)
}
func (record *mongoStore) allHandler(c echo.Context) error {
	zap.S().Info("Executing allHandler")
	ch, err := record.collection.Find(record.ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	var lists []bson.M
	if err := ch.All(record.ctx, &lists); err != nil {
		panic(err)
	}
	return c.JSONPretty(http.StatusOK, lists, "  ")
}
