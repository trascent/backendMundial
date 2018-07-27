package main

import (
	_"github.com/trascent/backendMundial/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_"github.com/lib/pq"
	_"github.com/astaxie/beego/context"
  "github.com/astaxie/beego/plugins/cors"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
	//orm.RegisterDataBase("default", "postgres", "postgres://postgres:postgres@127.0.0.1/equipos?sslmode=disable")
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
