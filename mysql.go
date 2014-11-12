package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id   int64
	Name string `xorm:"varchar(25) not null unique 'usr_name'"`
}

func main() {
	engine, _ := xorm.NewEngine("mysql", "root:@/my?charset=utf8")
	defer engine.Close()

	// user := new(User)
	// user.Name = "my020na33e121"
	// affected, err := engine.Insert(user)
	// if err != nil {
	// 	fmt.Println("has error", err.Error())
	// }

	pEveryOne := make([]*User, 0)
	engine.Cols("Id", "usr_name").Find(&pEveryOne)

	fmt.Println("...", len(pEveryOne))
	for _, userOne := range pEveryOne {
		fmt.Println(userOne.Name, " dfdfd ", userOne.Id)
	}
}
