package main

import (
	"context"
	"demo/hello_demo/conf"
	"demo/hello_demo/kitex_gen/hello/example"
	"demo/hello_demo/kitex_gen/hello/example/helloservice"
	"log"

	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	if err != nil {
		log.Fatal(err)
	}
	client, err := helloservice.NewClient("gomall.hello", client.WithResolver(r))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := client.HelloMethod(context.Background(), example.NewHelloReq())
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp)
}
