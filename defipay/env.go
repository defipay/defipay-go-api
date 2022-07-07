package defipay

type Env struct {
	Host   string
	PubKey string
}

func Sandbox() Env {
	return Env{Host: "https://api-test.defipay.biz/api-service", PubKey: "02a17fffb024cce6220ddf91b40711dc15fd8f830e23f6160c6a4eac8bc0eba820"}
}

func Prod() Env {
	return Env{Host: "https://api.custody.cobo.com", PubKey: "02adb46f0c10b5ec51d0df2183a812fdf7b330ef2c948e36cdb479f1af73a22753"}
}
