package main

import (
	_ "github.com/lib/pq"

	"github.com/staciewilhelm/chatterton-messenger-server/config"
)

func main() {
	config.StartServer()
}
