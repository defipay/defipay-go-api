package defipay

import (
	"encoding/hex"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/btcsuite/btcd/btcec"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Client struct {
	Signer LocalSigner
	Env    Env
	Debug  bool
}

func (c Client) CreateOrder(notifyUrl string, returnUrl string, memberTransNo string, amount string, currency string, tokenIds string) (*simplejson.Json, *ApiError) {
	return c.Request("POST", "/v1/external/pay/create", map[string]string{
		"notifyUrl":     notifyUrl,
		"returnUrl":     returnUrl,
		"memberTransNo": memberTransNo,
		"amount":        amount,
		"currency":      currency,
		"tokenIds":      tokenIds,
	})
}

func (c Client) QueryOrder(transNo string) (*simplejson.Json, *ApiError) {
	return c.Request("POST", "/v1/external/pay/query", map[string]string{
		"transNo": transNo,
	})
}

func (c Client) CreatePayoutOrder(notifyUrl string, memberTransNo string, amount string, currency string, toAddress string, tokenId string, payAmount string) (*simplejson.Json, *ApiError) {
	return c.Request("POST", "/v1/external/payout/create", map[string]string{
		"notifyUrl":     notifyUrl,
		"memberTransNo": memberTransNo,
		"amount":        amount,
		"currency":      currency,
		"toAddress":     toAddress,
		"tokenId":       tokenId,
		"payAmount":     payAmount,
	})
}

func (c Client) QueryPayoutOrder(transNo string) (*simplejson.Json, *ApiError) {
	return c.Request("POST", "/v1/external/payout/query", map[string]string{
		"transNo": transNo,
	})
}

func (c Client) QueryRate(base string, quote string) (*simplejson.Json, *ApiError) {
	return c.Request("POST", "/v1/external/rate/query", map[string]string{
		"base":  base,
		"quote": quote,
	})
}

func (c Client) QueryBillCurrency(offset string, limit string) (*simplejson.Json, *ApiError) {
	return c.Request("POST", "/v1/external/billCurrency/query", map[string]string{
		"offset": offset,
		"limit":  limit,
	})
}

func (c Client) TokenQuery(offset string, limit string) (*simplejson.Json, *ApiError) {
	return c.Request("POST", "/v1/external/token/query", map[string]string{
		"offset": offset,
		"limit":  limit,
	})
}

func (c Client) TokenDetail(tokenId string) (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/external/token/getDetail/"+tokenId, map[string]string{
		"tokenId": tokenId,
	})
}

func (c Client) AccountQuery() (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/external/account/query", map[string]string{})
}

func (c Client) QueryOrderList(offset string, limit string) (*simplejson.Json, *ApiError) {
	return c.Request("POST", "/v1/external/order/list", map[string]string{
		"offset": offset,
		"limit":  limit,
	})
}

func (c Client) OrderDetail(transNo string) (*simplejson.Json, *ApiError) {
	return c.Request("GET", "/v1/external/order/getDetail", map[string]string{
		"transNo": transNo,
	})
}

func (c Client) Request(method string, path string, params map[string]string) (*simplejson.Json, *ApiError) {
	jsonString := c.request(method, path, params)
	json, _ := simplejson.NewJson([]byte(jsonString))
	success, _ := json.Get("success").Bool()
	if !success {
		Success, _ := json.Get("success").Bool()
		Message, _ := json.Get("msg").String()
		Code, _ := json.Get("code").Int()
		apiError := ApiError{
			Success: Success,
			Message: Message,
			Code:    Code,
		}
		return nil, &apiError
	}

	result := json.Get("data")
	return result, nil
}

func (c Client) request(method string, path string, params map[string]string) string {
	httpClient := &http.Client{}
	nonce := fmt.Sprintf("%d", time.Now().Unix()*1000)
	sorted := SortParams(params)
	var req *http.Request
	if method == "POST" {
		req, _ = http.NewRequest(method, c.Env.Host+path, strings.NewReader(sorted))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	} else {
		req, _ = http.NewRequest(method, c.Env.Host+path+"?"+sorted, strings.NewReader(""))
	}
	content := strings.Join([]string{method, path, nonce, sorted}, "|")

	req.Header.Set("Biz-Api-Key", c.Signer.GetPublicKey())
	req.Header.Set("Biz-Api-Nonce", nonce)
	req.Header.Set("Biz-Api-Signature", c.Signer.Sign(content))

	if c.Debug {
		fmt.Println("request >>>>>>>>")
		fmt.Println(method, "\n", path, "\n", params, "\n", content, "\n", req.Header)
	}
	resp, _ := httpClient.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	//timestamp := resp.Header["Biz-Timestamp"][0]
	//signature := resp.Header["Biz-Resp-Signature"][0]
	//if c.Debug {
	//	fmt.Println("response <<<<<<<<")
	//	fmt.Println(string(body), "\n", timestamp, "\n", signature)
	//}
	return string(body)
}

func SortParams(params map[string]string) string {
	keys := make([]string, len(params))
	i := 0
	for k := range params {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	sorted := make([]string, len(params))
	i = 0
	for _, k := range keys {
		//sorted[i] = k + "=" + url.QueryEscape(params[k])
		sorted[i] = k + "=" + params[k]
		i++
	}
	return strings.Join(sorted, "&")
}

func (c Client) VerifyEcc(message string, signature string) bool {
	pubKeyBytes, _ := hex.DecodeString(c.Env.PubKey)
	pubKey, _ := btcec.ParsePubKey(pubKeyBytes, btcec.S256())

	sigBytes, _ := hex.DecodeString(signature)
	sigObj, _ := btcec.ParseSignature(sigBytes, btcec.S256())

	verified := sigObj.Verify([]byte(Hash256x2(message)), pubKey)
	return verified
}
