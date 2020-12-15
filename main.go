package main

import (
	"fmt"
	"log"

	_ "goMarketer/models"
	_ "goMarketer/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	driverErr := orm.RegisterDriver("postgres", orm.DRPostgres)

	if driverErr != nil {
		log.Fatal(fmt.Sprintf("Failed to register the driver %v", driverErr))
	}

	databaseErr := orm.RegisterDataBase("default", "postgres", beego.AppConfig.String("db_url"))

	if databaseErr != nil {
		log.Fatal(fmt.Sprintf("Failed to connect to the database %v", databaseErr))
	}
}

func main() {
	err := orm.RunSyncdb("default", false, true)

	if err != nil {
		fmt.Println(err)
	}

	beego.Run()
}
