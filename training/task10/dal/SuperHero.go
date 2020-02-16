package dal

import (
	"fmt"
	"gitlab.myteksi.net/gophers/go/commons/data"
	"strconv"
)

type Superhero struct {
	ID int64 `sql-col:"id" sql-key:"id" sql-insert:"false"`
	Name string `sql-col:"name"`
}

func (superHero *Superhero) SetID(ID string){
	idInt, err := strconv.ParseInt(ID, 10, 64)
	if err != nil {
		fmt.Errorf("Invalid value for ID", err)
		return
	}
	superHero.ID = idInt
}

func (superHero *Superhero) GetID() string {
	return strconv.FormatInt(superHero.ID, 10)
}

func (superHero *Superhero) NewEntity() data.Entity {
	return &Superhero{}
}

func (superHero *Superhero) GetTableName() string {
	return "superhero"
}


func (superHero *Superhero) GetNamespace() string {
	return "superhero"
}
