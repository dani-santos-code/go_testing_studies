package greet

import (
	"fmt"
	"io"
	"net/http"
)

// Greet will greet users
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// MyGreeterHandler will greet as a response
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func StartServer() {
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
