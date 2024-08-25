package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	var resp string
	err = conn.Call("GreetService.Greet", "client", &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
}
