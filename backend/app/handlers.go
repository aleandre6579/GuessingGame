package app

import (
	"encoding/json"
	"fmt"
	"guessing-game/db/models/guess"
	user "guessing-game/db/models/user"
	"net/http"
	"strconv"

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

	jwt, err := GenerateJWT(u.ID, u.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(jwt))
	w.WriteHeader(http.StatusOK)
}

func Register(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := user.CreateUser(db, req.Username, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jwt, err := GenerateJWT(u.ID, u.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write([]byte(jwt))
	w.WriteHeader(http.StatusOK)
}

func Guess(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var req GuessRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var correctGuess *guess.Guess
	var err error
	fmt.Println(guess.GetGuessAtLevel(db, req.LevelName))
	if correctGuess, err = guess.GetGuessAtLevel(db, req.LevelName); err != nil {
		return
	}

	n, err := strconv.Atoi(req.Number)
	if err != nil {
		fmt.Println("Failed atoi")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if n != correctGuess.Number {
		fmt.Printf("Incorrect guess (%d != %d)", n, correctGuess.Number)
		w.WriteHeader(http.StatusOK)
		return
	}

	fmt.Printf("Correct guess (%d = %d)", n, correctGuess.Number)
	guess.RegenerateGuess(db, req.LevelName)
	w.WriteHeader(http.StatusCreated)
}
