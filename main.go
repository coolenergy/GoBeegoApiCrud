package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/melardev/GoBeegoApiCrud/infrastructure"
	_ "github.com/melardev/GoBeegoApiCrud/routers"
	"github.com/melardev/GoBeegoApiCrud/seeds"
	"os"
)

func main() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
		os.Exit(0)
	}

	// if in develop mode
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	if err := infrastructure.Initialize(); err != nil {
		panic(err.Error())
	}

	seeds.Seed()

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		// AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Print SQL statements executed
	orm.Debug = true
	beego.Run()
}
