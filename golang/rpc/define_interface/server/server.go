package main

import (
	"define_interface/service"
	"fmt"
)

var _ service.HelloService = (*HelloServer)(nil)

type HelloServer struct {
}

func (s *HelloServer) Greet(request string, resp *string) error {
	fmt.Println(request)
	*resp = "Hello" + request
	return nil
}

func main() {

}
