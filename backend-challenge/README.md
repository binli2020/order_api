# Food Ordering API (Backend Challenge)
The goal of this folder is to build an API server in Golang based on the given OpenAPI spec, including product listing, getting a product, placing an order, and validating promo codes with concurrency.

The project is structured with controller, service, model, router, utils, and generated code from oapi-codegen.

## Features

- Implements the required API endpoints from the OpenAPI spec
- API types and interfaces generated using oapi-codegen
- Promo code validation uses Go concurrency:
    - Each promo file is scanned by a separate goroutine
    - Reads line-by-line (streaming, not loading whole file into memory)
    - Uses channels to collect matches
    - When a code is found in two different files, all goroutines stop immediately


## Folder Structure
```
backend-challenge
├── cmd/server              # Main server entry
├── internal
│   ├── controller
│   ├── router
│   ├── service
│   ├── model
│   ├── utils
│   └── generated           # Code created by oapi-codegen
├── data                    # Default location for promo files (not committed)
└── spec
```

## Promo Code File Location

To avoid committing very large files, the promo code files are not included in the repo and are ignored through .gitignore.

The application looks for promo files in two ways:

1. If $PROMO_DIR environment variable is set

It reads files from:
```
$PROMO_DIR/couponbase1.txt
$PROMO_DIR/couponbase2.txt
$PROMO_DIR/couponbase3.txt
```

2. If $PROMO_DIR is not set

It will look in the project root under:
```
backend-challenge/data/couponbase1.txt
backend-challenge/data/couponbase2.txt
backend-challenge/data/couponbase3.txt
```

Please make sure these three files exist before testing promo code validation.

## How to Run
1. Install dependencies
```
go mod tidy
```

2. (Optional) Set promo file directory
```
export PROMO_DIR="/path/to/promo/files"
```

3. Start server
```
go run cmd/server/main.go
```


Or:
```
./run.sh
```

Server runs on:
```
http://localhost:8080
```

## Test API Endpoints
1. List all products
```
curl -X GET http://localhost:8080/product
```

2. Get a product by ID
```
curl -X GET http://localhost:8080/product/1
```

3. Place an order (with promo code)
```
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{
        "couponCode": "HAPPYHRS",
        "items": [
          { "productId": "1", "quantity": 2 },
          { "productId": "3", "quantity": 1 }
        ]
      }'
```

## Promo Code Concurrency Test

Promo code is valid only if it appears exactly in two or more promo files.

Example test:
```
curl -X POST http://localhost:8080/order \
  -H "Content-Type: application/json" \
  -d '{
        "couponCode": "INVALIDONE",
        "items": [
          { "productId": "1", "quantity": 2 },
          { "productId": "3", "quantity": 1 }
        ]
      }'
```

If invalid:
```
{ "message": "Invalid promo code" }

```

## Concurrency logic highlights
- One goroutine per promo file
- Uses context.Context to cancel remaining goroutines when 2 matches found
- Streamed file reading avoids memory pressure
- Channel used to aggregate results safely
