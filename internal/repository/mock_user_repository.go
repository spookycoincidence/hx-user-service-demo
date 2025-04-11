package repository

import (
	"errors"
	"sync"

	"github.com/spookycoincidence/hx-user-service-demo/internal/model"
)

type MockUserRepository struct {
	data map[int]*model.User
	mu   sync.Mutex
	id   int
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		data: make(map[int]*model.User),
		id:   1,
	}
}

func (m *MockUserRepository) Create(user *model.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	user.ID = m.id
	m.data[m.id] = user
	m.id++
	return nil
}

func (m *MockUserRepository) GetByID(id int) (*model.User, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	user, ok := m.data[id]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}
