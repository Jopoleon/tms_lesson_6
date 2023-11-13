package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("server started on port 8080")
	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":8080", nil)
	//fmt.Println("this is golang program on linux")
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Request happened, method: ", req.Method)
	fmt.Println("Remote addr: ", req.RemoteAddr)
	fmt.Println("!!!! This is a NEW verions of server !!!!")
	//fmt.Println(req.Header.Get("X-Real-IP"))
	fmt.Fprintf(w, "hello there! Your IP is: %s", req.RemoteAddr)
}
