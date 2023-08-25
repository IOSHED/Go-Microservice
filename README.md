
## Run project

```bach
go run cmd/microservice/main.go --config "config/local.yaml"
```

## Migrations manually

migrat create:

```bach
migrate create -ext sql -dir migrations -seq [name_migrat]
```

up migrations:

```bach
migrate -path migrations -database "mysql://root:44541@tcp(localhost:3306)/go-task-manager?parseTime=true&multiStatements=true" up
```

down migrations:

```bach
migrate -path migrations -database "$STORAGE_PATH" down
```

## Dependencies

* **cleanenv** v1.5.0 load config
* **go-sqlite3** v1.14.17 init storage sqlite3
* **golang-migrate/migrate** v3.5.4+incompatible 
* **echo/v4** v4.11.1 router
