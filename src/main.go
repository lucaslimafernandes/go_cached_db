package main

import (
	"fmt"
	"net/http"
)

var p1 = map[string]string{
	"Lucas": "Lucas Lima",
}

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello\n%v\n", p1)

}

func main() {

	fmt.Println("http server")

	http.HandleFunc("/hello", hello)

	http.ListenAndServe("127.0.0.1:8090", nil)

}
