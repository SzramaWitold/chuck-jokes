package models

// GetAllModels for gorm migrations
func GetAllModels() []interface{} {
	var models []interface{}
	models = append(models, &Joke{})
	models = append(models, &Category{})
	models = append(models, &User{})

	return models
}
