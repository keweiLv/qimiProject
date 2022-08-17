package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Http Server start failed,err:%v", err)
		return
	}
}

func sayHello(writer http.ResponseWriter, request *http.Request) {
	// 解析模版
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("Parse Template failed,err:%v", err)
		return
	}
	// 渲染
	err = t.Execute(writer, "柯子")
	if err != nil {
		fmt.Println("Render Template failed,err:%v", err)
		return
	}
}
