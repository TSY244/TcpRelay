package tpFlag

import (
	"TcpRelay/config"
	"flag"
)

func banner() {
	banner := `
╔═╗  ╔╦╗┌─┐┌─┐╦═╗┌─┐┬  ┌─┐┬ ┬
╠═╣───║ │  ├─┘╠╦╝├┤ │  ├─┤└┬┘
╩ ╩   ╩ └─┘┴  ╩╚═└─┘┴─┘┴ ┴ ┴
Author: AU9U5T    Version: ` + config.Version + `
`
	println(banner)
}

func Flag(config *config.Config) {
	banner()
	flag.StringVar(&config.ListenHost, "lh", "0.0.0.0:10090", "IP address of the host you want to listen,for example: 0.0.0.0:10090")
	flag.StringVar(&config.ConnHost, "ch", "192.168.79.137:12345", "IP address of the host you want to connect,for example: 192.168.79.137:12345")
	flag.StringVar(&config.EncryptionMethod, "m", "6N_XOR", "encryption method")
	flag.StringVar(&config.Password, "p", "123456", "password")
	flag.StringVar(&config.TcpRelayMode, "t", "one2one", "TcpRelayMode: one2one")
	flag.Parse()
}

func PrintHelp() {
	flag.Usage()
}

func CheckFlag(config *config.Config) bool {
	// routine
	if config.ListenHost == "" || config.ConnHost == "" {
		return false
	}
	switch config.TcpRelayMode {
	case "one2one":
	default:
		return false

	}

	// encryption
	if config.EncryptionMethod == "6N_XOR" {
		if config.Password == "" {
			return false
		}
	}

	return true
}
