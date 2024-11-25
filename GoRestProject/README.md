# Event Booking RestAPI 🎟️

## Project Details 📋

This is an Event Management application where a user can create an event and make it public for other users to register for the event. JWT-based authentication is used to validate the credentials.

## Libraries Used 📚

- **Gin**: For server and middleware related configurations.
- **PSQL Driver**: For PostgreSQL related operations.
- **JWT**: For user authorization using JWT tokens.
- **Hash**: Passwords are stored as hash values in the database for security reasons.

## Config Setup ⚙️

### Setup Postgres Config 🛠️

1. Open `Go HandsOn\GoRestProject\db\db.go` file.
2. Update the constant variables `user` and `password` with your own PostgreSQL username and password.

### JWT Secret Key (Optional) 🔑

1. Open `Go HandsOn\GoRestProject\util\jwt.go`.
2. Update the constant variable `secretKey` with your own key.

## Project Setup 🚀

1. **Download Go (Version 1.22 or above)** from the official Go page.

2. **Install dependencies**:

   ```sh
   go mod tidy
   ```

3. **Run the application**:
   ```sh
   go run .
   ```

## Postman Endpoints 📬

### Get An Event

```sh
curl --location 'http://localhost:8080/events'
```

### Get An Event by Id

```sh
curl --location 'http://localhost:8080/events/1'
```

### Create an Event

```sh
curl --location 'http://localhost:8080/events' \
--header 'Content-Type: application/json' \
--data '{
    "name" :"Coldplay Event1",
    "description" : "Foreginer Band",
    "location" : "Pune",
    "dateTime" : "2025-01-01T15:30:00.000Z"
}'
```

### Update an Event

```sh
curl --location --request PUT 'http://localhost:8080/events/5' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1cmFqQGdtYWlsLmNvbSIsImV4cCI6MTczMjUzNzgzMiwidXNlcklkIjo0fQ.DvffddiOT9pkPnc-BOrJ1euockG8eVlX4-6RdPeexFE' \
--header 'Content-Type: application/json' \
--data '{
    "name" :"Coldplay Event1",
    "description" : "Foreginer Band Canceller",
    "location" : "Pune",
    "dateTime" : "2025-01-01T15:30:00.000Z"
}'
```

### Delete an Event

```sh
curl --location --request DELETE 'http://localhost:8080/events/3' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuYW5kQGdtYWlsLmNvbSIsImV4cCI6MTczMjUzNDU4NiwidXNlcklkIjozfQ.zYV4h5rIsGykaiI4MTLbh_69qQst2GNbAHzxF6AlyaM'
```

### User Registration

```sh
curl --location 'http://localhost:8080/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "suraj@gmail.com",
    "password" : "changeit"
}'
```

### User Login (Generates JWT Token)

```sh
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "suraj@gmail.com",
    "password" : "changeit"
}'
```

### User Registration for an event

```sh
curl --location --request POST 'http://localhost:8080/events/6/register' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1cmFqQGdtYWlsLmNvbSIsImV4cCI6MTczMjU0MjIzNCwidXNlcklkIjo0fQ.De7NJ5vQ59mEJzRRv_7Ykfld6UvUdyN_-fd83H5ho3w'

```

### Delete User Registration for an event

```sh
curl --location --request DELETE 'http://localhost:8080/events/6/register' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1cmFqQGdtYWlsLmNvbSIsImV4cCI6MTczMjU0MjIzNCwidXNlcklkIjo0fQ.De7NJ5vQ59mEJzRRv_7Ykfld6UvUdyN_-fd83H5ho3w'
```
