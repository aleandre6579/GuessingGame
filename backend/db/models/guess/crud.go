package guess

import "gorm.io/gorm"

func GetGuess(db *gorm.DB) (*Guess, error) {
	guess := &Guess{}
	err := db.First(guess, 0).Error
	return guess, err
}

func ReplaceGuess(db *gorm.DB, newNumber int) error {
	guess := &Guess{
		Number: newNumber,
	}

	// Delete guess with ID = 0
	if err := db.Delete(&Guess{}, 0).Error; err != nil {
		return err
	}

	if err := db.Create(guess).Error; err != nil {
		return err
	}

	return nil
}
