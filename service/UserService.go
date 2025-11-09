package service

import (
	"context"
	"errors"
	"time"

	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo        repository.IUserRepository
	jwtSecret   string
	tokenExpiry time.Duration
}

// Конструктор
func NewUserService(repo repository.IUserRepository, jwtSecret string, expiry time.Duration) *UserService {
	if jwtSecret == "" {
		panic("jwtSecret cannot be empty")
	}
	return &UserService{
		repo:        repo,
		jwtSecret:   jwtSecret,
		tokenExpiry: expiry,
	}
}

// Регистрация пользователя
func (s *UserService) Register(ctx context.Context, username, email, password string) (*model.User, error) {
	existing, _ := s.repo.GetByEmail(ctx, email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Username:  username,
		Email:     email,
		Password:  string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

// Логин пользователя (JWT)
func (s *UserService) LoginUser(ctx context.Context, email, password string) (*model.User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil || user == nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

// Получение профиля пользователя
func (s *UserService) GetProfile(ctx context.Context, userID string) (*model.User, error) {
	return s.repo.GetByID(ctx, userID)
}

// Обновление профиля пользователя
func (s *UserService) UpdateProfile(ctx context.Context, user *model.User) error {
	user.UpdatedAt = time.Now()
	return s.repo.Update(ctx, user)
}
