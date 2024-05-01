# Description

Student API

## Folder Structure

```
└── student_api
    └── db
        └── db.sql
    └── docker-compose.yml
    └── server
        └── cmd
            └── api
                └── db.go
                └── handlers.go
                └── main.go
                └── middlewares.go
                └── routes.go
                └── utils.go
        └── Dockerfile
        └── go.mod
        └── go.sum
        └── internal
            └── models
                └── StudentModel.go
            └── repository
                └── dbrepo
                    └── MySQL_dbrepo.go
                └── repository.go
```

## How to run localy

### Prerequisites

Having docker and docker-compose already install on you machine

Clone the project to your machine

Fill in the information for you database in the `docker-compose.yml` file: `MYSQL_ROOT_PASSWORD`, `MYSQL_USER`, `MYSQL_PASSWORD`, `DB_USER` and `DB_PASSWORD`

```bash
cd student_api
docker-compose up
```
