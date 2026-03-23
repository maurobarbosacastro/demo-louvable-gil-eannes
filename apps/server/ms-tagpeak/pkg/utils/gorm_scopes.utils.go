package utils

import (
	"gorm.io/gorm"
)

// ActiveScope filters out records where deleted = true
func ActiveScope() func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("deleted = ?", false)
	}
}

func StateScope(state *[]string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if state != nil && len(*state) > 0 {
			return db.Where("state in (?)", *state)
		}
		return db
	}
}

func UserScope(uuid *string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if uuid != nil {
			return db.Where("user = ?", *uuid)
		}
		return db
	}
}
