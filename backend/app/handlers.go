package app

import (
	"encoding/json"
	"fmt"
	"guessing-game/db/models/guess"
	user "guessing-game/db/models/user"
	"net/http"

	"gorm.io/gorm"
)

func Login(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := user.GetUserByUsernameAndPassword(db, req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jwt, err := GenerateJWT(u.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("JWT " + jwt)
	w.Write([]byte(jwt))
	w.WriteHeader(http.StatusOK)
}

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := user.CreateUser(db, req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Guess(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Guess")

	var req GuessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var correctGuess *guess.Guess
	var err error
	if correctGuess, err = guess.GetGuess(db); err != nil {
		return
	}

	if req.Number != correctGuess.Number {
		fmt.Printf("Incorrect guess (%d != %d)", req.Number, correctGuess.Number)
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Printf("Correct guess (%d = %d)", req.Number, correctGuess.Number)
	guess.RegenerateGuess(db, req.LevelName)
	w.WriteHeader(http.StatusCreated)
}
