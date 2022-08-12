# Chatterton
_A simple way to message._

## Chatterton API

### Ping
Confirm the server is running

```shell
curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/
```

Example success response
```shell
{"response":{"message":"Success","code":"200"}}
```

### Messages
Request recent messages from all senders. By default:
- only messages from the _last 30 days_ are returned
- a limit of 100 messages are returned

```shell
curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/messages | json_pp
```

A _limit_ can be passed as a query param, but will be disregarded if greater than 100.
```shell
curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/messages?limit=5 | json_pp
```

Example response
```shell
{
   "data" : [
      {
         "created_at" : "2022-08-11T14:29:53.924863-06:00",
         "id" : "0a235eb8-ee39-4d32-9158-e189c4c30791",
         "message_text" : "message",
         "message_type" : "I just got back from vacation.",
         "recipient_id" : "2",
         "sender_id" : "1"
      },
      {
         "created_at" : "2022-08-11T14:29:53.924863-06:00",
         "id" : "b279c2f8-2d39-4826-b86d-e02cca59f113",
         "message_text" : "message",
         "message_type" : "Hey Caro! How are you?",
         "recipient_id" : "2",
         "sender_id" : "1"
      }
   ],
   "response" : {
      "code" : "200",
      "message" : "Success"
   }
}
```

Example error response
```shell
{
   "response" : {
      "code" : "400",
      "message" : "Error returning messages from app: an error occurred"
   }
}
```

Request recent messages for a recipient from a specific sender by userID
- only messages from the _last 30 days_ are returned
- a limit of 100 messages are returned
- both sender and recipient are required

Messages _from sender_ **brewster** _for recipient_ **bklein**
```shell
curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/messages?sender_id=1&recipient_id=2 | json_pp
```


## Getting Started Locally
### Install
Chatterton runs using [Go](https://go.dev/) v1.19 and [PostgreSQL](https://www.postgresql.org/) v12.

## Setup
Ensure both `GOPATH` and `GOROOT` are set correctly. To install dependencies, either user `make deps` or install each
package from `go.mod` individually.

Rename `sample.env` to `.env` and set the environment variables for your PostgreSQL database:
```
DB_DATABASE=
DB_HOST=
DB_PASSWORD=
DB_PORT=
DB_USERNAME=
```

## Start the Server
Once installation and setup is complete, `make start` will start the server.

### Running Tests Locally
Chatterton testing suite uses both [Ginkgo](https://onsi.github.io/ginkgo) and [Gomega](https://onsi.github.io/gomega/).

To run the test suite, `make deps`, then `make run-tests`.

Note: in the case of errors, follow [Ginkgo's installation steps](https://onsi.github.io/ginkgo/#installing-ginkgo).


### Running migrations
Migrations use [golang-migrate](https://github.com/golang-migrate/migrate). Please reference [the documentation](https://github.com/golang-migrate/migrate/blob/03613f14ac4f975eb0070a23958123c5d84e6b87/database/postgres/TUTORIAL.md#create-migrations)
for information on how to generate new migration files.
