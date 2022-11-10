package migrations

import (
	"GoLaravel/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type Link struct {
		Info string `gorm:"type:varchar(255);not null;default:''"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Link{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropColumn(&Link{}, "Info")
	}

	migrate.Add("2022_11_10_160112_add_info_fields_to_link", up, down)
}
