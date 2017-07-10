package posts

import (
	"fmt"
	"net/http"
)

type Post struct {
	UID      string
	Text     string
	Username string
	Avatar   string
	Favorite bool
}

// func that calls before main
func init() {
	http.HandleFunc("/posts", listPosts)
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world!")
}
