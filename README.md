# Chatterton
_A simple way to message._

## Chatterton API
To return all messages

### Ping
Confirm the server is running

```shell
curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/
```

### Messages

WIP - Return all messages
```shell
curl -H "Content-Type: application/json" -X GET http://localhost:8080/api/messages/ | json_pp
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
