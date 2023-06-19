# NOTE: Refer from: 
    - Github: https://go-clean-architecture-example
    - Paper: https://pkritiotis.io/clean-architecture-in-golang/

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
