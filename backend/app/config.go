package app

import "os"

type Config struct {
	Addr           string
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	User     string
	Password string
	Addr     string
	DBName   string
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
	}

	if dbAddress, exists := os.LookupEnv("DB_ADDRESS"); exists {
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
