package guess

import "gorm.io/gorm"

type Guess struct {
	gorm.Model

	Number int
}
