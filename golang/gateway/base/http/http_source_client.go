package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, err := client.Get("https://baidu.com")
	defer resp.Body.Close()
	if err != nil {
		panic(err)
	}
	output, _ := io.ReadAll(resp.Body)
	fmt.Println(string(output))
}
