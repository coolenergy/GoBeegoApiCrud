package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/melardev/GoBeegoApiCrud/dtos"
	"github.com/melardev/GoBeegoApiCrud/models"
	"github.com/melardev/GoBeegoApiCrud/services"
	"net/http"
	"strconv"
)

type TodosController struct {
	beego.Controller
}

func (this *TodosController) GetAllTodos() {
	todos, err := services.FetchTodos()
	if err != nil {
		this.Data["json"] = dtos.CreateErrorDtoWithMessage(err.Error())
		this.ServeJSON()
	}
	this.Data["json"] = dtos.GetTodoListDto(todos)
	this.ServeJSON()
}

func (this *TodosController) GetAllPendingTodos() {
	todos, err := services.FetchPendingTodos()
	if err != nil {
		this.Data["json"] = dtos.CreateErrorDtoWithMessage(err.Error())
		this.ServeJSON()
	}
	this.Data["json"] = dtos.GetTodoListDto(todos)
	this.ServeJSON()
}
func (this *TodosController) GetAllCompletedTodos() {
	todos, err := services.FetchCompletedTodos()
	if err != nil {
		this.Data["json"] = dtos.CreateErrorDtoWithMessage(err.Error())
		this.ServeJSON()
	}
	this.Data["json"] = dtos.GetTodoListDto(todos)
	this.ServeJSON()
}

func (this *TodosController) GetTodoById() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	todo, err := services.FetchById(id)
	if err != nil {
		return
	}

	this.Data["json"] = dtos.GetTodoDetaislDto(todo)
	this.ServeJSON()
}

func (this *TodosController) CreateTodo() {
	todo := &models.Todo{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, todo); err != nil {
		// this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		this.Data["json"] = dtos.CreateErrorDtoWithMessage(err.Error())
		this.ServeJSON()
		return
	}

	todo, err := services.CreateTodo(todo.Title, todo.Description, todo.Completed)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		this.Data["json"] = dtos.CreateErrorDtoWithMessage(err.Error())
		this.ServeJSON()
		return
	}

	this.Data["json"] = dtos.GetTodoDetaislDto(todo)
	this.ServeJSON()
}

func (this *TodosController) UpdateTodo() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		this.Data["json"] = dtos.CreateErrorDtoWithMessage("You must set an ID")
		this.ServeJSON()
		return
	}

	var todoInput models.Todo
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &todoInput); err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		this.Data["json"] = dtos.CreateErrorDtoWithMessage(err.Error())
		this.ServeJSON()
		return
	}

	todo, err := services.UpdateTodo(id, todoInput.Title, todoInput.Description, todoInput.Completed)
	if err != nil {
		this.Data["json"] = dtos.CreateErrorDtoWithMessage(err.Error())
		this.ServeJSON()
		return
	}

	this.Data["json"] = dtos.GetTodoDetaislDto(todo)
	this.ServeJSON()
}

func (this *TodosController) DeleteTodo() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		this.Data["json"] = dtos.CreateErrorDtoWithMessage("You must set an ID")
		this.ServeJSON()
		return
	}

	todo, err := services.FetchById(id)

	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		this.Data["json"] = dtos.CreateErrorDtoWithMessage("todo not found")
		this.ServeJSON()

		return
	}

	err = services.DeleteTodo(todo)

	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		this.Data["json"] = dtos.CreateErrorDtoWithMessage("Could not delete Todo")
		this.ServeJSON()
		return
	}

	this.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
	this.ServeJSON()
}

func (this *TodosController) DeleteAllTodos() {
	services.DeleteAllTodos()
	this.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
	this.ServeJSON()
}
