## golang-mysql-api

[![Build Status](https://travis-ci.org/ChrisTheShark/golang-mysql-api.svg?branch=master)](https://travis-ci.org/ChrisTheShark/golang-mysql-api)

Lead Maintainer - [Chris Dyer](https://github.com/ChrisTheShark)

This project was constructed to demonstrate good organization and testing patterns for a Golang application that uses Mysql for persistance.

## Gettting Started

This repository uses the [dep](https://github.com/golang/dep) tool for dependency management. First, clone this repository and at the root of the project execute ```dep ensure```. This command will go get all dependencies. Next bootstrap a local mysql instance with the included schema.sql file. Provide your connection string in the form ```root:password@tcp(127.0.0.1:3306)/sample``` as an environment variable named MYSQL_HOST. Build or run the application using ```go run *.go``` or ```go build *.go```. If using build follow up with an execution of the created binary. 