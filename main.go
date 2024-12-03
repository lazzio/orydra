package main

import (
	"orydra/core/webserver"
)

func main() {
	r := webserver.Router()
	webserver.StartServer(r)
}
