package main

import (
	"fmt"
	"net/http"
)
/**
函数实现接口例子
 */
func main() {
	//http.ListenAndServe("",nil)

	var hanler http.Handler
	//HandlerFunc是接口类型；在这里竟然用到了闭包？
	hanler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		switch path{
		case "/":
			fmt.Fprintf(w, "Hello World!")
			break
		case "/index":
			fmt.Fprintf(w, "welcome to go web!")
			break
		case "/login":
			fmt.Fprintf(w, "please login!")
			break
		default:
			fmt.Fprintf(w, "invalid path %s",path)
		}

	})

	server := http.Server{
		Addr:              "127.0.0.1:8081",
		Handler:           hanler,
	}

	fmt.Println("Http Server Start...")
	server.ListenAndServe()

}
