package posts

import (
	"encoding/json"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
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
	c := appengine.NewContext(r)

	posts := []Post{}
	_, err := datastore.NewQuery("Post").GetAll(c, &posts)
	if err != nil {
		//c.Errorf("fetching posts: %v", err)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(posts)
	if err != nil {
		log.Printf("encoding: %v", err)
	}
}
