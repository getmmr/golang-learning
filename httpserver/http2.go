package httpserver

import (
	"fmt"
	"net/http"
	"os"
)

func Test2() {
	http.HandleFunc("/hello", HandleHello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error:%v\n", err)
	}
}

func HandleHello(writer http.ResponseWriter, req *http.Request) {
	clientInfo := fmt.Sprintf("HandleHello: %s %s %s\n", req.RemoteAddr, req.Method, req.RequestURI)
	fmt.Printf(clientInfo)
	_, _ = fmt.Fprintf(writer, clientInfo)
}
