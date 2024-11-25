package rpc

import (
	"gomall/app/frontend/conf"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"gomall.local/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient userservice.Client
	once       sync.Once
)

func InitClient() {
	once.Do(func() {
		InitUserClient()
	})
}

func InitUserClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	UserClient, err = userservice.NewClient("user", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}
