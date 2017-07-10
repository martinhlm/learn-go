package posts

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	UID      string
	Text     string
	Username string
	Avatar   string
	Favorite bool
}

var posts = []Post{
	{
		"1",
		"Go is awesome",
		"Gopher",
		"http://www.evanmiller.org/images/go-gopher3.png",
		true,
	},
	{
		"2",
		"Go is awesome",
		"Gopher",
		"http://www.evanmiller.org/images/go-gopher3.png",
		true,
	},
}

// func that calls before main
func init() {
	http.HandleFunc("/posts", listPosts)
}

func listPosts(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	enc.Encode(posts)
}
