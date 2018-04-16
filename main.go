package main

import (
	"Go-GoSAFE.converter/config"
	"Go-GoSAFE.converter/server"
)

func main() {
	config.CreateDBConnection()
	server.Start()
}
