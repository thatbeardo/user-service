// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"database/sql"
)

type Consumer struct {
	FanfitUserID int32
}

type Creator struct {
	FanfitUserID      int32
	PaymentInfo       string
	LogoPicture       string
	BackgroundPicture string
}

type User struct {
	ID             int32
	UserTypeID     int32
	FirstName      string
	LastName       string
	Email          string
	CreatedDate    interface{}
	Username       sql.NullString
	PhoneNo        sql.NullInt32
	Gender         sql.NullString
	ProfilePicture sql.NullString
	Bio            sql.NullString
}

type UserType struct {
	ID   int32
	Disc string
}
