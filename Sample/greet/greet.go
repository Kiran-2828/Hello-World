package greet

import (
	"fmt"
	"net/http"
)

func Helloworld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}
