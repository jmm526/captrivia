package session

import (
	"fmt"
	"math/rand"

	"github.com/jmm526/captrivia/internal/database"
)

func generateSessionID() string {
	randBytes := make([]byte, 16)
	rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}

func CreateEmptySessionStore() *SessionStore {
	return &SessionStore{Sessions: make(map[string]*GameSession)}
}

func (store *SessionStore) CreateSession() (string, string) {
	store.Lock()
	defer store.Unlock()

	uniqueSessionID := generateSessionID()
	uniquePlayerID := generateSessionID()
	var playerSessions map[string]*PlayerSession = make(map[string]*PlayerSession)
	playerSessions[uniquePlayerID] = &PlayerSession{Score: 0}

	store.Sessions[uniqueSessionID] = &GameSession{UrlCode: uniqueSessionID, Questions: make([]database.Question, 0), PlayerSessions: playerSessions}
	return uniqueSessionID, uniquePlayerID
}

func (store *SessionStore) GetSession(sessionID string) (*GameSession, bool) {
	store.Lock()
	defer store.Unlock()

	mySession, exists := store.Sessions[sessionID]
	return mySession, exists
}