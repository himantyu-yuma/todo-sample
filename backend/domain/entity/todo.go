package entity

import (
	"time"
)

// TODOアイテムの構造体（エンティティ）
type Todo struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// 新しいTODOアイテムを作成するファクトリ関数
func NewTodo(id, title, description string) *Todo {
	return &Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

// // TODOを完了済みにマークする
// func (t *Todo) MarkAsComplete() {
// 	t.Completed = true
// 	t.UpdatedAt = time.Now()
// }

// // TODOのタイトルと説明を更新する
// func (t *Todo) Update(title, description string) {
// 	t.Title = title
// 	t.Description = description
// 	t.UpdatedAt = time.Now()
// }
