# Student Registry
## Description
The following backend has been created using Echo Web Framework in Golang and the database used is MongoDB.
The backend has been created to store the details of the students and their respective details.

The following are the features of the backend:
* Add a new student entry
* Delete an existing student entry
* Filter all the student entries based on a particular parameter
* Update the details of a student

A student entry consists of the following parameters:
* Name
* Roll Number
* Branch
* UserID

## Requirements
* Golang 
* POSTMAN 
* MongoDB and Echo Web Framework libraries installed in Go\
**To run this server, one does not need MongoDB installed on the system as the database is stored on the cloud.**

## Routes
* Homepage (http://localhost:4000)  (GET)
* Add a new student entry (http://localhost:4000/add)  (POST)
* Delete an existing student entry (http://localhost:4000/delete)  (POST)
* Filter all the student entries based on a particular parameter (http://localhost:4000/filter)  (POST)
* Update the details of a student (http://localhost:4000/edit)  (POST)
* Get all the student entries (http://localhost:4000/all)  (GET)

To start the server, run the following command:
```shell
go run web/*.go
```

## POST Requests to the server
To make a POST request to the server, POSTMAN is used. \
[![Run in Postman](https://run.pstmn.io/button.svg)](https://god.gw.postman.com/run-collection/22128788-c6445aea-3acd-43a4-b53f-9f8aac12db36?action=collection%2Ffork&collection-url=entityId%3D22128788-c6445aea-3acd-43a4-b53f-9f8aac12db36%26entityType%3Dcollection%26workspaceId%3D41aa9d90-9dda-47ab-88ac-d573fd655c8c)
Use the above link to send key value pair for any of the routes.\
To edit any key value pair \
Head to **BODY->form-data->key-value pair** and change the value of the key.
The name of the keys are as follows:
* **name**
* **rollno**
* **branch**
* **userid**




