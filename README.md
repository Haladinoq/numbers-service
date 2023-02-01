# Reservation Numbers - API

```
Available commands:
	make postgres               Create docker image for database.
	make createdb               Create the database for the api.
	make dropdb                 Removes the database.
	make migrationup            Run UP migration SQL files.
	make migrationup1           Run ONE UP migration SQL files.
	make migrationdown          Run DOWN migration SQL files.
	make migrationdown2         Run ONE DOWN migration SQL files.
	make swagger                Generate swagger documentation
	make server                 Deploy api.
	make compose                Generates docker container with Api and database.
	make test                   Run test api.
```

## Installation:

```
git clone git@github.com:Haladinoq/numbers-service.git
```

## Project dirs:

| Dir                                |  Description                                                                 |
|:-----------------------------------|:-----------------------------------------------------------------------------|
| build                              |  Packaging and Continuous Integration..                                      |
| cmd                                |  Main applications for this project.                                         |
| config                             |  Configuration file with environment variables.                              |
| docs                               |  Design and user documents, Contains swagger20.json file and swagger client. |
| pkg/config                         |  All the internal configurations of our project.                             |
| pkg/middleware                     |  Middlewares for cors, auth, etc.                                            |
| pkg/rest                           |  REST helpers for use inside pkg/api. Contains pagination utils.             |
| pkg/swagger                        |  All internal swagger configurations.                                        |
| pkg/utils                          |  Internal methods that are useful to the rest of the application.            |
| pkg/numbers/api                    |  All handlers.                                                               |
| pkg/numbers/api/handlers           |  The initializer of our http server and handler of rest requests.            |
| pkg/numbers/api/mapper             |  The mappers functions of our handler.                                       |
| pkg/numbers/api/model              |  The structs data of our requests and responses.                             |
| pkg/numbers/core                   |  All business logic.                                                         |
| pkg/numbers/core/business          |  The business logic of our project.                                          |
| pkg/numbers/core/mappers           |  The mappers functions of our business logic.                                |
| pkg/numbers/core/model             |  The structs data of our DTOs.                                               |
| pkg/numbers/core/services          |  Domain services where business logic is executed.                           |
| pkg/numbers/data/persistence       |  All database entities and data access repositories.                         |
| pkg/numbers/data/persistence/model |  The structs data of our entities.                                           |
| pkg/numbers/data/persistence/repo  |  All interfaces of our database operations.                                  |
| pkg/numbers/data/persistence/sql   |  All database operations.                                                    |
| migrations                         |  SQL migrations.                                                             |

## Setup local development

### Install tools

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```
- [Mockery](https://github.com/vektra/mockery)

    ``` bash
    brew install mockery
    ```
### Setup infrastructure

- Start postgres container:

    ```bash
    make postgres
    ```

- Create numbers database:

    ```bash
    make createdb
    ```

- Run db migration up all versions:

    ```bash
    make migrateup
    ```

- Run db migration up 1 version:

    ```bash
    make migrateup1
    ```

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

- Run db migration down 1 version:

    ```bash
    make migratedown1
    ```
### Documentation

- Generate Api documentation:

    ```bash
    make swagger
    ```

- Access the Api documentation at path 
    ```
  /swagger/index.html
  ```
### How to run

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```


