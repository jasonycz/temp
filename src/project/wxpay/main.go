package main

import "github.com/objcoding/wxpay"

const (
	appID    = "wx4c8fd8d13ed047f5"
	mchID    = "1515254161"
	apiKey   = "0tbXF1kUrzrFEPKscth1ot6TzcvdejXZ"
	certPath = "/Users/yangchengzhi/yangchengzhi/test/src/project/wxpay/cert"
)

func main() {
	// 创建支付账户
	account := wxpay.NewAccount(appID, mchID, apiKey, false)
	// 新建微信支付客户端
	client := wxpay.NewClient(account)
	// 设置证书
	account.SetCertData(certPath)
	// 设置支付账户
	client.SetAccount(account)
	// 设置http请求超时时间
	client.SetHttpConnectTimeoutMs(2000)
	// 设置http读取信息流超时时间
	client.SetHttpReadTimeoutMs(1000)
	// 更改签名类型
	client.SetSignType(wxpay.HMACSHA256)
}
