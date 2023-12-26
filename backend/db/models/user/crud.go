package user

import (
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, username string, password string) (*User, error) {
	usr := &User{
		Username: username,
		Password: password,
	}

	return usr, db.Create(usr).Error
}

// Retrieve the user by ID.
func GetUser(db *gorm.DB, ID uint) (*User, error) {
	usr := &User{}
	err := db.First(usr, ID).Error
	return usr, err
}

// Retrieve the user by username.
func GetFirstUserByUsername(db *gorm.DB, username string) (*User, error) {
	usr := &User{}
	err := db.Where(&User{
		Username: username,
	}).First(usr).Error
	return usr, err
}

// Retrieve the user by username and password
func GetUserByUsernameAndPassword(db *gorm.DB, username string, password string) (*User, error) {
	usr := &User{}
	err := db.Where(&User{
		Username: username,
		Password: password,
	}).First(usr).Error
	return usr, err
}

func UpdateUserPassword(db *gorm.DB, ID uint, password string) (*User, error) {
	var usr *User = nil

	err := db.Transaction(func(tx *gorm.DB) error {

		var err error
		if usr, err = GetUser(tx, ID); err != nil {
			return err
		}

		usr.Password = password

		if err := tx.Save(&usr).Error; err != nil {
			return err
		}

		return nil
	})

	return usr, err
}
