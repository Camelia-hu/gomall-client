package service

import (
	"context"
	"github.com/Camelia-hu/gomall-client/module"
	"github.com/Camelia-hu/gomall-client/rpc"
	"github.com/Camelia-hu/gomall/order/kitex_gen/order"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

func PlaceOrder(ctx context.Context, c *app.RequestContext) {
	var Order module.OrderReq
	err := c.Bind(&Order)
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "orderReq bind err : " + err.Error(),
		})
		return
	}
	var orders []*order.OrderItem
	for _, item := range Order.OrderItems {
		oneOrder := &order.OrderItem{
			Item: &order.CartItem{
				ProductId: item.ProductId,
				Quantity:  item.Quantity,
			},
			Cost: item.Cost,
		}
		orders = append(orders, oneOrder)
	}
	placeResp, err := rpc.Orpc.PlaceOrder(ctx, &order.PlaceOrderReq{
		UserId:       Order.UserId,
		UserCurrency: Order.UserCurrency,
		Email:        Order.Email,
		OrderItems:   orders,
	})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "place order rpc err : " + err.Error(),
		})
		return
	}
	c.JSON(200, utils.H{
		"code":    0,
		"msg":     "success喵～",
		"orderId": placeResp.Order.OrderId,
	})
}

func ListOrder(ctx context.Context, c *app.RequestContext) {
	UID := c.Query("uid")
	Uid, _ := strconv.Atoi(UID)
	uid := uint32(Uid)
	listResp, err := rpc.Orpc.ListOrder(ctx, &order.ListOrderReq{UserId: uid})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "list order rpc err : " + err.Error(),
		})
		return
	}
	c.JSON(200, utils.H{
		"code":   0,
		"msg":    "success喵～",
		"orders": listResp.Orders,
	})
}
