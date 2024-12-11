package service

import (
	"context"
	"github.com/Camelia-hu/gomall-client/module"
	"github.com/Camelia-hu/gomall-client/rpc"
	"github.com/Camelia-hu/gomall/payment/kitex_gen/payment"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"strconv"
)

func CreateCredit(ctx context.Context, c *app.RequestContext) {
	var credit module.CreditCard
	err := c.BindJSON(&credit)
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "绑定参数失败  " + err.Error(),
		})
		return
	}

	createCreditResp, err := rpc.Payrpc.CreateCredit(ctx, &payment.CreateCreditReq{
		Money:                     credit.Money,
		CreditCardNumber:          credit.CreditCardNumber,
		CreditCardCvv:             credit.CreditCardCvv,
		CreditCardExpirationYear:  credit.CreditCardExpirationYear,
		CreditCardExpirationMonth: credit.CreditCardExpirationMonth,
		Uid:                       int32(credit.Uid),
	})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "createCredit rpc err : " + err.Error(),
		})
		return
	}
	c.JSON(200, utils.H{
		"code": 0,
		"msg":  "success喵～",
		"id":   createCreditResp.Id,
	})
}

func Charge(ctx context.Context, c *app.RequestContext) {
	ID := c.Query("userId")
	Id, _ := strconv.Atoi(ID)
	uid := uint32(Id)
	Oid := c.Query("orderId")
	Amount := c.Query("amount")
	AmounT, _ := strconv.Atoi(Amount)
	amount := float32(AmounT)
	var credit module.CreditCard
	err := c.BindJSON(&credit)
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "bind err 喵～ ：" + err.Error(),
		})
		return
	}
	chargeResp, err := rpc.Payrpc.Charge(ctx, &payment.ChargeReq{
		Amount: amount,
		CreditCard: &payment.CreditCardInfo{
			CreditCardNumber:          credit.CreditCardNumber,
			CreditCardCvv:             credit.CreditCardCvv,
			CreditCardExpirationYear:  credit.CreditCardExpirationYear,
			CreditCardExpirationMonth: credit.CreditCardExpirationMonth,
		},
		OrderId: Oid,
		UserId:  uid,
	})
	if err != nil {
		c.JSON(400, module.Response{
			Code: 1,
			Msg:  "charge rpc err : " + err.Error(),
		})
		return
	}
	c.JSON(200, utils.H{
		"code": 0,
		"msg":  "success喵～",
		"id":   chargeResp.TransactionId,
	})
}
