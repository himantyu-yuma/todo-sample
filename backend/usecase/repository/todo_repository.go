package repository

import (
	"context"
	"todo-sample/domain/entity"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *entity.Todo) error
	// FindByID(ctx context.Context, id string) (*entity.Todo, error)
	// FindAll(ctx context.Context) ([]*entity.Todo, error)
	// Update(ctx context.Context, todo *entity.Todo) error
	// Delete(ctx context.Context, id string) error
}
