package db

import (
	"fmt"

	"gorm.io/gorm"
)

func InitTable(db *gorm.DB, table interface{}) error {
	fmt.Println("MIGRATE")
	if db == nil {
		return fmt.Errorf("Database is invalid.")
	}

	migrator := db.Migrator()

	if !migrator.HasTable(table) {
		err := db.AutoMigrate(table)
		if err != nil {
			return err
		}
		fmt.Println("MIGRATING")
	}

	fmt.Println("DID MIGRATE")
	return nil
}
