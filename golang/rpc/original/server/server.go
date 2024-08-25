package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloServer struct {
}

func NewGreetService() *HelloServer {
	return &HelloServer{}
}

// 这里需要注意 方法必须满足rpc的约束 有点类似http，需要有一个request一个response
func (s *HelloServer) Greet(request string, resp *string) error {
	fmt.Println(request)
	*resp = request + "hello"
	return nil
}

func main() {
	rpc.RegisterName("GreetService", NewGreetService())

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}

		go rpc.ServeConn(conn)
	}

}
