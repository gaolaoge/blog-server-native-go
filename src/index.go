package main

import (
	"fmt"
	"net/http"
)

func hande(w http.ResponseWriter, res *http.Request) {
	w.Write([]byte("hello world"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", hande)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("服务器开启失败：", err)
	}

}
