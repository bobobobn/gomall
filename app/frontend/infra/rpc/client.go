package rpc

import (
	"gomall/app/frontend/conf"
	"sync"

	"github.com/bobobobn/gomall/common/clientsuite"
	"github.com/cloudwego/kitex/client"
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
	opts           []client.Option
	once           sync.Once
)

func InitClient() {
	once.Do(func() {
		opts = []client.Option{
			client.WithSuite(clientsuite.CommonGrpcClientSuite{
				RegistryAddr: conf.GetConf().Registry.RegistryAddress[0],
			}),
		}
		InitUserClient()
		InitProductClient()
		InitCartClient()
		InitCheckoutClient()
		InitOrderClient()
	})
}

func InitUserClient() {
	var err error
	UserClient, err = userservice.NewClient("user", opts...)
	if err != nil {
		panic(err)
	}
}

func InitProductClient() {
	var err error
	ProductClient, err = productservice.NewClient("product", opts...)
	if err != nil {
		panic(err)
	}
}

func InitCartClient() {
	var err error
	CartClient, err = cartservice.NewClient("cart", opts...)
	if err != nil {
		panic(err)
	}
}

func InitCheckoutClient() {
	var err error
	CheckoutCLient, err = checkoutservice.NewClient("checkout", opts...)
	if err != nil {
		panic(err)
	}
}

func InitOrderClient() {
	var err error
	OrderClient, err = orderservice.NewClient("order", opts...)
	if err != nil {
		panic(err)
	}
}
