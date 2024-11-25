package dal

import (
	"gomall/rpc_gen/biz/dal/mysql"
	"gomall/rpc_gen/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
