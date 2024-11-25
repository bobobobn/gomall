package dal

import (
	"demo/hello_demo/biz/dal/mysql"
)

func Init() {
	// redis.Init()
	mysql.Init()
}
