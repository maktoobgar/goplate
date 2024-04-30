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
