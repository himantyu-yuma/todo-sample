package usecase

import (
	"context"
	"todo-sample/domain/entity"
	"todo-sample/usecase/repository"

	"github.com/google/uuid"
)

type TodoUsecase interface {
	CreateTodo(ctx context.Context, title, description string) (*entity.Todo, error)
	// GetTodo(ctx context.Context, id string) (*entity.Todo, error)
	// ListTodos(ctx context.Context) ([]*entity.Todo, error)
	// UpdateTodo(ctx context.Context, id, title, description string) (*entity.Todo, error)
	// DeleteTodo(ctx context.Context, id string) error
	// CompleteTodo(ctx context.Context, id string) (*entity.Todo, error)
}

type todoUsecase struct {
	todoRepo repository.TodoRepository
}

func NewTodoUsecase(todoRepo repository.TodoRepository) TodoUsecase {
	return &todoUsecase{
		todoRepo: todoRepo,
	}
}

func (uc *todoUsecase) CreateTodo(ctx context.Context, title, description string) (*entity.Todo, error) {
	todo := entity.NewTodo(uuid.NewString(), title, description)
	if err := uc.todoRepo.Create(ctx, todo); err != nil {
		return nil, err
	}
	return todo, nil
}
