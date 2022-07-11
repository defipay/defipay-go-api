package main

import (
	"fmt"
	"go_demo/defipay"
)

var localSigner = defipay.LocalSigner{
	PrivateKey: "1c75baa01f02c1457c553385da66c805da408e472832218bd7fc104573a448de",
}

var client = defipay.Client{
	Signer: localSigner,
	Env:    defipay.Sandbox(),
	Debug:  true,
}

func main() {
	GenerateKeyPair()
}

func createOrder() {
	result, apiError := client.CreateOrder("http://bkmvpggq.nc/olxtgwn", "http://trlwgf.et/ogtl", "bcb9559afce6477b92e0d05c01cd4be0", "100", "TRX", "25")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("createOrder：")
	fmt.Println(string(str))
}

func queryOrder() {
	result, apiError := client.QueryOrder("9QLDKBED")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("queryOrder：")
	fmt.Println(string(str))
}

func CreatePayoutOrder() {
	result, apiError := client.CreatePayoutOrder("www.baidu.com", "payout20220509001", "0.01", "ETH", "0x88a611Ceb5Cb3f0Fc002261F47CC85EbEd304412", "2", "")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("CreatePayoutOrder：")
	fmt.Println(string(str))
}

func queryPayoutOrder() {
	result, apiError := client.QueryPayoutOrder("V6JZYFXQ")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("queryPayoutOrder：")
	fmt.Println(string(str))
}

func queryRate() {
	result, apiError := client.QueryRate("ETH", "USDT")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("queryRate：")
	fmt.Println(string(str))
}

func queryBillCurrency() {
	result, apiError := client.QueryBillCurrency("1", "10")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("queryBillCurrency：")
	fmt.Println(string(str))
}

func tokenQuery() {
	result, apiError := client.TokenQuery("1", "10")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("tokenQuery：")
	fmt.Println(string(str))
}

func tokenDetail() {
	result, apiError := client.TokenDetail("2")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("tokenDetail：")
	fmt.Println(string(str))
}

func accountQuery() {
	result, apiError := client.AccountQuery()
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("accountQuery：")
	fmt.Println(string(str))
}

func queryOrderList() {
	result, apiError := client.QueryOrderList("1", "10")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("queryOrderList：")
	fmt.Println(string(str))
}

func orderDetail() {
	result, apiError := client.OrderDetail("3TJ0L2M5")
	if apiError != nil {
		fmt.Println("Error >>>>>>>>")
		fmt.Println(apiError)
		return
	}
	str, _ := result.Encode()
	fmt.Println("orderDetail：")
	fmt.Println(string(str))
}

func GenerateKeyPair() {
	apiSecret, apiKey := defipay.GenerateKeyPair()
	println("API_SECRET:", apiSecret)
	println("API_KEY:", apiKey)
}

func VerifyEcc() {
	verified := client.VerifyEcc("", "")
	fmt.Println("VerifyEcc：")
	fmt.Println(verified)
}
