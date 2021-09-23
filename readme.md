# Simple REST API With Golang
This project is system REST API for data peoples, in this project using database MySQL and programming languange Go Lang


## Structure Database
```
- users
    - id int (primary key)
    - name varchar
    - username varchar
    - password varchar
    - role ENUM('admin','user')
    - date_added datetime
    - date_modified datetime
- peoples
    - nik varchar (primary key)
    - name varchar
    - birthday date
    - gender ENUM('man', 'woman')
    - address varchar
    - photo varchar
    - rt varchar
    - rw varchar
    - subdistrict varchar
    - district varchar
    - city varchar
    - province varchar
    - age int
    - date_added datetime
    - date_modified datetime
```

## Getting Started
### Install Golang

Make sure you have Go 1.13 or higher installed.

https://golang.org/doc/install

### Environment Config

Set-up the standard Go environment variables according to latest guidance (see https://golang.org/doc/install#install).

### Run App
From the project root, run:
```
go mod tidy
go run main.go
```
### API Documentation
https://www.postman.com/collections/90893a6cd48cac68d7bb

## Notes
Before you run this project, you should to create database with name crud_golang

You can use this project database inside the repository github with filename **database.sql**
