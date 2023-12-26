package guess

import "gorm.io/gorm"

type Guess struct {
	gorm.Model

	Level  string
	Number int
}
