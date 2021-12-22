package main

import (
	"fmt"
	"kyc/config"
	"kyc/database"
	"kyc/server"
)

func main() {
	fmt.Println("Hello world")
	config.Init()
	database.StartDatabase()
	s := server.NewServer()
	s.Run()
}
