// Code generated by sqlc. DO NOT EDIT.

package repository

import (
	"context"
)

type Querier interface {
	CreateClients(ctx context.Context, arg CreateClientsParams) (Client, error)
	DeleteUser(ctx context.Context, email string) (User, error)
	GetClients(ctx context.Context, fanfitUserID int32) (GetClientsRow, error)
}

var _ Querier = (*Queries)(nil)
