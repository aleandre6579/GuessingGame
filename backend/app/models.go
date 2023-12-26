package app

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type GuessRequest struct {
	Number    string `json:"number"`
	LevelName string `json:"level"`
}

type Level struct {
	LevelName string
	Upper     int
	Lower     int
}
