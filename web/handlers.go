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
		return c.String(http.StatusBadRequest,
			http.StatusText(http.StatusBadRequest))
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
	fmt.Println(r.InsertedID)
	return c.String(http.StatusOK, fmt.Sprintf("Added record %v", example))
	//if !record.m[rollno].present {
	//	//record.m[rollno] = true
	//	//record.info[rollno] = student{name, branch, userid}
	//	record.m[rollno] = student{true, name, branch, userid}
	//} else {
	//	return c.String(http.StatusOK, "Student already exists")
	//}
	//if !record.m[rollno] {
	//	record.m[rollno] = student{name, branch, userid}
	//}
	//record.stdLists = append(record.stdLists, student{rollno, r, branch, userid})
	//fmt.Println(record.m)
	//return c.String(http.StatusOK, fmt.Sprintf("%d %v", rollno, record.m[rollno]))
	//return c.String(http.StatusOK, fmt.Sprintf("%s %s %d %d", name, branch, rollno, userid))
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
	return c.String(http.StatusOK, fmt.Sprintf("Deleted record %d", rollno))
}
func (record *mongoStore) filterHandler(c echo.Context) error {
	zap.S().Info("Executing filterHandler")
	//parameter := c.QueryParam("parameter")
	//parameter := c.FormValue("parameter")
	parameter := c.Param("parameter")
	value := c.FormValue("parameter")
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
	//name := c.FormValue("name")
	//rollnoString := c.FormValue("rollno")
	//rollno, err := strconv.Atoi(rollnoString)
	//if err != nil {
	//	zap.S().Error(err)
	//	return c.String(http.StatusBadRequest,
	//		http.StatusText(http.StatusBadRequest))
	//}
	//branch := c.FormValue("branch")
	//userid := c.FormValue("userid")
	//filter := bson.D{{parameter, value}}

	ch, err := record.collection.Find(record.ctx, filter)
	if err != nil {
		panic(err)
	}
	var lists []bson.M
	if err := ch.All(record.ctx, &lists); err != nil {
		panic(err)
	}
	//var result []string
	//for _, v := range lists {
	//	result = append(result, fmt.Sprintf("%v + \n", v))
	//}
	//return c.String(http.StatusOK, fmt.Sprintf("%s", result))
	return c.JSON(http.StatusOK, lists)
	//return c.String(http.StatusOK, "Hello, World!")
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
	return c.JSON(http.StatusOK, result)
}

//func (record *mongoStore) updateHandler(c echo.Context) error {
//	zap.S().Info("Executing updateHandler")
//
//}

//func (record *studentList) getUser(c echo.Context) error {
//	zap.S().Info("Executing getUser")
//	// User ID from path `users/:id`
//	id := c.Param("id")
//	return c.String(http.StatusOK, id)
//}
