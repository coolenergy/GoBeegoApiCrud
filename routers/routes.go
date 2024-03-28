package routers

import (
	"github.com/astaxie/beego"
	"github.com/melardev/GoBeegoApiCrud/controllers"
)

func init() {

	// beego.Router("/api/todos", &controllers.TodosController{}, "get:GetAllTodos")
	namespace := beego.NewNamespace("/api",
		beego.NSNamespace("/todos",
			beego.NSRouter("/", &controllers.TodosController{}, "get:GetAllTodos"),
			beego.NSRouter("/completed", &controllers.TodosController{}, "get:GetAllCompletedTodos"),
			beego.NSRouter("/pending", &controllers.TodosController{}, "get:GetAllPendingTodos"),
			beego.NSRouter("/:id", &controllers.TodosController{}, "get:GetTodoById"),
			beego.NSRouter("/", &controllers.TodosController{}, "post:CreateTodo"),
			beego.NSRouter("/:id", &controllers.TodosController{}, "put:UpdateTodo"),
			beego.NSRouter("/:id", &controllers.TodosController{}, "delete:DeleteTodo"),
			beego.NSRouter("/", &controllers.TodosController{}, "delete:DeleteAllTodos"),
		),
	)

	beego.AddNamespace(namespace)
}
