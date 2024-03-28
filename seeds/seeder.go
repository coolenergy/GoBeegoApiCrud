package seeds

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/icrowley/fake"
	"github.com/melardev/GoBeegoApiCrud/models"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {

	return rand.Intn(max-min) + min
}
func Seed() {
	o := orm.NewOrm()
	qt := o.QueryTable(new(models.Todo))
	insert, err := qt.PrepareInsert()
	if err != nil {
		panic(err.Error())
	}

	fake.Seed(time.Now().Unix())

	qb, err := orm.NewQueryBuilder("mysql")
	if err != nil {
		panic(err.Error())
	}
	tableName := new(models.Todo).TableName()
	qb.Select(fmt.Sprintf("%s.id", tableName)).From(tableName)
	rawSQL := qb.String()
	println(rawSQL)
	var todos []models.Todo
	countTodos, err := o.Raw(rawSQL).QueryRows(&todos)
	todosToSeed := 12
	todosToSeed -= int(countTodos)

	for i := 0; i < todosToSeed; i++ {
		completed := true
		if randomInt(0, 20)%2 == 0 {
			completed = false
		}
		id, err := insert.Insert(&models.Todo{
			Title:       fake.WordsN(randomInt(2, 4)),
			Description: fake.SentencesN(randomInt(1, 3)),
			Completed:   completed,
		})

		if err != nil {
			panic(err.Error())
		} else {
			println(fmt.Sprintf("Created Todo with id %d", id))
		}
	}
}
