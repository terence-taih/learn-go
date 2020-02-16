package dal

import (
	"context"
	"fmt"
	"gitlab.myteksi.net/gophers/go/commons/data"
	"gitlab.myteksi.net/gophers/go/commons/util/redis/gredis"
	"gitlab.myteksi.net/gophers/go/commons/util/resilience/circuitbreaker"
)


type SuperHeroDao struct{
	data.Strategy
	dataRef data.Entity
}

func (dao *SuperHeroDao) Init() {
	cb := &circuitbreaker.Setting {
		TagStr: "superhero",
		Timeout: 1000,
	}
	dbConfig := &data.MysqlConfig{
		Dsn: "root:root@tcp(localhost:3306)/training?parseTime=true",
	}

	redisConfig := &gredis.Config{
		Host:    "localhost",
		Port:    6379,
		MaxConn: 20,
	}

	dbStrategy := data.NewMySQLHystrixStrategyWithCustomCB(dbConfig, cb)
	//dbStrategy := data.NewMySQLStrategy(dbConfig)
	redisStrategy := data.NewRedisStrategyV3(redisConfig, 60, data.DefaultRedisKeyGenerator{})
	dao.Strategy = data.NewMySQLRedisStrategyComposable(dbStrategy, redisStrategy)
	dao.dataRef = &Superhero{}
}


func (dao *SuperHeroDao) LoadByID(ctx context.Context, ID string) (*Superhero, error) {
	data, err := dao.Strategy.LoadByIDContext(ctx, dao.dataRef, ID)
	if err != nil {
		return nil, err
	}
	return data.(*Superhero), nil
}

func (dao *SuperHeroDao) Save(ctx context.Context, data *Superhero) error {
	return dao.Strategy.SaveContext(ctx, data)
}

func TestDal() {
	dao := SuperHeroDao{}
	dao.Init()
	if err := dao.Save(context.Background(), &Superhero{Name: "Thai Nguyen"}); err != nil {
		fmt.Println("Error when saving:", err)
	}

	superhero, e := dao.LoadByID(context.Background(), "1")

	if e != nil {
		fmt.Println("Error when read id 2", e)
	}
	fmt.Println(superhero.Name)

}