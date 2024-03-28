package dtos

import (
	"github.com/melardev/GoBeegoApiCrud/models"
)

type CreateTodo struct {
	Title       string `form:"title" json:"title" xml:"title" binding:"required"`
	Description string `form:"description" json:"description" xml:"description"`
	Completed   bool   `form:"completed" json:"completed" xml:"completed"`
}

func GetTodoDto(todo *models.Todo, includeDescription bool) map[string]interface{} {
	dto := map[string]interface{}{
		"id":    todo.Id,
		"title": todo.Title,
	}

	if includeDescription {
		dto["description"] = todo.Description
	}

	dto["completed"] = todo.Completed
	dto["created_at"] = todo.CreatedAt
	dto["updated_at"] = todo.UpdatedAt

	return dto
}

func GetTodoListDto(todos []models.Todo) []interface{} {

	result := make([]interface{}, len(todos))

	for i := 0; i < len(todos); i++ {
		result[i] = GetTodoDto(&todos[i], false)
	}

	return result
}
