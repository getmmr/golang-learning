package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/getmmr/golang-learning/httpserver"
)

func main() {
	handle := httpserver.HttpHandle{}
	err := http.ListenAndServe(":8000", &handle)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

}
