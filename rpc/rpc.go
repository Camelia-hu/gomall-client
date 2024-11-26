package rpc

import (
	"github.com/Camelia-hu/gomall/auth/kitex_gen/auth/authservice"
	"github.com/Camelia-hu/gomall/cart/kitex_gen/cart/cartservice"
	"github.com/Camelia-hu/gomall/product/kitex_gen/product/productcatalogservice"
	"github.com/Camelia-hu/gomall/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"time"
)

var (
	Urpc userservice.Client
	Arpc authservice.Client
	Prpc productcatalogservice.Client
	Crpc cartservice.Client
)

func RpcInit() {
	userR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("UserRpc find err : ", err)
	}
	Urpc = userservice.MustNewClient("user", client.WithResolver(userR), client.WithRPCTimeout(3*time.Second))

	authR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("AuthRpc find err : ", err)
	}
	Arpc = authservice.MustNewClient("auth", client.WithResolver(authR), client.WithRPCTimeout(3*time.Second))

	productR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("ProductRpc find err : ", err)
	}
	Prpc = productcatalogservice.MustNewClient("product", client.WithResolver(productR), client.WithRPCTimeout(3*time.Second))

	cartR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("CartRpc find err : ", err)
	}
	Crpc = cartservice.MustNewClient("cart", client.WithResolver(cartR), client.WithRPCTimeout(3*time.Second))
}