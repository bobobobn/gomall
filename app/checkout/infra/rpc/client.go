package rpc

import (
	"gomall/app/checkout/conf"
	"sync"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"gomall.local/rpc_gen/kitex_gen/cart/cartservice"
	"gomall.local/rpc_gen/kitex_gen/payment/paymentservice"
	"gomall.local/rpc_gen/kitex_gen/product/productservice"
)

var (
	ProductClient productservice.Client
	CartClient    cartservice.Client
	PaymentClient paymentservice.Client
	once          sync.Once
)

func InitClient() {
	once.Do(func() {
		InitProductClient()
		InitCartClient()
		InitpaymentClient()
	})
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

func InitpaymentClient() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		panic(err)
	}
	PaymentClient, err = paymentservice.NewClient("payment", client.WithResolver(r))
	if err != nil {
		panic(err)
	}
}
