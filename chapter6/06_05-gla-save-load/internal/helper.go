package helper

import (
	"encoding/json"
	"fmt"
	"gla/pkg/gai"
	"os"
	"time"
)

// Session struct contains the session ID, timestamp, model, token count, messages, and isSelected flag
// This is used to store the session history
type Session struct {
	Id         string               `json:"Id"`
	Timestamp  string               `json:"timestamp"`
	Model      string               `json:"model"`
	TokenCount int                  `json:"tokencount"`
	Messages   []gai.MessageContent `json:"messages"`
	IsSelected bool                 `json:"isselected"`
}

// sessionStore contains the path to the session store and a slice of Session entries
type SessionStore struct {
	Path     string
	FileName string
	Sessions []Session
}

// GetStore returns the path to the session store
func (s *SessionStore) getStore() (string, error) {
	if s.Path == "" || s.FileName == "" {
		return "", fmt.Errorf("session store is not set")
	}

	return fmt.Sprintf("%s/%s", s.Path, s.FileName), nil
}

// Exits checks if the session store exists
func (s *SessionStore) exists() bool {
	path, err := s.getStore()

	if err != nil {
		return false
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

// Initialize the session store
func InitializeSessionStore(path, filename string) (*SessionStore, error) {
	var err error
	ss := &SessionStore{
		Path:     path,
		FileName: filename,
	}

	if ss.exists() {
		ss.Sessions, err = ss.GetSessionEntries()
		if err != nil {
			return nil, err
		}
		return ss, nil
	}

	_, err = ss.createStore()
	if err != nil {
		return nil, err
	}

	return ss, nil
}

// CreateStore creates a new session store
func (s *SessionStore) createStore() (string, error) {
	path, _ := s.getStore()

	if !s.exists() {
		err := os.MkdirAll(s.Path, 0755)
		if err != nil {
			return "", err
		}
	}

	se := []Session{}
	sl, err := json.Marshal(se)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(path, sl, 0644)
	if err != nil {
		return "", err
	}

	return path, nil
}

// GetSessionEntries returns all the session entries
func (s *SessionStore) GetSessionEntries() (sessionEntries []Session, err error) {
	path, err := s.getStore()
	if err != nil {
		return nil, err
	}

	se, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(se, &sessionEntries)
	if err != nil {
		return nil, err
	}

	return sessionEntries, nil
}

// GetSessionEntry returns a single session entry
func (s *SessionStore) GetSessionEntry(sessionId string) (Session, error) {
	se, err := s.GetSessionEntries()
	if err != nil {
		return Session{}, err
	}

	for _, s := range se {
		if s.Id == sessionId {
			return s, nil
		}
	}

	return Session{}, nil
}

func (s *SessionStore) SaveSessionEntries(sessionEntries []Session) error {
	var path string
	path, err := s.getStore()
	if err != nil {
		return err
	}

	if !s.exists() {
		path, err = s.createStore()
		if err != nil {
			return err
		}
	}

	sObject, err := json.Marshal(sessionEntries)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, sObject, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (s *SessionStore) AddSession(session Session) ([]Session, error) {
	sessionEntries, err := s.GetSessionEntries()
	if err != nil {
		return nil, err
	}

	// always add the new session at the beginning of the sessions slice
	sessionEntries = append([]Session{session}, sessionEntries...)
	s.Sessions = sessionEntries

	err = s.SaveSessionEntries(sessionEntries)
	if err != nil {
		return nil, err
	}

	return sessionEntries, nil
}

func (s *SessionStore) AddMessage(sessionId string, message gai.MessageContent, tokenCount int) (Session, error) {
	se, err := s.GetSessionEntries()
	if err != nil {
		return Session{}, err
	}

	for i, s := range se {
		if s.Id == sessionId {
			se[i].Messages = append(s.Messages, message)
			se[i].TokenCount = tokenCount
			se[i].Timestamp = time.Now().Format("2006-01-02 15:04")
			break
		}
	}

	err = s.SaveSessionEntries(se)
	if err != nil {
		return Session{}, err
	}

	return s.GetSessionEntry(sessionId)
}

func (s *SessionStore) SelectSession(sessionId string) (Session, error) {
	se, err := s.GetSessionEntries()
	if err != nil {
		return Session{}, err
	}

	if len(se) == 0 {
		return Session{}, fmt.Errorf("no session entries found")
	}

	for i, s := range se {
		if s.Id == sessionId {
			se[i].IsSelected = true
		} else {
			se[i].IsSelected = false
		}
	}

	err = s.SaveSessionEntries(se)
	if err != nil {
		return Session{}, err
	}

	return s.GetSessionEntry(sessionId)
}

func (s *SessionStore) DeleteSession(sessionId string) error {
	se, err := s.GetSessionEntries()
	if err != nil {
		return err
	}

	entry, err := s.GetSessionEntry(sessionId)
	if err != nil {
		return err
	}

	if entry.Id == "" {
		return fmt.Errorf("session entry not found")
	}

	if entry.IsSelected {
		// set index 0 as selected
		se[0].IsSelected = true
	}

	var newSessions []Session
	for _, session := range se {
		if session.Id != sessionId {
			newSessions = append(newSessions, session)
		}
	}

	err = s.SaveSessionEntries(newSessions)
	if err != nil {
		return err
	}

	return nil
}
