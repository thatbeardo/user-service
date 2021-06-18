// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"database/sql"
	"time"
)

type User struct {
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
}

type UserType struct {
	ID   int32
	Disc string
}
