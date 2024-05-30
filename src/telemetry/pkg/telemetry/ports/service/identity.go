package service

import (
    "github.com/ntatschner/Tatux.Telemetry/src/telemetry/pkg/telemetry/domain"
)

type UserService struct {
    generator domain.IDGenerator
    repo      domain.UserRepository
}

func NewUserService(generator domain.IDGenerator, repo domain.UserRepository) *UserService {
    return &UserService{generator: generator, repo: repo}
}

func (s *UserService) CreateUser(firstName string, lastName string, email string) (*domain.User, error) {
    user := domain.NewUser(s.generator, firstName, lastName, email)
    err := s.repo.Save(user)
    if err != nil {
        return nil, err
    }
    return user, nil
}