package repositories

import (
	"service/pkg/copier"
	"time"
)

func NewUser[T any](instance *T) User {
	user := copier.Copy(&User{}, instance)
	if !user.IsHashed() {
		user.HashPassword()
	}
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	return user
}

func NewToken(content string, refreshToken bool, userId int32, expiresAt time.Time) Token {
	token := Token{
		Token:          content,
		IsRefreshToken: refreshToken,
		UserID:         userId,
		ExpiresAt:      expiresAt,
		CreatedAt:      time.Now(),
	}
	return token
}
