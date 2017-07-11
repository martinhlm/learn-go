package posts

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"

	"golang.org/x/net/context"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

// define all the endpoints of the posts API
type PostAPI struct{}

type Post struct {
	UID      *datastore.Key
	Text     string
	Username string
	Avatar   string
	Favorite bool
}

type Posts struct {
	Posts []Post
}

type AddRequest struct {
	Text     string
	Username string
	Avatar   string
}

// func that calls before main
func init() {
	endpoints.HandleHTTP()
}

// method
func (PostAPI) list(w http.ResponseWriter, r *http.Request) {
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

func (PostAPI) Add(c context.Context, r *AddRequest) (*Post, error) {
	p := Post{
		Text:     r.Text,
		Username: r.Username,
		Avatar:   r.Avatar,
	}

	key := datastore.NewIncompleteKey(c, "Post", nil)
	key, err := datastore.Put(c, key, &p)
	if err != nil {
		return nil, fmt.Errorf("put post: %v", err)
	}
	p.UID = key

	return &p, nil
}
