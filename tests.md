# Testing space-trouble API

Below are the active launchpad ids on space-X API
5e9e4501f509094ba4566f84
5e9e4502f509092b78566f87
5e9e4502f509094188566f88

Use one of the above launchpad ids for validation to be successful

Below are the supported Destination Ids
Mars
Moon
Pluto
Asteroid Belt
Europa
Titan
Ganymede

Use one of the above destination ids for validation to be successful

# Testing with curl

## Create Booking
```
curl http://localhost:8080/api/v1/bookings/create --header 'Content-Type: application/json' -d "{\"first_name\":\"abc\",\"last_name\":\"def\",\"gender\":\"M\",\"birthday\":\"1996-01-02\",\"launchpad_id\":\"5e9e4501f509094ba4566f84\",\"destination_id\":\"Titan\",\"launch_date\":\"2026-01-03\"}"
```
This will return booking-id if successful.

## Get All Bookings
```
curl http://localhost:8080/api/v1/bookings/all
```
Returns all the bookings in Json.

## Get All Bookings
```
curl -X DELETE http://localhost:8080/api/v1/bookings/<Booking-ID>
```
Will return success if booking id is deleted.

