package migrations

import (
	"github.com/Leonardo-lucas-pereira/tcc-api/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(models.User{})
	//db.AutoMigrate(models.User{})
}
