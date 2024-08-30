package task_repository

import (
	"portfolio/simple-Kanban/dto"
	"portfolio/simple-Kanban/entity"
	"portfolio/simple-Kanban/pkg/errs"
)

type Repository interface {
	CreateTask(taskPayload *entity.Task) (*entity.Task, errs.Error)
	GetTasks() (*[]dto.GetTasksResponse, errs.Error)
	UpdateTask(taskPayload *entity.Task) (*entity.Task, errs.Error)
	UpdateStatus(taskPayload *entity.Task) (*entity.Task, errs.Error)
	UpdateTaskCategory(taskPayload *entity.Task) (*entity.Task, errs.Error)
	DeleteTask(id int) errs.Error
}
