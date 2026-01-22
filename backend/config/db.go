package config

import (
	"sync"
)

// UserStore holds users in memory
type UserStore struct {
	Users map[string]*User
	Mu    sync.RWMutex
}

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

var Store *UserStore

// InitStore initializes the in-memory user store
func InitStore() {
	Store = &UserStore{
		Users: make(map[string]*User),
	}
}
