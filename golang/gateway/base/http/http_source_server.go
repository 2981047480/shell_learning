// package http_source_sever

package main

import "net/http"

func server() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello from http source server"))
	})
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello from /hello"))
	})
	http.HandleFunc("/hello/world", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world from /hello/world"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}
