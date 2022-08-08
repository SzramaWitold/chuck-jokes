package repositories

import "gorm.io/gorm"

//Paginate scope for paginate lists from GORM
func paginate[T interface{}](p *Pagination[T]) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Page - 1) * p.PerPage
		return db.Offset(offset).Limit(p.PerPage)
	}
}
