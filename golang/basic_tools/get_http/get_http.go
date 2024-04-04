package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Args struct {
	args []string
}

func (a *Args) Read_Args() {
	a.args = os.Args[1:]
}

// func Get_Res(url string) (content_str string) {
func Get_Res(url string) {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	// var content []byte
	var err_read error
	// content, err_read = io.ReadAll(res.Body)
	_, err_read = io.Copy(os.Stdout, res.Body)
	res.Body.Close()
	if err_read != nil {
		panic(err_read)
	}
	fmt.Println(os.Stdout)
	// return string(content)
}

func main() {
	a := new(Args)
	a.Read_Args()
	for _, url := range a.args {
		// fmt.Println(Get_Res(url))
		go Get_Res(url)
	}
}
