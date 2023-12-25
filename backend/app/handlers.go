package app

import (
	"fmt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login ", r.Body)
	w.WriteHeader(http.StatusOK)
}

func Register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register ", r.Body)
	w.WriteHeader(http.StatusOK)
}

func Guess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Guess ", r.Body)
	w.WriteHeader(http.StatusOK)
}
