package database

// BaseSeed seed database with fake data
func BaseSeed() {
	OpenConnection()
	CreateJokes(10)
	users := CreateUsers(5, 5)
	for _, user := range users {
		CreateCategories(user,3,4)
	}
}

// CreateJokes create and return specify amount of jokes
func CreateJokes(amount int) []*Joke {
	jokes := make([]*Joke, amount)
	for i := 0; i < amount; i++ {
		jokes[i] = CreateJoke(nil)
	}

	return jokes
}

// CreateUsers create with favourite jokes
func CreateUsers(amount, favAmount int) []*User {
	users := make([]*User, amount)
	for i := 0; i < amount; i++ {
		user := CreateUser(nil)
		for j := 0; j < favAmount; j++ {
			joke := CreateJoke(nil)
			user.Favourites = append(user.Favourites, *joke)
		}
		users[i] = user
		DB.Save(&user)
	}

	return users
}

// CreateCategories create categories with attached jokes
func CreateCategories(user *User, amount, jokesAmount int) {
	for i := 0; i < amount; i++ {
		category := CreateCategory(user, nil)
		for j := 0; j < jokesAmount; j++ {
			joke := CreateJoke(nil)
			category.Jokes = append(category.Jokes, *joke)
		}
		DB.Save(&category)
	}
}
