package service

//后续看看能不能用hertz的jwt库来实现
import (
	"context"
	"errors"
	"github.com/Camelia-hu/gomall-client/module"
	"github.com/Camelia-hu/gomall-client/rpc"
	"github.com/Camelia-hu/gomall/auth/kitex_gen/auth"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"log"
)

func AccessTokenAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		authorization := c.GetHeader("Authorization")
		if string(authorization) == "" || string(authorization[:7]) != "Bearer " {
			c.JSON(400, module.Response{
				Code: 1,
				Msg:  "你不是狗修金（",
			})
			c.Abort()
			return
		}
		token := authorization[7:]
		verifyResp, err := rpc.Arpc.VerifyTokenByRPC(ctx, &auth.VerifyTokenReq{Token: string(token)})
		if err != nil && errors.Is(err, errors.New("token 过期喵～")) {
			c.JSON(400, module.Response{
				Code: 1,
				Msg:  "帮您刷新一下令牌喵～",
			})
			c.Abort()
			return
		}
		if err != nil || verifyResp.Res != true {
			c.JSON(400, module.Response{
				Code: 1,
				Msg:  "杂鱼假扮狗修金（",
			})
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}

func RefreshToken(ctx context.Context, c *app.RequestContext) {
	accessToken := c.Query("accessToken")
	refreshToken := c.Query("refreshToken")
	refreshResp, err := rpc.Arpc.ReFreshTokenByRPC(ctx, &auth.RefreshReq{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		log.Println("哈？不信，回去重新登陆我看看（")
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "哈？不信，回去重新登陆我看看（",
		})
		return
	}
	c.JSON(200, utils.H{
		"accessToken":  refreshResp.AccessToken,
		"refreshToken": refreshResp.RefreshToken,
	})
}
