package session

import (
	"sync"

	"github.com/jmm526/captrivia/internal/database"
)

type PlayerSession struct {
	Score int
}

type GameSession struct {
	UrlCode string
	Questions []database.Question
	PlayerSessions map[string]*PlayerSession
}

type SessionStore struct {
	sync.Mutex
	Sessions map[string]*GameSession
}