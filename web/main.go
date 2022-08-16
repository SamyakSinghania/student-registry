package main

import (
	"github.com/labstack/echo/v4"
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

type studentList struct {
	//stdLists []student
	m map[int]student
	//info map[int]student
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
	record := new(studentList)
	record.m = make(map[int]student)
	e := echo.New()
	e.GET("/", record.homeHandler)
	//e.GET("/add", record.addHandler)
	e.POST("/add", record.addHandler)
	//e.GET("/delete", record.deleteHandler)
	e.POST("/delete", record.deleteHandler)
	//e.GET("/update", record.updateHandler)
	e.POST("/edit", record.editHandler)
	//e.GET("/filter", record.filterHandler)
	e.POST("/filter", record.filterHandler)

	//e.POST("/", record.homeHandler)
	//e.GET("/users/:id", record.getUser)
	e.Logger.Fatal(e.Start(":8080"))
	//e.Start(":8080")

}
