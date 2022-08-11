package main

import (
	_ "github.com/lib/pq"

	"chatterton-messenger-server/config"
)

func main() {
	config.StartServer()
}
