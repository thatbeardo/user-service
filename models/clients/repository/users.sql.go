// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package repository

import (
	"context"
	"database/sql"
	"time"
)

const createClient = `-- name: CreateClient :one
INSERT INTO clients(
  fanfit_user_id,
  temp_field
) VALUES(
  $1,
  $2
)
RETURNING fanfit_user_id, temp_field
`

type CreateClientParams struct {
	FanfitUserID int32
	TempField    sql.NullString
}

func (q *Queries) CreateClient(ctx context.Context, arg CreateClientParams) (Client, error) {
	row := q.db.QueryRowContext(ctx, createClient, arg.FanfitUserID, arg.TempField)
	var i Client
	err := row.Scan(&i.FanfitUserID, &i.TempField)
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

const getClientByEmail = `-- name: GetClientByEmail :one
SELECT id, user_type_id, first_name, last_name, email, created_date, username, phone_no, gender, profile_picture, bio, fanfit_user_id, temp_field FROM users INNER JOIN clients
ON users.fanfit_user_id = clients.fanfit_user_id
WHERE email = $1
`

type GetClientByEmailRow struct {
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

func (q *Queries) GetClientByEmail(ctx context.Context, email string) (GetClientByEmailRow, error) {
	row := q.db.QueryRowContext(ctx, getClientByEmail, email)
	var i GetClientByEmailRow
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

const getClientByID = `-- name: GetClientByID :one
SELECT id, user_type_id, first_name, last_name, email, created_date, username, phone_no, gender, profile_picture, bio, fanfit_user_id, temp_field FROM users INNER JOIN clients
ON users.fanfit_user_id = clients.fanfit_user_id
WHERE fanfit_user_id = $1
`

type GetClientByIDRow struct {
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

func (q *Queries) GetClientByID(ctx context.Context, fanfitUserID int32) (GetClientByIDRow, error) {
	row := q.db.QueryRowContext(ctx, getClientByID, fanfitUserID)
	var i GetClientByIDRow
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
