#

To start the server on localhost:3000:
```
go run cmd/main.go
```

or via docker compose:

```
docker compose up
```

Available endpoints:
```
GET /products - get all products
GET /products/{id} - get a product
POST /availability - get availability
GET /bookings/{id} - get a bookings
POST /bookings - create a booking
POST /bookings/{id}/confirm - confirm a booking
```

Use `Capability: pricing` header for response with prices, see examples below.

Examples:

```
# get products
curl -X GET -H "Capability: pricing" http://localhost:3000/products

# get a product
curl -X GET -H "Capability: pricing" http://localhost:3000/products/foo

# get availability
curl -X POST -H "Capability: pricing" http://localhost:3000/availability -d '{"product_id": "foo", "local_date_start": "2024-05-11", "local_date_end": "2024-05-12"}'

# create a booking
curl -X POST -H "Capability: pricing" http://localhost:3000/bookings -d '{ "product_id": "foo", "availability_id": "93a20721-1e06-4aeb-9e3a-a7437800f6d8", "units": 3}'

# confirm a booking
curl -X POST http://localhost:3000/bookings/0db49e4e-df54-430a-b388-617d3aac883d/confirm

# get a booking
curl -X GET http://localhost:3000/bookings/0db49e4e-df54-430a-b388-617d3aac883d

```

To run unit tests:
```
go test ./internal/... -v
```


To run integration tests (this commands expects the server to be running):
```
go test ./test/... -v
```

## Implementation notes
* The lifecycle for a product is not specified which means there is no way to create a product (via exposed endpoints). Without a product, it is not possible to check product availability or create a booking, so I seed 2 products into the data store on start with ids `foo` and `bar` which can be thought of as 'default' or 'initial' offering. While seeding data can be useful in some cases, we probably want a well-defined user story for managing products in production scenarios.
