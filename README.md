# space-trouble

A microservice application

Code Organization

    - cmd - Main & Binaries
        - bin
            - space_controller
    - config - config data
        - .env
    - internal
        - competitors - Data from Competitors
        - handlers - API Endpoint handlers
        - pkg
            - apicalls - Generic wrappers for external API calls
            - logger - zap logger wrapper
            - middleware - middleware (can be used for auth, code commented)
            - model - Bookings model and validations
            - response - Endpoint Responses
        - repository - DB Calls
        - server - server and routes
        - service - Endpoint associated services
        - sql - schemas
    docker-compose.yaml
    Dockerfile
    go.mod
    go.sum
    Makefile
    README.md

## Build

```bash
  make build
```

## Run Locally
Assumption - postgres DB is already running locally
```bash
  cp .sample.env .env
  cmd/bin/space-trouble
```

## docker-componse run

```bash
  docker-compose up
```

## Running Tests

To run tests, run the following command

```bash
  make test
```

## Working

    Application loads the config and connects to the postgres db.
    Mounts the below endpoints
        - /api/v1/bookings/create - POST
        - /api/v1/bookings/all - GET
        - /api/v1/bookings/:id - DELETE  
    When create booking is called
        - checks if the input parameters are valid
            - destination is valid if it is one of the configured destinations(Mars, Moon, Pluto, Asteroid Belt, Europa, Titan, Ganymede)
            - launchpad is valid if it is one of the "active" launchpads that can be obtained from https://api.spacexdata.com/v4/launchpads
            - valid launch_date and birthday
        - checks if there's a conflicting booking
        - checks if there's a competing flight in space-X by querying https://api.spacexdata.com/v5/launches/query
        - if the checks are successful then a booking is saved into the DB
        - booking id is returned
        - if any of the validations or checks fail, error response of booking couldn't be created is sent back as response
    When Get all the bookings is called
        - Reads the DB and returns all the bookings
    When delete a booking is called
        - If the booking with the booking id exists in the DB the booking is deleted, else error response of booking not found is sent back as response
    
