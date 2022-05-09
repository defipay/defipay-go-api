package defipay

type Env struct {
	Host   string
	PubKey string
}

func Sandbox() Env {
	return Env{Host: "http://api-test.defipay.biz/api-service", PubKey: "03412208e920ba78d97c33c1476db11506cb6d4b3fc218a8207ca523c8d392a3f0"}
}

func Prod() Env {
	return Env{Host: "https://api.custody.cobo.com", PubKey: "02c3e5bacf436fbf4da78597e791579f022a2e85073ae36c54a361ff97f2811376"}
}
