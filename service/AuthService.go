package service

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/yourname/pet_messenger/model"
	"github.com/yourname/pet_messenger/repository"
)

type AuthService struct {
	jwtSecret   string
	accessTTL   time.Duration
	refreshTTL  time.Duration
	refreshRepo repository.IRefreshTokenRepository
}

func NewAuthService(jwtSecret string, accessTTL, refreshTTL time.Duration, refreshRepo repository.IRefreshTokenRepository) *AuthService {
	return &AuthService{
		jwtSecret:   jwtSecret,
		accessTTL:   accessTTL,
		refreshTTL:  refreshTTL,
		refreshRepo: refreshRepo,
	}
}

func (s *AuthService) GenerateTokens(ctx context.Context, userID string) (accessToken string, refreshToken string, err error) {
	accessClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(s.accessTTL).Unix(),
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err = at.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", "", err
	}

	refreshClaims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(s.refreshTTL).Unix(),
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err = rt.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", "", err
	}

	tokenModel := &model.RefreshToken{
		ID:        uuid.New().String(),
		UserID:    userID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.refreshTTL),
		CreatedAt: time.Now(),
	}
	if err := s.refreshRepo.Create(ctx, tokenModel); err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

// Обновление токенов по refresh
func (s *AuthService) RefreshTokens(ctx context.Context, refreshToken string) (newAccess, newRefresh string, err error) {
	// Проверяем наличие refresh в базе
	rt, err := s.refreshRepo.GetByToken(ctx, refreshToken)
	if err != nil {
		return "", "", errors.New("refresh token invalid")
	}

	if rt.ExpiresAt.Before(time.Now()) {
		_ = s.refreshRepo.Delete(ctx, refreshToken)
		return "", "", errors.New("refresh token expired")
	}

	// Валидируем подпись
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return "", "", errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("invalid claims")
	}

	userID := claims["user_id"].(string)

	// Удаляем старый refresh
	_ = s.refreshRepo.Delete(ctx, refreshToken)

	// Генерируем новые токены
	return s.GenerateTokens(ctx, userID)
}

// Logout — удаление всех refresh токенов пользователя
func (s *AuthService) Logout(ctx context.Context, userID string) error {
	return s.refreshRepo.DeleteAllByUserID(ctx, userID)
}
