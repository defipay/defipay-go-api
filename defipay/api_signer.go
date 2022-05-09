package defipay

type ApiSigner interface {
	Sign(message string) string
	GetPublicKey() string
}
