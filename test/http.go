package main

import (
	"fmt"
	"net/http"
)

//go语言实现 http服务端
func main() {

	//注册路由
	http.HandleFunc("/", sayhello)
	//建立监听
	err := http.ListenAndServe("127.0.0.1:8080", nil)

	if err != nil {
		fmt.Println("网络错误")
		return
	}

}

func sayhello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("message : hello"))
	//fmt.Fprint(writer,"<p style='color:red;'>你好啊哈哈哈哈哈</p>")

}
