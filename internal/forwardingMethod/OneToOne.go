package forwardingMethod

import (
	"TcpRelay/pkg/encryption"
	"TcpRelay/pkg/socket"
	"errors"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

//func test(server *socket.Server, clientIp, clientPort string) error {
//	// test server read msg and write msg
//	if server == nil {
//		return errors.New("server or client is nil")
//	}
//	defer server.Close()
//
//	srvConn, err := server.Accept()
//	if err != nil {
//		return err
//	}
//	defer srvConn.Close()
//
//	client, err := socket.NewClient(clientIp, clientPort)
//	if err != nil {
//		return err
//	}

func Do(server *socket.Server, clientIp, clientPort string, encryptionMethod string, key string) error {
	if server == nil {
		return errors.New("server or client is nil")
	}
	defer server.Close()

	var srvConn, cliConn *net.Conn

	toCliSignal := make(chan struct{}, 1)
	toSrvSignal := make(chan struct{}, 1)
	exitSignal := make(chan struct{}, 1)

	flag := true
	lock := sync.Mutex{}

	srvKeyCopy := key
	cliKeyCopy := key
	go func() {
		sc, err := server.Accept()
		toCliSignal <- struct{}{}
		if err != nil {
			return
		}
		defer sc.Close()
		srvConn = &sc
		buf := make([]byte, 1024)
		for {
			n, err := (*srvConn).Read(buf)
			// EOF -> continue
			if errors.Is(err, io.EOF) {
				continue
			}
			if err != nil {
				log.Println(" read error:", err)
				exitSignal <- struct{}{}
				break
			}
			log.Println("[192.168.79.1:10090] read from :", (*srvConn).RemoteAddr(), " size is ", n)
			if flag {
				<-toSrvSignal
				log.Println("server signal")
				flag = false
			}
			lock.Lock()
			_, err = socket.Write(*cliConn, encryption.Encrypt(encryptionMethod, buf[:n], srvKeyCopy))
			lock.Unlock()
			if err != nil {
				log.Println("client write error:", err)
				exitSignal <- struct{}{}
				break
			}
			log.Println("[192.168.79.1:10090] write to :", (*cliConn).RemoteAddr())
		}
	}()

	go func() {
		<-toCliSignal
		log.Println("client signal")
		client, err := socket.NewClient(clientIp, clientPort)
		if err != nil {
			log.Println("[192.168.79.1:10090] client connect error:", err)
			exitSignal <- struct{}{}
			return
		}
		defer client.Close()
		cliConn = client.Conn
		toSrvSignal <- struct{}{}
		buf := make([]byte, 1024)
		for {
			n, err := (*cliConn).Read(buf)
			if errors.Is(err, io.EOF) {
				continue
			}
			if err != nil {
				log.Println("client read error:", err)
				exitSignal <- struct{}{}
				break
			}
			log.Println("[192.168.79.1:10090] read from :", (*cliConn).RemoteAddr(), " size is ", n)
			lock.Lock()
			_, err = socket.Write(*srvConn, encryption.Decrypt(encryptionMethod, buf[:n], cliKeyCopy))
			lock.Unlock()
			if err != nil {
				log.Println("server write error:", err)
				exitSignal <- struct{}{}
				break
			}
			log.Println("[192.168.79.1:10090] write to :", (*srvConn).RemoteAddr())
		}
		exitSignal <- struct{}{}
	}()
	for {
		select {
		case <-exitSignal:
			log.Println("exit.....")
			os.Exit(1)
		default:
			//log.Println("default")
		}
	}
	return nil
}
