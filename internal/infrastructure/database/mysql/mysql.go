package mysql

import (
	"context"
	"log"

	"github.com/tomo1227/dust/tool/ent"

	_ "github.com/go-sql-driver/mysql"
)

func start() {
	client, err := ent.Open("mysql", "dust:test@tcp(mysql)/dust?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
