package main

import (
	"TcpRelay/config"
	OneToOne "TcpRelay/internal/forwardingMethod"
	"TcpRelay/internal/tpFlag"
	"TcpRelay/pkg/socket"
	"strings"
)

func main() {
	cfg := &config.Config{}
	tpFlag.Flag(cfg)

	if !tpFlag.CheckFlag(cfg) {
		tpFlag.PrintHelp()
		return
	}

	server, err := socket.NewServer(cfg.ListenHost)
	if err != nil {
		panic(err)
	}
	defer server.Close()

	r := strings.Split(cfg.ConnHost, ":")
	if len(r) != 2 {
		panic("connect address error")
	}
	ip, port := r[0], r[1]

	err = OneToOne.Do(server, ip, port, cfg.EncryptionMethod, cfg.Password)

}
