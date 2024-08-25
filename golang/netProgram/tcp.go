package netProgram

import (
	"log"
	"net"
	"sync"
)

const tcp = "tcp4"

// 服务端
func TcpServer() {
	// 基于某个地址建立监听
	address := "127.0.0.1:5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Println("server is listening on ", address)

	// 接受连接请求
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// 处理连接
		log.Println("remot address:", conn.RemoteAddr())
	}
}

func TcpClient() {
	addr := "127.0.0.1:5678"
	wg := sync.WaitGroup{}
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			conn, err := net.Dial(tcp, addr)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			log.Println("connected to server:", conn.RemoteAddr())
		}
		wg.Done()
	}()
	wg.Wait()
}
