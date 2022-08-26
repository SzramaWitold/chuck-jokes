package main

import (
	"chuck-jokes/cmd"

	"github.com/joho/godotenv"
)

// @title           Chuck jokes api doc
// @version         1.0
// @description     Chuck jokes manage api

// @host      localhost:8080
// @BasePath  /
func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
