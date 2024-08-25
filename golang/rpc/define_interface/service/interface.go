package service

type HelloService interface {
	Greet(request string, resp *string) error
}

type HelloRequest struct {
	Message string
}

type HelloResponse struct {
	Message string
}
