// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package repositories

import (
	"database/sql"
	"time"
)

type Group struct {
	ID        int32     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type Permission struct {
	ID           int32         `json:"id"`
	PermissionID int32         `json:"permission_id"`
	Name         string        `json:"name"`
	UserID       sql.NullInt32 `json:"user_id"`
	GroupID      sql.NullInt32 `json:"group_id"`
	CreatedAt    time.Time     `json:"created_at"`
}

type Token struct {
	ID             int32     `json:"id"`
	Token          string    `json:"token"`
	IsRefreshToken bool      `json:"is_refresh_token"`
	UserID         int32     `json:"user_id"`
	ExpiresAt      time.Time `json:"expires_at"`
	CreatedAt      time.Time `json:"created_at"`
}

type User struct {
	ID                   int32          `json:"id"`
	PhoneNumber          string         `json:"phone_number"`
	Email                sql.NullString `json:"email"`
	Password             string         `json:"password"`
	Profile              sql.NullString `json:"profile"`
	FirstName            sql.NullString `json:"first_name"`
	LastName             sql.NullString `json:"last_name"`
	DisplayName          string         `json:"display_name"`
	Gender               int32          `json:"gender"`
	IsActive             bool           `json:"is_active"`
	Registered           bool           `json:"registered"`
	DeactivationReason   sql.NullString `json:"deactivation_reason"`
	IsAdmin              bool           `json:"is_admin"`
	OtpRemainingAttempts int32          `json:"otp_remaining_attempts"`
	OtpCode              sql.NullInt32  `json:"otp_code"`
	OtpDueDate           sql.NullTime   `json:"otp_due_date"`
	IsSuperuser          bool           `json:"is_superuser"`
	CreatedAt            time.Time      `json:"created_at"`
}

type UsersGroup struct {
	ID        int32     `json:"id"`
	UserID    int32     `json:"user_id"`
	GroupID   int32     `json:"group_id"`
	CreatedAt time.Time `json:"created_at"`
}
