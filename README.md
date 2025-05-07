# Stock Trades API

A simple RESTful API built with Golang, Gin, and GORM to manage stock trades. The API uses JWT authentication to secure endpoints and allows users to create and retrieve their trades.

---

## Features

- JWT-based user authentication
- Create and retrieve stock trades
- User-level trade access control
- SQLite integration via GORM
- RESTful design using Gin framework

---

## Tech Stack

- **Language**: Go
- **Web Framework**: Gin
- **ORM**: GORM
- **Database**: SQLite
- **Authentication**: JWT

---

Assuming that the user registration will be done by another service and username and password will be stored in db.

For testing purpose i have created a user with details ( Username: "testuser", Password: "testpass" )

after the application starts running( here i am using port 8080 it can be changed using environment variable) we can make a post request for endpoint (http://localhost:8080/login) with body as json (
{
"username": "testuser",
"password": "testpass"
}
)
you will get a jwt token as response

{
"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDY3MjQwODMsInVzZXJfaWQiOjF9.uL7NswacAWpAl4aLTUWLdZrlBvZHPFwd6diFv7Q3u1w"
}

now we can make a trade or get all trades of the user by sending this token in header as ( key : Authorization and value : token)

for making trade make a post request using endpoint (http://localhost:8080/trade) with body
{
"stock": "AAPL",
"quantity": 10,
"trade_type": "BUY"
}

we will get response like

{
"id": 1,
"user_id": 1,
"stock": "AAPL",
"quantity": 10,
"trade_type": "BUY"
}

To see all the trades of a user make a get request to endpoint (http://localhost:8080/trades) by passing the jwt-token in header

we will get a response like

[
{
"id": 1,
"user_id": 1,
"stock": "AAPL",
"quantity": 10,
"trade_type": "BUY"
}
]
