package service

import (
	"context"
	"github.com/Camelia-hu/gomall-client/module"
	"github.com/Camelia-hu/gomall-client/rpc"
	"github.com/Camelia-hu/gomall/auth/kitex_gen/auth"
	"github.com/Camelia-hu/gomall/user/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"log"
)

func Register(ctx context.Context, c *app.RequestContext) {
	email := c.Query("email")
	password := c.Query("password")
	rePassword := c.Query("rePassword")
	registerResp, err := rpc.Urpc.Register(ctx, &user.RegisterReq{
		Email:           email,
		Password:        password,
		ConfirmPassword: rePassword,
	})
	if err != nil {
		log.Println("rpc register err : ", err)
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "杂鱼～ " + err.Error(),
		})
		return
	}
	DeliveryResp, err := rpc.Arpc.DeliverTokenByRPC(ctx, &auth.DeliverTokenReq{UserId: registerResp.UserId})
	if err != nil {
		log.Println("rpc generate token err : ", err)
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "杂鱼杂鱼～ " + err.Error(),
		})
		return
	}
	c.JSON(200, utils.H{
		"code":         0,
		"msg":          "祝狗修金有愉快的一天～",
		"id":           registerResp.UserId,
		"accesstoken":  DeliveryResp.AccessToken,
		"refreshtoken": DeliveryResp.RefreshToken,
	})

}

func Login(ctx context.Context, c *app.RequestContext) {
	email := c.Query("email")
	password := c.Query("password")
	loginResp, err := rpc.Urpc.Login(ctx, &user.LoginReq{
		Email:    email,
		Password: password,
	})
	if err != nil {
		log.Println("zaku～")
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "zaku～",
		})
		return
	}
	DeliverResp, err := rpc.Arpc.DeliverTokenByRPC(ctx, &auth.DeliverTokenReq{UserId: loginResp.UserId})
	if err != nil {
		log.Println("zakuzaku～")
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "zakuzaku～",
		})
		return
	}

	c.JSON(200, utils.H{
		"code":         0,
		"msg":          "祝狗修金有愉快的一天～",
		"id":           loginResp.UserId,
		"accesstoken":  DeliverResp.AccessToken,
		"refreshtoken": DeliverResp.RefreshToken,
	})
}
