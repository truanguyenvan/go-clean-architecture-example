#### 👨‍💻 Full list what has been used:
* [go](https://go.dev/dl/) 1.18 or later, to use generics
* [wire](https://github.com/google/wire) for Dependency injection
* [fiber](https://github.com/gofiber/fiber) - Web framework
* [sqlx](https://github.com/jmoiron/sqlx) - Extensions to database/sql.
* [pgx](https://github.com/jackc/pgx) - PostgreSQL driver and toolkit for Go
* [viper](https://github.com/spf13/viper) - Go configuration with fangs
* [go-redis](https://github.com/go-redis/redis) - Type-safe Redis client for Golang
* [zap](https://github.com/uber-go/zap) - Logger
* [validator](https://github.com/go-playground/validator) - Go Struct and Field validation
* [uuid](https://github.com/google/uuid) - UUID
* [migrate](https://github.com/golang-migrate/migrate) - Database migrations. CLI and Golang library.
* [minio-go](https://github.com/minio/minio-go) - AWS S3 MinIO Client SDK for Go
* [swag](https://github.com/swaggo/swag) - Swagger
* [testify](https://github.com/stretchr/testify) - Testing toolkit
* [gomock](https://github.com/golang/mock) - Mocking framework
* [Docker](https://www.docker.com/) - Docker

# Code Design
![Code Design](./docs/graphics/clean-architecture-go-climb.png)

# Application Overview

### Use Cases
As a **web client**, I want to be able to
* Get all available crags.
* Get a crag by ID.
* Add a crag by providing a name, country, and description.
* Update a crag by providing a name, country, and description.
* Remove a crag by ID.

As an **application administrator**
* When a new crag is added, I want to receive a notification at a pre-agreed channel.

### Technical requirements
* Operations should be exposed via an HTTP restful interface.
* For *simplicity* purposes,
    * The crags should be stored in memory; no need for persistence storage.
    * Notifications should be sent in a console application.

**Project Structure**
- `cmd` contains the `main.go` file, the entry point of the application
- `docs` contains documentation about the application
- `internal` contains the main implementation of our application. It consists of the three layers of clean architecture + server 
    - infra
        - outputadapters
        - inputports
    - app
    - domain
    - server
- `pkg` shared utility code

  Each of these directories contains its corresponding components, following the group-by-feature approach.
- `vendor` contains the dependencies of our project


# Developer's Handbook
```makefile
make run  ## Run the application
make lint  ## Perform linting
make test  ## Run unit tests
make build  ## Build the app executable for Linux
make fmt  ## Format the source code
```

# Inspirations
[Clean Architecture in Go ](https://pkritiotis.io/clean-architecture-in-golang) \
[Clean Architecture Uncle Bob](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)