# Event Booking RestAPI

## Project Details

This an an Event Management application where a user can create an event and make to public for other users to register for the event.
JWT Based Authentication is used to validate the credentials

## Libraries Used

Gin -> For server and middleware related configurations.
PSQL Driver -> For PostgreSQL related operations.
JWT -> For user authorization JWT token is used.
Hash -> Password is stored as an Hash value in DB for security reasons.

## Config Setup

### Setup Postgres Config

    Open Go HandsOn\GoRestProject\db\db.go file
    Update constant variable "user" and "password" with your own Postgres Username and Password

## JWT Secrete Key (Optional)

Open \Go HandsOn\GoRestProject\util\jwt.go
Update constant variable "secretKey" with your own key

## Project SetUp

Step1 : Download Go( Version above 1.22) from the below Go official Page
https://go.dev/doc/install

Step2: Run
go mod tidy
This will install all the project dependency

go run .
This will start the application

## PostMan Endpoints

Get An Event

curl --location 'http://localhost:8080/events'

Get Event By ID

curl --location 'http://localhost:8080/events/1'

Create an Event

curl --location 'http://localhost:8080/events' \
--header 'Content-Type: application/json' \
--data '{
"name" :"Coldplay Event1",
"description" : "Foreginer Band",
"location" : "Pune",
"dateTime" : "2025-01-01T15:30:00.000Z"
}'

Update an Event

curl --location --request PUT 'http://localhost:8080/events/5' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1cmFqQGdtYWlsLmNvbSIsImV4cCI6MTczMjUzNzgzMiwidXNlcklkIjo0fQ.DvffddiOT9pkPnc-BOrJ1euockG8eVlX4-6RdPeexFE' \
--header 'Content-Type: application/json' \
--data '{
"name" :"Coldplay Event1",
"description" : "Foreginer Band Canceller",
"location" : "Pune",
"dateTime" : "2025-01-01T15:30:00.000Z"
}'

Delete an Event

curl --location --request DELETE 'http://localhost:8080/events/3' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImFuYW5kQGdtYWlsLmNvbSIsImV4cCI6MTczMjUzNDU4NiwidXNlcklkIjozfQ.zYV4h5rIsGykaiI4MTLbh_69qQst2GNbAHzxF6AlyaM'

User Registration

curl --location 'http://localhost:8080/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
"email" : "suraj@gmail.com",
"password" : "changeit"
}'

User Login

curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
"email" : "suraj@gmail.com",
"password" : "changeit"
}'

User Registration for an event

curl --location --request POST 'http://localhost:8080/events/6/register' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1cmFqQGdtYWlsLmNvbSIsImV4cCI6MTczMjU0MjIzNCwidXNlcklkIjo0fQ.De7NJ5vQ59mEJzRRv*7Ykfld6UvUdyN*-fd83H5ho3w'

Delete User Registration for an event

curl --location --request DELETE 'http://localhost:8080/events/6/register' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InN1cmFqQGdtYWlsLmNvbSIsImV4cCI6MTczMjU0MjIzNCwidXNlcklkIjo0fQ.De7NJ5vQ59mEJzRRv*7Ykfld6UvUdyN*-fd83H5ho3w'
