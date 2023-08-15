package httpserver

import (
	"fmt"
	"net/http"
	"os"
)

type HttpHandle struct {
}

func (h *HttpHandle) ServeHTTP(respWriter http.ResponseWriter, req *http.Request) {
	fmt.Printf("ClientAddr: %s\n", req.Host)

	respWriter.Header().Add("Key-Hello", "Value-World")
	respWriter.Header().Set("Content-Type", "text/plain; charset=utf-8")

	respWriter.WriteHeader(200)

	clientInfo := fmt.Sprintf("ServeHTTP: %s %s %s\n", req.RemoteAddr, req.Method, req.RequestURI)

	_, _ = respWriter.Write([]byte(clientInfo))
}

func Test() {
	handle := &HttpHandle{}
	err := http.ListenAndServe(":8000", handle)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}
}
