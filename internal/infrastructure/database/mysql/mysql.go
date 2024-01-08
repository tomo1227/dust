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

// CreateUser 関数を呼び出してユーザーを作成する
// ctx := context.Background()
// _, err = CreateUser(ctx, client)
// if err != nil {
// 	log.Fatalf("failed creating user: %v", err)
// }

// func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
// 	u, err := client.User.
// 		Create().
// 		SetAge(30).
// 		SetName("a8m").
// 		Save(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed creating user: %w", err)
// 	}
// 	log.Println("user was created: ", u)
// 	return u, nil
// }

// func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
// 	u, err := client.User.
// 		Query().
// 		Where(user.Name("a8m")).
// 		// `Only` fails if no user found,
// 		// or more than 1 user returned.
// 		Only(ctx)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed querying user: %w", err)
// 	}
// 	log.Println("user returned: ", u)
// 	return u, nil
// }
