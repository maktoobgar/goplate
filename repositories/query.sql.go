// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package repositories

import (
	"context"
	"database/sql"
	"time"
	"service/pkg/errors"
	"service/global"
	"service/i18n/i18n_interfaces"
)

const createToken = `-- name: CreateToken :one
INSERT INTO tokens (
  user_id, created_at
) VALUES (
  $1, $2
)
RETURNING id, user_id, created_at
`

type CreateTokenParams struct {
	UserID    int32     `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func (q *Queries) CreateToken(ctx context.Context, arg CreateTokenParams) (Token, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, createToken, arg.UserID, arg.CreatedAt)
	var i Token
	err := row.Scan(&i.ID, &i.UserID, &i.CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}

const getToken = `-- name: GetToken :one
SELECT id, user_id, created_at FROM tokens WHERE id = $1
`

func (q *Queries) GetToken(ctx context.Context, id int32) (Token, error) {
	translator := ctx.Value(g.TranslatorKey).(i18n_interfaces.TranslatorI)
	row := q.db.QueryRowContext(ctx, getToken, id)
	var i Token
	err := row.Scan(&i.ID, &i.UserID, &i.CreatedAt)
	if err != nil && err != sql.ErrNoRows {
		panic(errors.New(errors.UnexpectedStatus, translator.StatusCodes().InternalServerError(), err.Error()))
	}
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, phone_number, email, password, profile, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at FROM users WHERE id = $1
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
		&i.Profile,
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
SELECT u.id, u.phone_number, u.email, u.password, u.profile, u.first_name, u.last_name, u.display_name, u.gender, u.is_active, u.registered, u.deactivation_reason, u.is_admin, u.otp_remaining_attempts, u.otp_code, u.otp_due_date, u.is_superuser, u.created_at FROM users u JOIN tokens t ON u.id = t.user_id WHERE t.id = $1
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
		&i.Profile,
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
SELECT id, phone_number, email, password, profile, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at FROM users WHERE email = $1
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
		&i.Profile,
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
SELECT id, phone_number, email, password, profile, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at FROM users WHERE phone_number = $1
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
		&i.Profile,
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
RETURNING id, phone_number, email, password, profile, first_name, last_name, display_name, gender, is_active, registered, deactivation_reason, is_admin, otp_remaining_attempts, otp_code, otp_due_date, is_superuser, created_at
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
		&i.Profile,
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
