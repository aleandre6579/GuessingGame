package guess

import (
	"fmt"
	"guessing-game/config"
	"math/rand"

	"gorm.io/gorm"
)

func GetGuessAtLevel(db *gorm.DB, levelName string) (*Guess, error) {
	guess := &Guess{}
	err := db.Where("level = ?", levelName).First(guess).Error
	return guess, err
}

func RegenerateGuess(db *gorm.DB, levelName string) error {
	randomNumber, err := randomNumberAtLevel(levelName)
	if err != nil {
		return err
	}

	guess := &Guess{
		Level:  levelName,
		Number: randomNumber,
	}

	// Delete guess then create new one at that level
	if err := db.Where("level = ?", levelName).Delete(&Guess{}).Error; err != nil {
		fmt.Println("No guess yet to delete at level " + levelName)
		return err
	}
	if err := db.Create(guess).Error; err != nil {
		fmt.Println("Failed to create new guess")
		return err
	}

	return nil
}

func randomNumberAtLevel(levelName string) (int, error) {
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
