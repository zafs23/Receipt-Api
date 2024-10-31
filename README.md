## Receipt-Api ![workflow](https://github.com/zafs23/Receipt-Api/actions/workflows/go.yml/badge.svg)
This project handles GET and POST requests. For POST request this project takes a JSON receipt and returns a JSON object with an ID. For the GET request, with provided ID of a receipt returns a JSON response with points earned. 

### Getting Started
Before running this project on your local machine for development and testing, complete the following steps. 

#### Prerequisites
Before running the server, have Go and Git installed on your machine.  The project is built on ```Go version 1.22.0.```
[Go installation offical website](https://go.dev/learn/)
[Git installtion](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

#### Installation
To set up the project, clone the repositoy to your local machine:
```
git clone git@github.com:zafs23/Receipt-Api.git
cd receipt-api
```
Then, run the server using: 
```
go run main.go
```
The API will listen to http://localhost:8000 to handle GET and POST request.

### Services
#### POST request 
 - endpoint: /receipts/process
 - Payload: Receipt JSON
 - Response: JSON containing an id for the receipt.
  example response: 
 ```json
   { "id": "9aedae7d3b59b2ba3b364924bf39f2f797653b4eb4b7fb17ec104bf7ab064b9b" }
 ```
#### GET request 
 - endpoint: /receipts/{id}/points
 - Response: A JSON object containing the number of points awarded.
 example response: 
 ```json
   { "points": 32 }
 ```

 #### Payload Example
 ```json
    {
    "retailer": "Target",
    "purchaseDate": "2022-01-01",
    "purchaseTime": "13:01",
    "items": [
        {
        "shortDescription": "Mountain Dew 12PK",
        "price": "6.49"
        },{
        "shortDescription": "Emils Cheese Pizza",
        "price": "12.25"
        },{
        "shortDescription": "Knorr Creamy Chicken",
        "price": "1.26"
        },{
        "shortDescription": "Doritos Nacho Cheese",
        "price": "3.35"
        },{
        "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
        "price": "12.00"
        }
    ],
    "total": "35.35"
    }
 ```
Expected points:
```
    Total Points: 28
    Breakdown:
        6 points - retailer name has 6 characters
        10 points - 4 items (2 pairs @ 5 points each)
        3 Points - "Emils Cheese Pizza" is 18 characters (a multiple of 3)
                    item price of 12.25 * 0.2 = 2.45, rounded up is 3 points
        3 Points - "Klarbrunn 12-PK 12 FL OZ" is 24 characters (a multiple of 3)
                    item price of 12.00 * 0.2 = 2.4, rounded up is 3 points
        6 points - purchase day is odd
    + ---------
    = 28 points
```

#### Service Logic
These rules collectively define how many points should be awarded to a receipt.

- One point for every alphanumeric character in the retailer name.
- 50 points if the total is a round dollar amount with no cents.
- 25 points if the total is a multiple of 0.25.
- 5 points for every two items on the receipt.
- If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
- 6 points if the day in the purchase date is odd.
- 10 points if the time of purchase is after 2:00pm and before 4:00pm.

### Assumptions and Design scope
Assumptions in the design
- Each short description assumped to be 1-50 character long
- When determining total is multiple of 0.25, the tolerance is assumed to be 0..00001
- Expect at most two decimal places in the item price

##### Scalibility and Memory Usage
To simulate scalibility in the in-memory storage system, data sharding is implemented with highest of 10 shards. To keep the memory usage limited, the generated `ID` is unique for identical payloads. 


### Testing
To execute the automated tests, run the following command from the project directory:
```
go test -v ./tests/...
```
