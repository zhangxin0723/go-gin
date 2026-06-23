package main

import (
	"io"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	if r.Method != "GET" {
		byteData, _ := io.ReadAll(r.Body)
		log.Println(string(byteData))
	}
	log.Println(r.Header)
	w.Write([]byte("Hello World!"))
}

func main() {

	http.HandleFunc("/index", Index)
	log.Println("服务已启动....")
	http.ListenAndServe(":8899", nil)
}
