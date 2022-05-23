package defipay

type Env struct {
	Host   string
	PubKey string
}

func Sandbox() Env {
	return Env{Host: "https://api-test.defipay.biz/api-service", PubKey: "0314c127b69c03545b49aaf365c56575ce9b1640cad8eb6ca47c34c8322fe9f4d4"}
}

func Prod() Env {
	return Env{Host: "https://api.custody.cobo.com", PubKey: "02adb46f0c10b5ec51d0df2183a812fdf7b330ef2c948e36cdb479f1af73a22753"}
}
