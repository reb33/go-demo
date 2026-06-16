package auth

import (
	"adv_demo/internal/user"
	"adv_demo/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo *user.UserRepository
	JWT      *jwt.JWT
}

func NewAuthService(userRepo *user.UserRepository, jwt *jwt.JWT) *AuthService {
	return &AuthService{
		UserRepo: userRepo,
		JWT:      jwt,
	}
}

func (s *AuthService) Register(email, password, name string) (string, error) {
	_, isFound, _ := s.UserRepo.FindByEmail(email)
	if isFound {
		return "", ErrUserAlreadyExist
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := &user.User{
		Email:    email,
		Password: string(hashPassword),
		Name:     name,
	}
	_, err = s.UserRepo.Create(user)
	if err != nil {
		return "", err
	}
	signedString, err := s.JWT.Create(email)
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func (s *AuthService) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", ErrInvalidCredentials
	}
	userDB, _ := s.UserRepo.FirstByEmail(email)
	if userDB == nil {
		return "", ErrInvalidCredentials
	}
	err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	} 
	signedString, err := s.JWT.Create(email)
	if err != nil {
		return "", err
	}
	return signedString, nil
}
