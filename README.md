# Defipay go API

defipay-go-api 是一個輕量級的 go 庫，用於與[Defipay API](http://doc.defipay.biz/)交互，提供完整的 API 覆蓋。


* [安裝](#安裝)
* [測試](#測試)
* [用法](#用法)
    * [初始化](#初始化)
        * [生成密鑰對](#生成密鑰對)
        * [初始化 RestClient](#初始化RestClient)
        * [初始化 ApiSigner](#初始化-apisigner)
    * [充值](#充值)
        * [充值請求下單](#充值請求下單)
        * [充值交易查詢](#充值交易查詢)
    * [提現](#提現)
        * [提現請求下單](#提現請求下單)
        * [提現交易查詢](#提現交易查詢)
    * [交易](#交易)
        * [獲取已入賬的交易詳情](#獲取已入賬的交易詳情)
        * [通過ID查詢獲取已確認的交易記錄](#通過ID查詢獲取已確認的交易記錄)
    * [賬戶余額查詢](#賬戶余額查詢)
        * [商戶賬戶余額查詢](#商戶賬戶余額查詢)
        * [支持賬單幣種查詢](#支持賬單幣種查詢)
        * [支持支付幣種查詢](#支持支付幣種查詢)
        * [支付幣種详情](#支付幣種详情)
    * [匯率查詢](#匯率查詢)
        * [幣種匯率查詢](#幣種匯率查詢)

## 安裝

添加依賴

``` 
go get github.com/btcsuite/btcd/btcec@v0.22
go get github.com/defipay/defipay-go-sdk/defipay
```

## 測試

```
main.go
```

## 用法

### 初始化

#### 生成密鑰對

```
import "github.com/defipay/defipay-go-sdk/defipay"

apiSecret, apiKey := defipay.GenerateKeyPair()
println("API_SECRET:", apiSecret)
println("API_KEY:", apiKey)
```

#### 初始化RestClient

```
import "github.com/defipay/defipay-go-sdk/defipay"

var client = defipay.Client{
	Signer: localSigner,
	Env:    defipay.Sandbox(),
	Debug:  true,
}
```

#### 初始化-apiSigner


`ApiSigner` 可以通過實例化
```
import "github.com/defipay/defipay-go-sdk/defipay"

var localSigner = defipay.LocalSigner{
	PrivateKey: "apiSecret ",
}
```

在某些情況下，您的私鑰無法導出，例如，您的私鑰在 aws kms 中，您應該通過實現`ApiSigner`接口傳入您自己的實現：

### 充值

#### 充值請求下單

```
client.CreateOrder("http://xcsewvb.ao/nhhcn", "http://da1.ao/djjke", "2a1ca442077e49edb91306701b3b0255", "100", "TRX", "25")
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":{"amount":"100","cashierUrl":"http://www.defipay.biz/customer/#/order/9QLDFBED","currency":"TRX","currencyLogoUrl":"https://defipay-test.oss-cn-hangzhou.aliyuncs.com/defipay_v_1.0/c205901ff60f4d588d4617bbaa0bc939.png","memberTransNo":"bcb9559afce6477b92e0d05c01cd4be0","tokenInfo":[{"address":"TVCsvmAvKyubwJtLweHhwDqKWtn8iXrWhJ","amount":"100","chainAssertDecimal":6,"chainAssertId":null,"chainId":null,"coinType":"Mainnet","displayName":"TRX","logoUrl":"https://defipay-test.oss-cn-hangzhou.aliyuncs.com/defipay_v_1.0/c205901ff60f4d588d4617bbaa0bc939.png","shortName":"TRX"}],"transNo":"9QLDFBED"},"success":true}
```
</details>

#### 充值交易查詢
```
client.QueryOrder("9QLDFBED")
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":{"amount":"100","billAmount":"","createTime":"1652076718","currency":"TRX","id":null,"memberSeq":"80000032","memberTransNo":"bcb9559afce6477b92e0d05c01cd4be0","settleTime":"0","shortName":"","state":"300","toAddress":"","transNo":"9QLDFBED"},"success":true}
```
</details>


### 提現
#### 提現請求下單

```
client.CreatePayoutOrder("http://mjyfxsum.ci/awckrxllk", "c412524cb01c4bd383adgc903a08284d", "0.01", "ETH", "0x88a611Ceb5Cb3f0Fc002261F47CC85EbEd304412", "2", "")
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":{"transNo":"C5LOTLWK","memberTransNo":"c412524cb01c4bd383adgc903a08284d","currency":"ETH","amount":"0.01","tokenId":2,"tokenAmount":"0.01"},"success":true}
```
</details>

#### 提現交易查詢
```
client.QueryPayoutOrder("9QLDFBED")
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":{"amount":"100","billAmount":"","createTime":"1652076718","currency":"TRX","id":null,"memberSeq":"80000032","memberTransNo":"bcb9559afce6477b92e0d05c01cd4be0","settleTime":"0","shortName":"","state":"300","toAddress":"","transNo":"9QLDFBED"},"success":true}
```
</details>

### 交易

#### 獲取已入賬的交易列表

```
client.QueryOrderList("1", "10")
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":[{"id":4546,"transNo":"C5LOTLWK","memberTransNo":"payouttest111","memberSeq":"80000029","amount":"1000","currency":"USDT","toAddress":"0x3531C5F7540aDC5e5d640De11DE524cD379CC717","shortName":"ETH","billAmount":"0.374510748524758788","state":"300","createTime":"1651906147","settleTime":"0"},{"id":3170,"transNo":"1SQKTD17","memberTransNo":"test122234124","memberSeq":"80000029","amount":"1000","currency":"USDT","toAddress":"","shortName":"","billAmount":"","state":"300","createTime":"1651905865","settleTime":"0"}],"success":true}
```
</details>

#### 查詢獲取已確認的交易記錄
```
client.OrderDetail("29N3FVHO")
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":{"id":2873,"transNo":"29N3FVHO","memberTransNo":"202205050000000001","memberSeq":"80000020","amount":"0.1","currency":"ETH","toAddress":"0x8034761a6b9b2aa11f5b5fc9c5539a8061e0d759","shortName":"ETH","billAmount":"0.1","state":"200","createTime":"1651750925","settleTime":"1651754339"},"success":true}
```

</details>

### 賬戶余額查詢
#### 商戶賬戶余額查詢

```
client.AccountQuery()
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":[{"totalAmount":0,"frozenAmount":0,"availableAmount":0,"tokenId":2,"shortName":"ETH","name":"Ether","displayName":"ETH","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/ba40da70bed74489a7ed6adaed495763.png"},{"totalAmount":0,"frozenAmount":0,"availableAmount":0,"tokenId":3,"shortName":"USDT","name":"Tether","displayName":"USDT-ERC20","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/1b88c0c2dba04080bd3165843de3ffae.png"}],"success":true}
```
</details>

#### 支持賬單幣種查詢

```
client.QueryBillCurrency("1", "10")
```

<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":[{"assertId":"AED","typeIsCrypto":0},{"assertId":"AUD","typeIsCrypto":0},{"assertId":"BRC","typeIsCrypto":0},{"assertId":"BYN","typeIsCrypto":0},{"assertId":"CAD","typeIsCrypto":0},{"assertId":"CHF","typeIsCrypto":0},{"assertId":"CLP","typeIsCrypto":0},{"assertId":"CNY","typeIsCrypto":0},{"assertId":"DEM","typeIsCrypto":0},{"assertId":"DKK","typeIsCrypto":0}],"success":true}
```
</details>

#### 支持支付幣種查詢

```
client.TokenQuery("1", "10")
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":[{"id":1,"name":"BTC","displayName":"BTC","shortName":"BTC","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/8f6e5e2382f94028b87307ad5c73c52e.png","chainAssertId":null,"chainAssertDecimal":"6"},{"id":2,"name":"Ether","displayName":"ETH","shortName":"ETH","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/ba40da70bed74489a7ed6adaed495763.png","chainAssertId":null,"chainAssertDecimal":"18"},{"id":3,"name":"Tether","displayName":"USDT-ERC20","shortName":"USDT","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/1b88c0c2dba04080bd3165843de3ffae.png","chainAssertId":"0xdac17f958d2ee523a2206206994597c13d831ec7","chainAssertDecimal":"6"},{"id":4,"name":"Binance Coin","displayName":"BNB","shortName":"BNB","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/6300d391da1642c58c6673f32235db89.png","chainAssertId":null,"chainAssertDecimal":"18"},{"id":5,"name":"USDC","displayName":"USDC-ERC20","shortName":"USDC","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/b5afa9c1d46842cea775d3e63c9287b7.png","chainAssertId":"0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48","chainAssertDecimal":"6"},{"id":6,"name":"Ripple","displayName":"XRP-BEP20","shortName":"XRP","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/7df1c4bc20054894b52b4aa6a7ae1e81.png","chainAssertId":"0x1d2f0da169ceb9fc7b3144628db156f3f6c60dbe","chainAssertDecimal":"18"},{"id":7,"name":"Cardano","displayName":"ADA-BEP20","shortName":"ADA","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/35d6f3167e9c4a9dae0ef34c50a1deb0.png","chainAssertId":"0x3ee2200efb3400fabb9aacf31297cbdd1d435d47","chainAssertDecimal":"18"},{"id":8,"name":"Solana","displayName":"SOL","shortName":"SOL","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/068abeb1076c40189d64c69cf70ed6e6.png","chainAssertId":null,"chainAssertDecimal":"18"},{"id":9,"name":"Luna Coin","displayName":"LUNA","shortName":"LUNA","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/a8a96dd9240c425291bb87178aae935c.png","chainAssertId":null,"chainAssertDecimal":"18"},{"id":10,"name":"AVAX","displayName":"AVAX-C","shortName":"AVAX","logoUrl":"https:\/\/defipay-test.oss-cn-hangzhou.aliyuncs.com\/defipay_v_1.0\/62f25f5cf608415587a7cda95d9238c0.png","chainAssertId":null,"chainAssertDecimal":"18"}],"success":true}
```
</details>

#### 支付幣種详情

```
client.TokenDetail("2")
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":{"availableAmount":0,"createTime":1651909485,"frozenAmount":0,"id":245,"memberSeq":"80000032","shortName":"ETH","tokenId":2,"totalAmount":0,"updateTime":1651909485,"version":0},"success":true}
```
</details>


### 匯率查詢
#### 幣種匯率查詢

```
client.QueryRate("ETH", "USDT")
```
<details>
<summary>響應視圖</summary>

```
{"code":200,"msg":"OK","data":{"rate":"2669.8356867640886695401071913","rateTime":1651896014},"success":true}
```
</details>

