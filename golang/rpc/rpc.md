先定义数据结构
```go
type HelloServer struct {
}

func NewGreetService() *HelloServer {
	return &HelloServer{}
}
```

定义方法
```go
// 这里需要注意 方法必须满足rpc的约束 有点类似http，需要有一个request一个response 
func (s *HelloServer) Greet(request string, resp *string) error {
	fmt.Println(request)
	*resp = request + "hello"
	return nil
}
```

提供服务
```go
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
```

对于服务端来说大概就这么几步，下面我们看客户端：
```go
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
```

最终结果：
服务端：
![服务端](image.png)
客户端：
![客户端](image-1.png)