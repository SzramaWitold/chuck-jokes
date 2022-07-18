package main

import (
	"chuck-jokes/api"
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
	server := api.StartEngine()
	server.Engine.Run(":8080")
}
