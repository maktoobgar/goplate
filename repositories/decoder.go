package repositories

import (
	"encoding/json"
	"service/pkg/copier"
	"time"
)

type group struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func (u Group) MarshalJSON() ([]byte, error) {
	return json.Marshal(copier.Copy(&group{}, &u))
}

type permission struct {
	ID           int32     `json:"id"`
	PermissionID int32     `json:"permission_id"`
	Name         string    `json:"name"`
	UserID       int32     `json:"user_id"`
	GroupID      int32     `json:"group_id"`
	CreatedAt    time.Time `json:"created_at"`
}

func (u Permission) MarshalJSON() ([]byte, error) {
	return json.Marshal(copier.Copy(&permission{}, &u))
}

type token struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (u Token) MarshalJSON() ([]byte, error) {
	return json.Marshal(copier.Copy(&token{}, &u))
}

type user struct {
	ID                   int32     `json:"id"`
	PhoneNumber          string    `json:"phone_number"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	Avatar               string    `json:"avatar"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	DisplayName          string    `json:"display_name"`
	Gender               int32     `json:"gender"`
	IsActive             bool      `json:"is_active"`
	Registered           bool      `json:"registered"`
	DeactivationReason   string    `json:"deactivation_reason"`
	IsAdmin              bool      `json:"is_admin"`
	OtpRemainingAttempts int32     `json:"otp_remaining_attempts"`
	OtpCode              int32     `json:"otp_code"`
	OtpDueDate           time.Time `json:"otp_due_date"`
	IsSuperuser          bool      `json:"is_superuser"`
	CreatedAt            time.Time `json:"created_at"`
}

func (u User) MarshalJSON() ([]byte, error) {
	return json.Marshal(copier.Copy(&user{}, &u))
}

type usersGroup struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	GroupID   int32     `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (u UsersGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(copier.Copy(&usersGroup{}, &u))
}
