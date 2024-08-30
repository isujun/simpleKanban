package category_repository

import (
	"portfolio/simple-Kanban/dto"
	"portfolio/simple-Kanban/entity"
	"portfolio/simple-Kanban/pkg/errs"
)

type Repository interface {
	CreateCategory(categoryPayload *entity.Category) (*entity.Category, errs.Error)
	GetCategoriesWithTasks() (*[]dto.GetCategoriesResponse, errs.Error)
	UpdateCategory(categoryPayload *entity.Category) (*entity.Category, errs.Error)
	DeleteCategory(id int) errs.Error
}
