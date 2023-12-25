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
	fmt.Println("Login")

	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := user.GetUserByUsernameAndPassword(db, u.Username, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register")

	var u user.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := user.CreateUser(db, u.Username, u.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Guess(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Guess")

	var userGuess guess.Guess
	if err := json.NewDecoder(r.Body).Decode(&userGuess); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var correctGuess *guess.Guess
	var err error
	if correctGuess, err = guess.GetGuess(db); err != nil {
		return
	}

	if userGuess.Number != correctGuess.Number {
		fmt.Printf("Incorrect guess (%d != %d)", userGuess.Number, correctGuess.Number)
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Printf("Correct guess (%d = %d)", userGuess.Number, correctGuess.Number)
	w.WriteHeader(http.StatusCreated)
}
