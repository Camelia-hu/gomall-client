package service

import (
	"context"
	"github.com/Camelia-hu/gomall-client/module"
	"github.com/Camelia-hu/gomall-client/rpc"
	"github.com/Camelia-hu/gomall/product/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"log"
	"strconv"
)

func Create(ctx context.Context, c *app.RequestContext) {
	var p module.Product
	err := c.BindJSON(&p)
	if err != nil {
		log.Println("zaku～绑定参数失败喵～")
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "zaku～绑定参数失败喵～",
		})
		return
	}
	createResp, err := rpc.Prpc.CreateProduct(ctx, &product.CreateReq{
		Name:        p.Name,
		Description: p.Description,
		Picture:     p.Picture,
		Price:       p.Price,
		Categories:  p.Categories,
	})
	if err != nil {
		log.Println("就凭你也想创建商品？）")
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "就凭你也想创建商品？）",
		})
		return
	}
	c.JSON(200, utils.H{
		"code": 0,
		"msg":  "success）",
		"id":   createResp.Id,
	})
}

func Delete(ctx context.Context, c *app.RequestContext) {
	id := c.Query("id")
	Id, _ := strconv.Atoi(id)
	ID := uint32(Id)
	deleteResp, err := rpc.Prpc.DeleteProduct(ctx, &product.DeleteReq{Id: ID})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "删除失败喵～",
		})
		return
	}
	c.JSON(200, utils.H{
		"code": 0,
		"msg":  "它，永远的从你的世界里走了。",
		"resp": deleteResp.Is,
	})
}

func List(ctx context.Context, c *app.RequestContext) {
	category := c.Query("category")
	listResp, err := rpc.Prpc.ListProducts(ctx, &product.ListProductsReq{CategoryName: category})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "找不到哦喵～",
		})
		return
	}
	c.JSON(200, utils.H{
		"code":    0,
		"msg":     "那很好了）",
		"product": listResp.Products,
	})
}

func Search(ctx context.Context, c *app.RequestContext) {
	query := c.Query("query")
	searchResp, err := rpc.Prpc.SearchProducts(ctx, &product.SearchProductsReq{Query: query})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "找不到哦喵～",
		})
		return
	}
	c.JSON(200, utils.H{
		"code":    0,
		"msg":     "那很好了）",
		"product": searchResp.Results,
	})
}

func Get(ctx context.Context, c *app.RequestContext) {
	ID := c.Query("id")
	Id, _ := strconv.Atoi(ID)
	id := uint32(Id)
	getResp, err := rpc.Prpc.GetProduct(ctx, &product.GetProductReq{Id: id})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "找不到哦喵～",
		})
		return
	}
	c.JSON(200, utils.H{
		"code":    0,
		"msg":     "那很好了）",
		"product": getResp.Product,
	})
}
