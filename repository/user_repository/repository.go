package user_repository

import (
	"portfolio/simple-Kanban/entity"
	"portfolio/simple-Kanban/pkg/errs"
)

type Repository interface {
	CreateUser(userPayload *entity.User) (*entity.User, errs.Error)
	UpdateUser(userPayload *entity.User) (*entity.User, errs.Error)
	FindOneUserByEmail(email string) (*entity.User, errs.Error)
}
