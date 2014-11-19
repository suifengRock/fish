package main

import (
	"fmt"
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
	IdNum       string `xorm:"varchar(18) index unique not null"`
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
	return xorm.NewEngine("mysql", "test:1234@/myData?charset=utf8")
}

func sync(engine *xorm.Engine) error {
	return engine.Sync(&BaseInfo{}, &tools.ProvincesCode{})
}

func RangeInsertData(orm *xorm.Engine, provincesIds []string) {

	rangeNum := 10

	for i := 0; i < rangeNum; i++ {

		randID := tools.Random(provincesIds, 1)
		pId, _ := strconv.ParseInt(randID, 10, 64)
		obj := NewBaseInfo(pId)
		orm.Insert(obj)
		// _, err := orm.Insert(obj)
		// if err != nil {
		// fmt.Println(err)
		// }
	}

}

func GenerationBigData(orm *xorm.Engine, count int) {

	queue := make(chan int, count)
	provincesIds := tools.GetAllProvincesID(orm)
	success := 0
	for i := 0; i < count; i++ {
		go func(xrom *xorm.Engine, ids []string, x int) {
			RangeInsertData(xrom, ids)
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
		return
	}
	defer orm.Close()
	// orm.ShowSql = true
	orm.SetMaxIdleConns(1024)
	orm.SetMaxOpenConns(5120)
	err = sync(orm)
	if err != nil {
		fmt.Println(err)
		return
	}

	tools.CheckProvincesData(orm)

	begin := 50
	for i := 0; i < begin; i++ {

		GenerationBigData(orm, 1024)

	}

}
