package config

import "os"

type Config struct {
	Addr             string
	DatabaseConfig   DatabaseConfig
	DifficultyConfig map[string]LevelBounds
}

type DatabaseConfig struct {
	User     string
	Password string
	Addr     string
	DBName   string
}

type LevelBounds struct {
	Lower int
	Upper int
}

func CreateConfig() Config {
	config := Config{
		Addr: ":8080",
		DatabaseConfig: DatabaseConfig{
			User:     "user",
			Password: "password",
			Addr:     "localhost:3306",
			DBName:   "db",
		},
		DifficultyConfig: map[string]LevelBounds{
			"easy":   LevelBounds{1, 5},
			"medium": LevelBounds{1, 10},
			"hard":   LevelBounds{1, 20},
		},
	}

	if dbAddress, exists := os.LookupEnv("DB_ADDR"); exists {
		config.DatabaseConfig.Addr = dbAddress
	}

	if dbUser, exists := os.LookupEnv("DB_USER"); exists {
		config.DatabaseConfig.User = dbUser
	}

	if dbPassword, exists := os.LookupEnv("DB_PASSWORD"); exists {
		config.DatabaseConfig.Password = dbPassword
	}

	if dbName, exists := os.LookupEnv("DB_NAME"); exists {
		config.DatabaseConfig.DBName = dbName
	}

	return config
}
