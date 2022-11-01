package initialize

import (
	"context"
	"fmt"
	"log"
	"os"

	"blog/ent"
	"blog/global"
	"gopkg.in/yaml.v3"

	_ "github.com/go-sql-driver/mysql"
)

func MySQL() {

	data, err := os.ReadFile("./config/yml/mysql.yml")
	if err != nil {
		log.Fatal(err)
	}

	db := global.CONFIG.Mysql
	err = yaml.Unmarshal(data, &db)
	if err != nil {
		log.Fatal(err)
	}

	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=True", db.User, db.Pass, db.Address, db.Database)
	client, err := ent.Open("mysql", dataSource)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	global.DB = client

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
