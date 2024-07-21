package model

import (
	"context"
	"time"
	"todo-sample/domain/entity"
	"todo-sample/usecase/repository"

	"gorm.io/gorm"
)

// gorm.DB のインスタンスを保持し、データベース操作を行う
type todoRepository struct {
	db *gorm.DB
}

// NewTodoRepository は TodoRepository の実装を返す
func NewTodoRepository(db *gorm.DB) repository.TodoRepository {
	return &todoRepository{
		db: db,
	}
}

// GORMのモデル定義
type TodoModel struct {
	ID          string `gorm:"primaryKey"`
	UserID      string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string
	Completed   bool
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}

// エンティティからモデルへの変換
// インフラストラクチャ層とドメイン層の分離を維持
func (m *TodoModel) ToEntity() (*entity.Todo, error) {
	return entity.NewTodo(m.ID, m.Title, m.Description), nil
}

// DB操作の実装
// contextを使うことで、タイムアウトやキャンセレーションの制御を可能にしている
func (r *todoRepository) Create(ctx context.Context, todo *entity.Todo) error {
	model := &TodoModel{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
	}

	result := r.db.WithContext(ctx).Create(model)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
