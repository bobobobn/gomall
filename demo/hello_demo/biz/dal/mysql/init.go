package mysql

import (
	"demo/hello_demo/biz/model"
	"demo/hello_demo/conf"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	dsn := fmt.Sprintf(conf.GetConf().MySQL.DSN,
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	type Version struct {
		Version string
	}
	var v Version
	err = DB.Raw("SELECT VERSION() AS version").Scan(&v).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("MySQL Version:", v.Version)
	DB.AutoMigrate(&model.User{})
}