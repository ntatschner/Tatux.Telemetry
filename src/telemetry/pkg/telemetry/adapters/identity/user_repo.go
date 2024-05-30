package identity

import (
	"github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/telemetry/domain"
	"github.com/google/uuid"
)

type UUIDGenerator struct {}

func (g *UUIDGenerator) NewID() string {
    return uuid.New().String()
}
var _ domain.IDGenerator = &UUIDGenerator{}

type UserRepository struct {
	users map[string]*domain.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[string]*domain.User)}
}

func (r *UserRepository) FindById(id string) (*domain.User, error) {
	user, ok := r.users[id]
	if !ok {
		return nil, &domain.NotFoundError{Entity: "User", Input: id}
	}
	return user, nil
}

func (r *UserRepository) Save(user *domain.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *UserRepository) FindByEmail(email string) (*domain.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, &domain.NotFoundError{Entity: "User", Input: email}
}

func (r *UserRepository) FetchAll() ([]domain.User, error) {
	users := make([]domain.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, *user)
	}
	return users, nil
}

func (r *UserRepository) FindByLastName(lastName string) ([]domain.User, error) {
	users := make([]domain.User, 0)
	for _, user := range r.users {
		if user.LastName == lastName {
			users = append(users, *user)
		}
	}
	return users, nil
}

func (r *UserRepository) FindByFirstName(firstName string) ([]domain.User, error) {
	users := make([]domain.User, 0)
	for _, user := range r.users {
		if user.FirstName == firstName {
			users = append(users, *user)
		}
	}
	return users, nil
}

var _ domain.UserRepository = &UserRepository{}

type GroupRepository struct {
	groups map[string]*domain.Group
}

func NewGroupRepository() *GroupRepository {
	return &GroupRepository{groups: make(map[string]*domain.Group)}
}

func (r *GroupRepository) FindById(id string) (*domain.Group, error) {
	group, ok := r.groups[id]
	if !ok {
		return nil, &domain.NotFoundError{Entity: "Group", Input: id}
	}
	return group, nil
}

func (r *GroupRepository) Save(group *domain.Group) error {
	r.groups[group.ID] = group
	return nil
}

func (r *GroupRepository) FetchAll() ([]domain.Group, error) {
	groups := make([]domain.Group, 0, len(r.groups))
	for _, group := range r.groups {
		groups = append(groups, *group)
	}
	return groups, nil
}

var _ domain.GroupRepository = &GroupRepository{}