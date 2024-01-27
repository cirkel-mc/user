package repository

import (
	"cirkel/user/internal/domain/model"
	"context"
)

type Cache interface{}

type Psql interface {
	FindClientByClientId(ctx context.Context, clientId string) (*model.Client, error)
	GetAllGenre(ctx context.Context) ([]*model.Genre, error)
	CreateUserDevice(ctx context.Context, ud *model.UserDevice) error
	GetUserGenreByUser(ctx context.Context, userId int) ([]*model.UserGenre, error)
	GetUserGenreByGenre(ctx context.Context, genreId int) ([]*model.UserGenre, error)
	FindUserById(ctx context.Context, id int) (*model.User, error)
	FindUserByUsernameOrEmail(ctx context.Context, val string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) error
	DeleteUser(ctx context.Context, user *model.User) error
	GetInstrumentals(ctx context.Context) ([]*model.Instrument, error)
	GetSkills(ctx context.Context) ([]*model.Skill, error)
}
