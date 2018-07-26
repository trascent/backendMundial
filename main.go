package main

import (
	_"github.com/trascent/practica1backend/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/lib/pq"
	_"github.com/astaxie/beego/context"
  "github.com/astaxie/beego/plugins/cors"
)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://postgres:123@127.0.0.1/equipos?sslmode=disable")
	beego.Debug("Filters init...")
    beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
        AllowAllOrigins: true,
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"X-ACCESS_TOKEN", "Access-Control-Allow-Origin", "Authorization", "Origin", "x-requested-with", "Content-Type", "Content-Range", "Content-Disposition", "Content-Description","Access-Control-Allow-Headers"},
        ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
        AllowCredentials: true,
    }))
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
