package service

import (
	"context"
	"github.com/Camelia-hu/gomall-client/module"
	"github.com/Camelia-hu/gomall-client/rpc"
	"github.com/Camelia-hu/gomall/cart/kitex_gen/cart"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

func Add(ctx context.Context, c *app.RequestContext) {
	var item module.Item
	err := c.BindJSON(&item)
	id := c.Query("uid")
	Uid, _ := strconv.Atoi(id)
	uid := uint32(Uid)
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "绑定参数失败喵～",
		})
		return
	}

	_, err = rpc.Crpc.AddItem(ctx, &cart.AddItemReq{
		UserId: uid,
		Item: &cart.CartItem{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		},
	})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "rpc调用失败喵～",
		})
		return
	}
	c.JSON(200, utils.H{
		"code": 0,
		"msg":  "加购成功～",
	})
}

func GetCart(ctx context.Context, c *app.RequestContext) {
	id := c.Query("uid")
	Uid, _ := strconv.Atoi(id)
	uid := uint32(Uid)
	getResp, err := rpc.Crpc.GetCart(ctx, &cart.GetCartReq{UserId: uid})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "牙勒牙勒～",
		})
	}
	var items []module.Item
	for _, item := range getResp.Cart.Items {
		oneItem := module.Item{
			ProductId: item.ProductId,
			Quantity:  item.Quantity,
		}
		items = append(items, oneItem)
	}
	c.JSON(200, utils.H{
		"code":  0,
		"msg":   "okk～",
		"items": items,
	})
}

func DeleteCart(ctx context.Context, c *app.RequestContext) {
	id := c.Query("uid")
	Uid, _ := strconv.Atoi(id)
	uid := uint32(Uid)
	_, err := rpc.Crpc.EmptyCart(ctx, &cart.EmptyCartReq{UserId: uid})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "牙勒牙勒～",
		})
		return
	}
	c.JSON(200, module.Response{
		Code: 0,
		Msg:  "okk～",
	})
}
