package httpserver

import (
	"fmt"
	"net/http"
	"os"
)

func Test2() {
	serverMux := http.NewServeMux()
	serverMux.HandleFunc("/hello", HandleHello)

	serverMux.HandleFunc("/post", HandlePost)

	err := http.ListenAndServe(":8000", serverMux)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error:%v\n", err)
	}
}

func HandleHello(writer http.ResponseWriter, req *http.Request) {
	clientInfo := fmt.Sprintf("HandleHello: %s %s %s\n", req.RemoteAddr, req.Method, req.RequestURI)
	fmt.Printf(clientInfo)
	_, _ = fmt.Fprintf(writer, clientInfo)
}

func HandlePost(writer http.ResponseWriter, req *http.Request) {
	clientInfo := fmt.Sprintf("HandlePost: %s %s %s \n", req.RemoteAddr, req.Method, req.RequestURI)
	fmt.Printf(clientInfo)

	err := req.ParseForm()
	if err != nil {
		http.Error(writer, fmt.Sprintf("error: %v\n", err), http.StatusInternalServerError)
		return
	}
	formFields := req.PostForm
	fmt.Printf("Form fields: %v\n", formFields)
	writer.WriteHeader(http.StatusOK)
	_, _ = writer.Write([]byte("提交成功\n"))
}
