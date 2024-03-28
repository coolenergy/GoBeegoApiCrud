package services

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/melardev/GoBeegoApiCrud/models"
)

func FetchTodos() ([]models.Todo, error) {
	o := orm.NewOrm()
	var todos []models.Todo
	// o.QueryTable(new(models.Todo)).All(&todos)

	qb, _ := orm.NewQueryBuilder("mysql")

	tableName := new(models.Todo).TableName()
	qb.Select(
		fmt.Sprintf("%s.id", tableName),
		fmt.Sprintf("%s.title", tableName),
		fmt.Sprintf("%s.completed", tableName),
		fmt.Sprintf("%s.created_at", tableName),
		fmt.Sprintf("%s.updated_at", tableName)).
		From(tableName).
		OrderBy("created_at").Desc()

	rawQuery := qb.String()
	count, err := o.Raw(rawQuery).QueryRows(&todos)
	println("fetched", count, "rows")
	return todos, err
}

func FetchPendingTodos() (todos []models.Todo, err error) {
	return FetchTodosByCompleted(false)
}

func FetchCompletedTodos() (todos []models.Todo, err error) {
	return FetchTodosByCompleted(true)
}

func FetchTodosByCompleted(completed bool) (todos []models.Todo, err error) {
	o := orm.NewOrm()
	qb, _ := orm.NewQueryBuilder("mysql")

	tableName := new(models.Todo).TableName()
	qb.Select(
		fmt.Sprintf("%s.id", tableName),
		fmt.Sprintf("%s.title", tableName),
		fmt.Sprintf("%s.completed", tableName),
		fmt.Sprintf("%s.created_at", tableName),
		fmt.Sprintf("%s.updated_at", tableName)).
		From(tableName).
		Where("completed = ?").
		OrderBy("created_at").Desc()

	rawQuery := qb.String()
	count, err := o.Raw(rawQuery, completed).QueryRows(&todos)
	println("fetched", count, "rows")
	return todos, err
}

func FetchById(id int) (todo *models.Todo, err error) {
	o := orm.NewOrm()
	todo = &models.Todo{Id: id}
	err = o.Read(todo)
	return
}

func CreateTodo(title, description string, completed bool) (todo *models.Todo, err error) {
	todo = &models.Todo{
		Title:       title,
		Description: description,
		Completed:   completed,
	}

	o := orm.NewOrm()
	qs := o.QueryTable(todo)
	inserter, err := qs.PrepareInsert()
	if err != nil {
		return
	}

	id, err := inserter.Insert(todo)

	// We assume the Todo has been saved as such in the Db, but it may not be
	// For example if a string truncation took place
	todo.Id = int(id)
	return todo, err
}

func UpdateTodo(id int, title, description string, completed bool) (todo *models.Todo, err error) {
	todo, err = FetchById(id)

	if err != nil {
		return
	}

	todo.Title = title

	// TODO: handle this in a better way, the user should be able to set description to empty string
	// The intention is to check against nil but in go there are no nil strings, so we can not know
	// if the user intended to update the description to empty string or just update the other fields other than description.
	if description != "" {
		todo.Description = description
	}

	todo.Completed = completed
	o := orm.NewOrm()

	count, err := o.Update(todo)
	println("updated", count, "rows")
	return
}

func DeleteTodo(todo *models.Todo) error {
	o := orm.NewOrm()
	count, err := o.Delete(todo)
	println("deleted", count, "rows")
	return err
}

func DeleteAllTodos() (err error) {
	o := orm.NewOrm()
	todo := models.Todo{}
	count, err := o.QueryTable(todo.TableName()).Filter("id__isnull", false).Delete()
	println("Deleted", count, "rows")
	return err
}
