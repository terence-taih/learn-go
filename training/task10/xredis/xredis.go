package xredis

import (
	"context"
	"fmt"
	"gitlab.myteksi.net/gophers/go/staples/gredis3"
	"gitlab.myteksi.net/gophers/go/staples/gredis3/conman"
	"gitlab.myteksi.net/gophers/go/staples/gredis3/gredisapi"
	"time"
)

func TestRedis() {
	ctx := context.Background()

	cm := conman.NewSingleHost(ctx,
		conman.Host("localhost"),
		conman.Port(6379),
		conman.MaxActive(20),
		conman.IdleTimeout(2*time.Second),
		conman.ReadTimeout(2*time.Second),
		conman.WriteTimeout(2*time.Second),
	)

	client, err := gredis3.NewClient(ctx, cm)
	if err != nil {
		panic(err)
	}

	defer client.ShutDown(ctx)

	_, err = client.Do(ctx, "SET", "abc", "123")
	if err != nil {
		panic(err)
	}

	var value []byte
	value, err = gredisapi.Bytes(client.DoReadOnly(ctx, "GET", "abc"))
	if err != nil {
		panic(err)
	}

	fmt.Println("Loaded value:", string(value))
}