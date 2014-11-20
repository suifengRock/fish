package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
	"tools"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func init() {
	runtime.GOMAXPROCS(3)
}

type BaseInfo struct {
	Id          int64
	Name        string `xorm:"index"`
	Age         string
	Gender      string
	IdNum       string `xorm:"varchar(18) index not null"`
	ProvincesID int64
	CreatedAt   time.Time `xorm:"created"`
}

func NewBaseInfo(pId int64) (obj *BaseInfo) {
	obj = new(BaseInfo)
	obj.Name = tools.NoSpeStr(13)
	obj.Age = tools.NumberOnly(2)
	obj.Gender = tools.StringRand("男女", 1)
	obj.ProvincesID = pId
	obj.IdNum = tools.StringRand(tools.Capital+tools.Number+tools.Lowercase+tools.SpeStr, 18)
	return
}

func mysqlEngine() (*xorm.Engine, error) {
	addr := os.Getenv("DB_PORT_3306_TCP_ADDR")
	port := os.Getenv("DB_PORT_3306_TCP_PORT")
	proto := os.Getenv("DB_PORT_3306_TCP_PROTO")
	user := os.Getenv("DB_ENV_MYSQL_USER")
	password := os.Getenv("DB_ENV_MYSQL_PASSWORD")
	database := os.Getenv("DB_ENV_MYSQL_DATABASE")

	conn := "test:1234@/myData?charset=utf8"

	if addr != "" {
		conn = fmt.Sprintf("%v:%v@%v(%v:%v)/%v?charset=utf8", user, password, proto, addr, port, database)
		fmt.Println("the connection is " + conn)
	}

	return xorm.NewEngine("mysql", conn)
}

func sync(engine *xorm.Engine) error {
	return engine.Sync(&BaseInfo{}, &tools.ProvincesCode{})
}

func RangeInsertData(orm *xorm.Engine, provincesIds []string) (err error) {

	rangeNum := 10

	for i := 0; i < rangeNum; i++ {

		randID := tools.Random(provincesIds, 1)
		pId, _ := strconv.ParseInt(randID, 10, 64)
		obj := NewBaseInfo(pId)
		// orm.Insert(obj)
		_, err = orm.Insert(obj)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	return

}

func GenerationBigData(orm *xorm.Engine, count int) {

	queue := make(chan int, count)
	provincesIds := tools.GetAllProvincesID(orm)
	success := 0
	for i := 0; i < count; i++ {
		go func(xrom *xorm.Engine, ids []string, x int) {
			err := RangeInsertData(xrom, ids)
			if err != nil {
				return
			}
			success++
			queue <- x
		}(orm, provincesIds, i)
	}

	to := time.NewTimer(time.Second)
	for i := 0; i < count; i++ {
		to.Reset(time.Second * 3)
		select {
		case <-queue:
		case <-to.C:
			fmt.Println("......timeout")
		}
	}
	fmt.Println("success:", success)
}

func main() {

	orm, err := mysqlEngine()
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	defer orm.Close()
	// orm.ShowSql = true
	orm.SetMaxIdleConns(1024)
	orm.SetMaxOpenConns(5120)
	err = orm.DropTables(&BaseInfo{})
	if err != nil {
		fmt.Println(err.Error())
	}
	err = sync(orm)
	if err != nil {
		panic(err)
		fmt.Println(err)
		return
	}
	tools.CheckProvincesData(orm)

	begin := 10
	for i := 0; i < begin; i++ {

		GenerationBigData(orm, 1024)

	}

	count, _ := orm.Count(&BaseInfo{})
	fmt.Println("this is count:", count)

}
