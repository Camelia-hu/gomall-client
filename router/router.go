package router

import (
	"context"
	"github.com/Camelia-hu/gomall-client/service"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	"github.com/hertz-contrib/obs-opentelemetry/tracing"
)

func RouterInit() {
	serviceName := "gomall-client"

	p := provider.NewOpenTelemetryProvider(
		provider.WithServiceName(serviceName),
		provider.WithExportEndpoint("localhost:4317"),
		provider.WithInsecure(),
	)
	defer p.Shutdown(context.Background())

	tracer, cfg := tracing.NewServerTracer()
	h := server.Default(tracer)
	h.Use(tracing.ServerMiddleware(cfg))
	user := h.Group("/user")
	{
		user.POST("/register", service.Register)
		user.GET("/login", service.Login)
		user.GET("/refreshToken", service.RefreshToken)
	}

	product := h.Group("/product")
	product.Use(service.AccessTokenAuth())
	{
		product.POST("/create", service.Create)
		product.GET("/delete", service.Delete)
		product.GET("/list", service.List)
		product.GET("/search", service.Search)
		product.GET("/get", service.Get)
	}

	cart := h.Group("/cart")
	cart.Use(service.AccessTokenAuth())
	{
		cart.POST("/add", service.Add)
		cart.GET("/get", service.GetCart)
		cart.GET("/delete", service.DeleteCart)
	}

	order := h.Group("/order")
	order.Use(service.AccessTokenAuth())
	{
		order.POST("/place", service.PlaceOrder)
		order.GET("/list", service.ListOrder)
	}

	payment := h.Group("/payment")
	payment.Use(service.AccessTokenAuth())
	{
		payment.POST("/createCredit", service.CreateCredit)
		payment.POST("/charge", service.Charge)
	}
	h.Spin()
}

//service ProductCatalogService {
//rpc ListProducts(ListProductsReq) returns (ListProductsResp) {}
//rpc GetProduct(GetProductReq) returns (GetProductResp) {}
//rpc SearchProducts(SearchProductsReq) returns (SearchProductsResp) {}
//rpc CreateProduct(CreateReq) returns (CreateResp) {}
//rpc DeleteProduct(DeleteReq) returns (DeleteResp) {}
//}
