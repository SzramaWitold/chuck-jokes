package main

import (
	"chuck-jokes/cmd"

	"github.com/joho/godotenv"
)

// Plan to GO
// Get random joke from external api
// Add dockerize database
// Save jokes to database with current date
// Authorize user
// Many to Many users-jokes
// Many to many categorie
// Create seeders
//

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
