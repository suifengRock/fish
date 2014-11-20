package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type User struct {
	Id   int64
	Name string `xorm:"varchar(25) not null unique 'usr_name'"`
}

func main() {

	addr := os.Getenv("DB_PORT_3306_TCP_ADDR")
	port := os.Getenv("DB_PORT_3306_TCP_PORT")
	proto := os.Getenv("DB_PORT_3306_TCP_PROTO")
	user := os.Getenv("DB_ENV_MYSQL_USER")
	password := os.Getenv("DB_ENV_MYSQL_PASSWORD")
	database := os.Getenv("DB_ENV_MYSQL_DATABASE")

	conn := "root:123456@/my?charset=utf8"

	if addr != "" {
		conn = fmt.Sprintf("%v:%v@%v(%v:%v)/%v?charset=utf8", user, password, proto, addr, port, database)
		fmt.Println("the connection is " + conn)
	}

	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		panic(err)
	}
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
