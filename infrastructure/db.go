package infrastructure

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/melardev/GoBeegoApiCrud/models"
	"os"
)

// Opening a database and save the reference to `Database` struct.
func Initialize() error {
	// This app uses Query Builder which is not supported with SQLite ...
	// So this is why I use mysql
	// Register driver which is just a mapping between driverName and Database Type
	// orm.RegisterDriver("sqlite3", orm.DRSqlite)
	err := orm.RegisterDriver("mysql", orm.DRMySQL)

	// If you save the variables in conf/app.conf you can read them through beego.AppConfig.String("DB_USER")
	// But in production apps you should really use .env(and exclude this file from Git) instead,
	// otherwise you will be leaking your secrets
	username := os.Getenv("DB_USER")
	password :=  os.Getenv("DB_PASSWORD")
	dbName :=  os.Getenv("DB_NAME")
	host :=  os.Getenv("DB_HOST")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, dbName)
	// Register database with alias default(self explanatory)
	// orm.RegisterDataBase("default", "sqlite3", "app.db")
	if err = orm.RegisterDataBase("default", "mysql", connectionString); err != nil {
		panic(err.Error())
	}

	// Database alias
	name := "default"

	// Drop table and re-create
	force := false

	// Print log
	verbose := true

	orm.RegisterModel(new(models.Todo))
	err = orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}

	return err
}
