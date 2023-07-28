package main

import (
	"context"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/pank-art/features_practice/features"
	"log"
)

func main() {
	//подключаемся к базе данных
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{"http://localhost:8529"},
	})
	if err != nil {
		log.Fatal(err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "artem"),
	})
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := client.Database(ctx, "third")
	if err != nil {
		log.Fatal(err)
	}

	// работа с библиотекой:

	balance, err := features.BalanceAddr(ctx, db, "btcAddress/1")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	count, err := features.CountOutClust(ctx, db, "dITWeUoEbaxbmiVXpM1TbmFlmXJP2ZEe4QR7RqAL7M8BcMrWwiq2jkgsVwBCW5Ot")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	// таким образом можно запустить любой метод
}
