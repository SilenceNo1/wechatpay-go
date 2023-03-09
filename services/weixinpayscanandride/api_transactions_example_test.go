// Copyright 2021 Tencent Inc. All rights reserved.
//
// 公共出行平台代扣服务对外API
//
// 公共出行平台代扣服务对外API
//
// API version: 1.0.0

// Code generated by WechatPay APIv3 Generator based on [OpenAPI Generator](https://openapi-generator.tech); DO NOT EDIT.

package weixinpayscanandride_test

import (
	"context"
	"log"

	"github.com/SilenceNo1/wechatpay-go/core"
	"github.com/SilenceNo1/wechatpay-go/core/option"
	"github.com/SilenceNo1/wechatpay-go/services/weixinpayscanandride"
	"github.com/SilenceNo1/wechatpay-go/utils"
)

func ExampleTransactionsApiService_CreateTransaction() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := weixinpayscanandride.TransactionsApiService{Client: client}
	resp, result, err := svc.CreateTransaction(ctx,
		weixinpayscanandride.CreateTransactionRequest{
			Appid:       core.String("wxcbda96de0b165486"),
			SubAppid:    core.String("wxcbda96de0b165486"),
			SubMchid:    core.String("1900000109"),
			Description: core.String("地铁扣费"),
			Attach:      core.String("深圳分店"),
			OutTradeNo:  core.String("20150806125346"),
			TradeScene:  weixinpayscanandride.TRADESCENE_BUS.Ptr(),
			GoodsTag:    core.String("WXG"),
			ContractId:  core.String("Wx15463511252015071056489715"),
			NotifyUrl:   core.String("https://yoursite.com/wxpay.html"),
			Amount: &weixinpayscanandride.OrderAmount{
				Total:    core.Int64(600),
				Currency: core.String("CNY"),
			},
			BusInfo: &weixinpayscanandride.BusSceneInfo{
				StartTime:   core.String("2017-08-26T10:43:39+08:00"),
				LineName:    core.String("1路公车"),
				PlateNumber: core.String("粤B888888"),
			},
			MetroInfo: &weixinpayscanandride.MetroSceneInfo{
				StartTime:    core.String("2017-08-26T10:43:39+08:00"),
				EndTime:      core.String("2017-08-26T10:43:39+08:00"),
				StartStation: core.String("西单"),
				EndStation:   core.String("天安门西"),
			},
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call CreateTransaction err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}

func ExampleTransactionsApiService_QueryTransaction() {
	var (
		mchID                      string = "190000****"                               // 商户号
		mchCertificateSerialNumber string = "3775************************************" // 商户证书序列号
		mchAPIv3Key                string = "2ab9****************************"         // 商户APIv3密钥
	)

	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("/path/to/merchant/apiclient_key.pem")
	if err != nil {
		log.Printf("load merchant private key error:%s", err)
		return
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Printf("new wechat pay client err:%s", err)
		return
	}

	svc := weixinpayscanandride.TransactionsApiService{Client: client}
	resp, result, err := svc.QueryTransaction(ctx,
		weixinpayscanandride.QueryTransactionRequest{
			OutTradeNo: core.String("20150806125346"),
			SubMchid:   core.String("1900000109"),
		},
	)

	if err != nil {
		// 处理错误
		log.Printf("call QueryTransaction err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", result.Response.StatusCode, resp)
	}
}
