package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func (record *studentList) homeHandler(c echo.Context) error {
	zap.S().Info("Executing homeHandler")
	return c.String(http.StatusOK, "Hello, World!")
}
func (record *studentList) addHandler(c echo.Context) error {
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

	if !record.m[rollno].present {
		//record.m[rollno] = true
		//record.info[rollno] = student{name, branch, userid}
		record.m[rollno] = student{true, name, branch, userid}
	} else {
		return c.String(http.StatusOK, "Student already exists")
	}
	//if !record.m[rollno] {
	//	record.m[rollno] = student{name, branch, userid}
	//}
	//record.stdLists = append(record.stdLists, student{rollno, r, branch, userid})
	fmt.Println(record.m)
	return c.String(http.StatusOK, fmt.Sprintf("%d %v", rollno, record.m[rollno]))
	//return c.String(http.StatusOK, fmt.Sprintf("%s %s %d %d", name, branch, rollno, userid))
}
func (record *studentList) deleteHandler(c echo.Context) error {
	zap.S().Info("Executing deleteHandler")
	return c.String(http.StatusOK, "Hello, World!")
}
func (record *studentList) editHandler(c echo.Context) error {
	zap.S().Info("Executing updateHandler")
	return c.String(http.StatusOK, "Hello, World!")
}
func (record *studentList) filterHandler(c echo.Context) error {
	zap.S().Info("Executing filterHandler")
	return c.String(http.StatusOK, "Hello, World!")
}

//func (record *studentList) getUser(c echo.Context) error {
//	zap.S().Info("Executing getUser")
//	// User ID from path `users/:id`
//	id := c.Param("id")
//	return c.String(http.StatusOK, id)
//}
