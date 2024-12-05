package domain

import (
	"fmt"

	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {
	err := db.AutoMigrate(&Movie{})
	if err != nil {
		return fmt.Errorf("failed to migrate models: %w", err)
	}
	return nil
}
