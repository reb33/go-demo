package auth

import (
	"adv_demo/internal/user"
)

type AuthService struct {
	UserRepo *user.UserRepository
}

func NewAuthService(userRepo *user.UserRepository) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
	}
}

func (s *AuthService) Register(email, password, name string) (string, error) {
	_, isFound, _ := s.UserRepo.FindByEmail(email)
	if isFound {
		return "", ErrUserAlreadyExist
	}

	user := &user.User{
		Email:    email,
		Password: "",
		Name:     name,
	}
	_, err := s.UserRepo.Create(user)
	if err != nil {
		return "", err
	}
	return user.Email, nil
}
