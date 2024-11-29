package rpc

import (
	"gomall/app/frontend/conf"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"gomall.local/rpc_gen/kitex_gen/cart/cartservice"
	"gomall.local/rpc_gen/kitex_gen/checkout/checkoutservice"
	"gomall.local/rpc_gen/kitex_gen/order/orderservice"
	"gomall.local/rpc_gen/kitex_gen/product/productservice"
	"gomall.local/rpc_gen/kitex_gen/user/userservice"
)

var (
	UserClient     userservice.Client
	ProductClient  productservice.Client
	CartClient     cartservice.Client
	CheckoutCLient checkoutservice.Client
	OrderClient    orderservice.Client
	once           sync.Once
)

func InitClient() {
	once.Do(func() {
		InitUserClient()
		InitProductClient()
		InitCartClient()
		InitCheckoutClient()
		InitOrderClient()
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

func InitProductClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	ProductClient, err = productservice.NewClient("product", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func InitCartClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	CartClient, err = cartservice.NewClient("cart", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func InitCheckoutClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	CheckoutCLient, err = checkoutservice.NewClient("checkout", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}

func InitOrderClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	OrderClient, err = orderservice.NewClient("order", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}
