# Event Booking RestAPI

## Project Details

This is an Event Management application where a user can create an event and make it public for other users to register for the event. JWT-based authentication is used to validate the credentials.

## Libraries Used

- **Gin**: For server and middleware related configurations.
- **PSQL Driver**: For PostgreSQL related operations.
- **JWT**: For user authorization using JWT tokens.
- **Hash**: Passwords are stored as hash values in the database for security reasons.

## Config Setup

### Setup Postgres Config

1. Open `Go HandsOn\GoRestProject\db\db.go` file.
2. Update the constant variables `user` and `password` with your own PostgreSQL username and password.

### JWT Secret Key (Optional)

1. Open `Go HandsOn\GoRestProject\util\jwt.go`.
2. Update the constant variable `secretKey` with your own key.

## Project Setup

1. **Download Go (Version 1.22 or above)** from the official Go page.

2. **Install dependencies**:

   ```sh
   go mod tidy
   ```

3. **Run the application**:
   ```sh
   go run .
   ```

## Postman Endpoints

### Get An Event

```sh
curl --location 'http://localhost:8080/events'
```
