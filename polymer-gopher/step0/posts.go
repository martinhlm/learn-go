package posts

import (
	"fmt"
	"net/http"
)

// func that calls before main
func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
