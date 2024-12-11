package rpc

import (
	"context"
	"github.com/Camelia-hu/gomall/auth/kitex_gen/auth/authservice"
	"github.com/Camelia-hu/gomall/cart/kitex_gen/cart/cartservice"
	"github.com/Camelia-hu/gomall/order/kitex_gen/order/orderservice"
	"github.com/Camelia-hu/gomall/payment/kitex_gen/payment/paymentservice"
	"github.com/Camelia-hu/gomall/product/kitex_gen/product/productcatalogservice"
	"github.com/Camelia-hu/gomall/user/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	consul "github.com/kitex-contrib/registry-consul"
	"log"
	"time"
)

var (
	Urpc   userservice.Client
	Arpc   authservice.Client
	Prpc   productcatalogservice.Client
	Crpc   cartservice.Client
	Orpc   orderservice.Client
	Payrpc paymentservice.Client
)

func RpcInit() {
	p := provider.NewOpenTelemetryProvider(
		provider.WithInsecure(),
		provider.WithServiceName("gomall-rpc"),
		provider.WithExportEndpoint("localhost:4317"),
	)
	defer p.Shutdown(context.Background())
	userR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("UserRpc find err : ", err)
	}
	Urpc = userservice.MustNewClient("user", client.WithResolver(userR), client.WithRPCTimeout(3*time.Second), client.WithSuite(tracing.NewClientSuite()))

	authR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("AuthRpc find err : ", err)
	}
	Arpc = authservice.MustNewClient("auth", client.WithResolver(authR), client.WithRPCTimeout(3*time.Second), client.WithSuite(tracing.NewClientSuite()))

	productR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("ProductRpc find err : ", err)
	}
	Prpc = productcatalogservice.MustNewClient("product", client.WithResolver(productR), client.WithRPCTimeout(3*time.Second), client.WithSuite(tracing.NewClientSuite()))

	cartR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("CartRpc find err : ", err)
	}
	Crpc = cartservice.MustNewClient("cart", client.WithResolver(cartR), client.WithRPCTimeout(3*time.Second), client.WithSuite(tracing.NewClientSuite()))

	orderR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("OrderRpc find err : ", err)
	}
	Orpc = orderservice.MustNewClient("order", client.WithResolver(orderR), client.WithRPCTimeout(3*time.Second), client.WithSuite(tracing.NewClientSuite()))

	paymentR, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Panicln("PaymentRpc find err : ", err)
	}
	Payrpc = paymentservice.MustNewClient("payment", client.WithResolver(paymentR), client.WithRPCTimeout(3*time.Second), client.WithSuite(tracing.NewClientSuite()))
}
