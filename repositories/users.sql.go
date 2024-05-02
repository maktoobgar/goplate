// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package repositories

import (
	"context"
	"database/sql"
	"time"
	"service/pkg/errors"
	"service/global"
	"service/i18n/i18n_interfaces"
)

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, phone_number, email, password, avatar, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at FROM users WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email sql.NullString) (User, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.FirstName,
		&i.LastName,
		&i.DisplayName,
		&i.Gender,
		&i.IsActive,
		&i.Registered,
		&i.DeactivationReason,
		&i.IsAdmin,
		&i.OtpRemainingAttempts,
		&i.OtpCode,
		&i.OtpDueDate,
		&i.IsSuperuser,
		&i.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, phone_number, email, password, avatar, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at FROM users WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.FirstName,
		&i.LastName,
		&i.DisplayName,
		&i.Gender,
		&i.IsActive,
		&i.Registered,
		&i.DeactivationReason,
		&i.IsAdmin,
		&i.OtpRemainingAttempts,
		&i.OtpCode,
		&i.OtpDueDate,
		&i.IsSuperuser,
		&i.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}

const getUserWithTokenId = `-- name: GetUserWithTokenId :one
SELECT u.id, u.phone_number, u.email, u.password, u.avatar, u.first_name, u.last_name, u.display_name, u.gender, u.is_active, u.registered, u.deactivation_reason, u.is_admin, u.otp_remaining_attempts, u.otp_code, u.otp_due_date, u.is_superuser, u.created_at FROM users u JOIN tokens t ON u.id = t.user_id WHERE t.id = $1
`

func (q *Queries) GetUserWithTokenId(ctx context.Context, id int32) (User, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, getUserWithTokenId, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.FirstName,
		&i.LastName,
		&i.DisplayName,
		&i.Gender,
		&i.IsActive,
		&i.Registered,
		&i.DeactivationReason,
		&i.IsAdmin,
		&i.OtpRemainingAttempts,
		&i.OtpCode,
		&i.OtpDueDate,
		&i.IsSuperuser,
		&i.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}

const loginUserWithEmail = `-- name: LoginUserWithEmail :one
SELECT id, phone_number, email, password, avatar, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at FROM users WHERE email = $1
`

func (q *Queries) LoginUserWithEmail(ctx context.Context, email sql.NullString) (User, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, loginUserWithEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.FirstName,
		&i.LastName,
		&i.DisplayName,
		&i.Gender,
		&i.IsActive,
		&i.Registered,
		&i.DeactivationReason,
		&i.IsAdmin,
		&i.OtpRemainingAttempts,
		&i.OtpCode,
		&i.OtpDueDate,
		&i.IsSuperuser,
		&i.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}

const loginUserWithPhoneNumber = `-- name: LoginUserWithPhoneNumber :one
SELECT id, phone_number, email, password, avatar, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at FROM users WHERE phone_number = $1
`

func (q *Queries) LoginUserWithPhoneNumber(ctx context.Context, phoneNumber string) (User, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, loginUserWithPhoneNumber, phoneNumber)
	var i User
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.FirstName,
		&i.LastName,
		&i.DisplayName,
		&i.Gender,
		&i.IsActive,
		&i.Registered,
		&i.DeactivationReason,
		&i.IsAdmin,
		&i.OtpRemainingAttempts,
		&i.OtpCode,
		&i.OtpDueDate,
		&i.IsSuperuser,
		&i.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}

const registerUser = `-- name: RegisterUser :one
INSERT INTO users (
  phone_number, email, display_name, password, created_at
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id, phone_number, email, password, avatar, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at
`

type RegisterUserParams struct {
	PhoneNumber string         `json:"phone_number"`
	Email       sql.NullString `json:"email"`
	DisplayName string         `json:"display_name"`
	Password    string         `json:"password"`
	CreatedAt   time.Time      `json:"created_at"`
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) (User, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, registerUser,
		arg.PhoneNumber,
		arg.Email,
		arg.DisplayName,
		arg.Password,
		arg.CreatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.FirstName,
		&i.LastName,
		&i.DisplayName,
		&i.Gender,
		&i.IsActive,
		&i.Registered,
		&i.DeactivationReason,
		&i.IsAdmin,
		&i.OtpRemainingAttempts,
		&i.OtpCode,
		&i.OtpDueDate,
		&i.IsSuperuser,
		&i.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}

const updateAvatar = `-- name: UpdateAvatar :one
UPDATE users SET avatar = $1 WHERE id = $2 RETURNING id, phone_number, email, password, avatar, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at
`

type UpdateAvatarParams struct {
	Avatar sql.NullString `json:"avatar"`
	ID     int32          `json:"id"`
}

func (q *Queries) UpdateAvatar(ctx context.Context, arg UpdateAvatarParams) (User, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, updateAvatar, arg.Avatar, arg.ID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.FirstName,
		&i.LastName,
		&i.DisplayName,
		&i.Gender,
		&i.IsActive,
		&i.Registered,
		&i.DeactivationReason,
		&i.IsAdmin,
		&i.OtpRemainingAttempts,
		&i.OtpCode,
		&i.OtpDueDate,
		&i.IsSuperuser,
		&i.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}

const updateMe = `-- name: UpdateMe :one
UPDATE users SET first_name = $1, last_name = $2, display_name = $3, gender = $4 WHERE id = $5 RETURNING id, phone_number, email, password, avatar, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at
`

type UpdateMeParams struct {
	FirstName   sql.NullString `json:"first_name"`
	LastName    sql.NullString `json:"last_name"`
	DisplayName string         `json:"display_name"`
	Gender      int32          `json:"gender"`
	ID          int32          `json:"id"`
}

func (q *Queries) UpdateMe(ctx context.Context, arg UpdateMeParams) (User, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, updateMe,
		arg.FirstName,
		arg.LastName,
		arg.DisplayName,
		arg.Gender,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.PhoneNumber,
		&i.Email,
		&i.Password,
		&i.Avatar,
		&i.FirstName,
		&i.LastName,
		&i.DisplayName,
		&i.Gender,
		&i.IsActive,
		&i.Registered,
		&i.DeactivationReason,
		&i.IsAdmin,
		&i.OtpRemainingAttempts,
		&i.OtpCode,
		&i.OtpDueDate,
		&i.IsSuperuser,
		&i.CreatedAt,
	)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}
