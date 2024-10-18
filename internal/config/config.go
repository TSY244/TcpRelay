package config

var (
	Version = "0.0.1"
)

type Config struct {
	ListenHost       string
	ConnHost         string
	EncryptionMethod string
	Password         string
	TcpRelayMode     string
}
