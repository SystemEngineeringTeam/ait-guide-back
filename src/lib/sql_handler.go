package lib

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/go-sql-driver/mysql"
	// "github.com/kajiLabTeam/dx-waiting-time/conf"
)

func SqlConnect() (database *sql.DB) {
	var err error
	var db *sql.DB
	// c := conf.GetMysqlConfig()
	// dsn := c.GetString("mysql.user") + ":" + c.GetString("mysql.password") + "@" + c.GetString("mysql.protocol") + "/" + c.GetString("mysql.dbname") + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	dsn := "root:admin@tcp(localhost:3309)/aitguid?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

	if db, err = sql.Open("mysql", dsn); err != nil {
		db = connect(dsn, 10)
	}
	fmt.Println("db connected!!")

	return db
}

func connect(dsn string, count uint) *sql.DB {
	var err error
	var db *sql.DB
	if db, err = sql.Open("mysql", dsn); err != nil {
		if count > 1 {
			time.Sleep(time.Second * 2)
			count--
			fmt.Printf("retry... count:%v\n", count)
			connect(dsn, count)
		}
		panic(err.Error())
	}
	return (db)
}

// func SqlConnect() (database *gorm.DB) {
// 	var err error
// 	var db *gorm.DB
// 	// c := conf.GetMysqlConfig()
// 	// dsn := c.GetString("mysql.user") + ":" + c.GetString("mysql.password") + "@" + c.GetString("mysql.protocol") + "/" + c.GetString("mysql.dbname") + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
// 	dsn := "root:admin@tcp(localhost:3309)/aitguid?charset=utf8&parseTime=true&loc=Asia%2FTokyo"

// 	dialector := mysql.Open(dsn)

// 	if db, err = gorm.Open(dialector); err != nil {
// 		db = connect(dialector, 10)
// 	}
// 	fmt.Println("db connected!!")

// 	return db
// }

// func connect(dialector gorm.Dialector, count uint) *gorm.DB {
// 	var err error
// 	var db *gorm.DB
// 	if db, err = gorm.Open(dialector); err != nil {
// 		if count > 1 {
// 			time.Sleep(time.Second * 2)
// 			count--
// 			fmt.Printf("retry... count:%v\n", count)
// 			connect(dialector, count)
// 		}
// 		panic(err.Error())
// 	}
// 	return (db)
// }
