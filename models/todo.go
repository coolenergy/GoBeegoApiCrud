package models

import "time"

type Todo struct {
	Id          int
	Title       string    `json:"title"`
	Description string    `orm:"default(null)" json:"description"`
	Completed   bool      `orm:"default(false)" json:"completed"`
	CreatedAt   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt   time.Time `orm:"auto_now;type(datetime)"`
}

func (this *Todo) TableName() string {
	return "todos"
}
