package helpers

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"time"
)

func GenerateMigrationID(name string) string {
	timestamp := time.Now().Format("20060102_150405")
	return fmt.Sprintf("%s_%s", timestamp, name)
}

func GetMigrations[T any](tableName string, tableType *T) []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: GenerateMigrationID(fmt.Sprintf("add_%s_table", tableName)),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(tableType)
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(tableName)
			},
		},
	}
}
