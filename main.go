package main

import (
	"fmt"

	_ "goMarketer/models"
	_ "goMarketer/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("db_url"))
}

func main() {
	err := orm.RunSyncdb("default", false, true)
	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
