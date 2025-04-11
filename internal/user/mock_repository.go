package user

import (
	"errors"
	"sync"
)

type MockUserRepository struct {
	users map[string]*User
	mu    sync.Mutex
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*User),
	}
}

func (m *MockUserRepository) Create(user *User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.users[user.Email]; exists {
		return errors.New("user already exists")
	}

	m.users[user.Email] = user
	return nil
}

func (m *MockUserRepository) GetByEmail(email string) (*User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, exists := m.users[email]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}
