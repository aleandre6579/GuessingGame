package guess

import (
	"fmt"
	"guessing-game/config"
	"math/rand"

	"gorm.io/gorm"
)

func GetGuess(db *gorm.DB) (*Guess, error) {
	guess := &Guess{}
	err := db.First(guess, 0).Error
	return guess, err
}

func RegenerateGuess(db *gorm.DB, levelName string) error {
	randomNumber, err := randomNumber(levelName)
	if err != nil {
		return err
	}

	guess := &Guess{
		Number: randomNumber,
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

func randomNumber(levelName string) (int, error) {
	bounds, err := getLevel(levelName)
	if err != nil {
		return -1, err
	}

	number := rand.Intn(bounds.Upper+1-bounds.Lower) + bounds.Lower
	return number, nil
}

func getLevel(levelName string) (config.LevelBounds, error) {
	levels := config.CreateConfig().DifficultyConfig

	for name, bounds := range levels {
		if name == levelName {
			return bounds, nil
		}
	}
	return config.LevelBounds{}, fmt.Errorf("Level not found")
}
