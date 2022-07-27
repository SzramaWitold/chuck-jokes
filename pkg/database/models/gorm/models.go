package gorm

// GetAllModels for gorm migrations
func GetAllModels() []interface{} {
	models := []interface{}{}
	models = append(models, &Joke{})
	models = append(models, &Category{})
	models = append(models, &User{})

	return models
}
