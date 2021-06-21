// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package repository

import (
	"context"
	"database/sql"
	"time"
)

const createConsumer = `-- name: CreateConsumer :one
INSERT INTO consumers(
  fanfit_user_id,
  temp_field
) VALUES(
  $1,
  $2
)
RETURNING fanfit_user_id, temp_field
`

type CreateConsumerParams struct {
	FanfitUserID int32
	TempField    sql.NullString
}

func (q *Queries) CreateConsumer(ctx context.Context, arg CreateConsumerParams) (Consumer, error) {
	row := q.db.QueryRowContext(ctx, createConsumer, arg.FanfitUserID, arg.TempField)
	var i Consumer
	err := row.Scan(&i.FanfitUserID, &i.TempField)
	return i, err
}

const createCreator = `-- name: CreateCreator :one
INSERT INTO creators(
  fanfit_user_id,
  payment_info,
  logo_picture,
  background_picture
) VALUES(
  $1,
  $2,
  $3,
  $4
)
RETURNING fanfit_user_id, payment_info, logo_picture, background_picture
`

type CreateCreatorParams struct {
	FanfitUserID      int32
	PaymentInfo       string
	LogoPicture       string
	BackgroundPicture string
}

func (q *Queries) CreateCreator(ctx context.Context, arg CreateCreatorParams) (Creator, error) {
	row := q.db.QueryRowContext(ctx, createCreator,
		arg.FanfitUserID,
		arg.PaymentInfo,
		arg.LogoPicture,
		arg.BackgroundPicture,
	)
	var i Creator
	err := row.Scan(
		&i.FanfitUserID,
		&i.PaymentInfo,
		&i.LogoPicture,
		&i.BackgroundPicture,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users ( 
  first_name,
  last_name, 
  email,
  user_type_id,
  username,
  phone_no,
  gender,
  profile_picture,
  bio
) VALUES (
  $1, 
  $2, 
  $3,
  $4, 
  $5, 
  $6,
  $7, 
  $8, 
  $9
)
RETURNING id, user_type_id, first_name, last_name, email, created_date, username, phone_no, gender, profile_picture, bio
`

type CreateUserParams struct {
	FirstName      string
	LastName       string
	Email          string
	UserTypeID     int32
	Username       sql.NullString
	PhoneNo        sql.NullString
	Gender         sql.NullString
	ProfilePicture sql.NullString
	Bio            sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.UserTypeID,
		arg.Username,
		arg.PhoneNo,
		arg.Gender,
		arg.ProfilePicture,
		arg.Bio,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserTypeID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedDate,
		&i.Username,
		&i.PhoneNo,
		&i.Gender,
		&i.ProfilePicture,
		&i.Bio,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE email = $1
RETURNING id, user_type_id, first_name, last_name, email, created_date, username, phone_no, gender, profile_picture, bio
`

func (q *Queries) DeleteUser(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, deleteUser, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserTypeID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedDate,
		&i.Username,
		&i.PhoneNo,
		&i.Gender,
		&i.ProfilePicture,
		&i.Bio,
	)
	return i, err
}

const getClient = `-- name: GetClient :one
SELECT id, user_type_id, first_name, last_name, email, created_date, username, phone_no, gender, profile_picture, bio, fanfit_user_id, temp_field FROM users
INNER JOIN consumers on consumers.fanfit_user_id = users.id
WHERE users.email = $1
`

type GetClientRow struct {
	ID             int32
	UserTypeID     int32
	FirstName      string
	LastName       string
	Email          string
	CreatedDate    time.Time
	Username       sql.NullString
	PhoneNo        sql.NullString
	Gender         sql.NullString
	ProfilePicture sql.NullString
	Bio            sql.NullString
	FanfitUserID   int32
	TempField      sql.NullString
}

func (q *Queries) GetClient(ctx context.Context, email string) (GetClientRow, error) {
	row := q.db.QueryRowContext(ctx, getClient, email)
	var i GetClientRow
	err := row.Scan(
		&i.ID,
		&i.UserTypeID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedDate,
		&i.Username,
		&i.PhoneNo,
		&i.Gender,
		&i.ProfilePicture,
		&i.Bio,
		&i.FanfitUserID,
		&i.TempField,
	)
	return i, err
}

const getCreator = `-- name: GetCreator :one
SELECT id, user_type_id, first_name, last_name, email, created_date, username, phone_no, gender, profile_picture, bio, fanfit_user_id, payment_info, logo_picture, background_picture FROM users
INNER JOIN creators on creators.fanfit_user_id = users.id
WHERE users.email = $1
`

type GetCreatorRow struct {
	ID                int32
	UserTypeID        int32
	FirstName         string
	LastName          string
	Email             string
	CreatedDate       time.Time
	Username          sql.NullString
	PhoneNo           sql.NullString
	Gender            sql.NullString
	ProfilePicture    sql.NullString
	Bio               sql.NullString
	FanfitUserID      int32
	PaymentInfo       string
	LogoPicture       string
	BackgroundPicture string
}

func (q *Queries) GetCreator(ctx context.Context, email string) (GetCreatorRow, error) {
	row := q.db.QueryRowContext(ctx, getCreator, email)
	var i GetCreatorRow
	err := row.Scan(
		&i.ID,
		&i.UserTypeID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedDate,
		&i.Username,
		&i.PhoneNo,
		&i.Gender,
		&i.ProfilePicture,
		&i.Bio,
		&i.FanfitUserID,
		&i.PaymentInfo,
		&i.LogoPicture,
		&i.BackgroundPicture,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, user_type_id, first_name, last_name, email, created_date, username, phone_no, gender, profile_picture, bio FROM users 
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserTypeID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.CreatedDate,
		&i.Username,
		&i.PhoneNo,
		&i.Gender,
		&i.ProfilePicture,
		&i.Bio,
	)
	return i, err
}
