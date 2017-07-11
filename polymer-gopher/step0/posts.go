package posts

import (
	"fmt"

	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"

	"golang.org/x/net/context"

	"google.golang.org/appengine/datastore"
)

// define all the endpoints of the posts API
type PostAPI struct{}

type Post struct {
	UID      *datastore.Key `json:"uid" datastore:"-"` // name in datastore => don't store this
	Text     string         `json:"text"`
	Username string         `json:"username"`
	Avatar   string         `json:"avatar"`
	Favorite bool           `json:"favorite"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}

type AddRequest struct {
	Text     string
	Username string
	Avatar   string
}

// func that calls before main
func init() {
	endpoints.RegisterService(PostAPI{}, "posts", "v1", "posts api", true)
	endpoints.HandleHTTP()
}

// method
func (PostAPI) List(c context.Context) (*Posts, error) {
	posts := []Post{}
	keys, err := datastore.NewQuery("Post").GetAll(c, &posts)
	if err != nil {
		return nil, err
	}

	for i, k := range keys {
		posts[i].UID = k
	}

	return &Posts{posts}, nil
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

type SetFavoriteRequest struct {
	UID      *datastore.Key
	Favorite bool
}

func (PostAPI) SetFavorite(c context.Context, r *SetFavoriteRequest) error {
	return datastore.RunInTransaction(c, func(c context.Context) error {
		var post Post
		err := datastore.Get(c, r.UID, &post)
		if err != nil {
			return fmt.Errorf("get post: %v", err)
		}

		post.Favorite = r.Favorite

		_, err = datastore.Put(c, r.UID, &post)
		if err != nil {
			return fmt.Errorf("update post: %v", err)
		}

		return nil
	}, nil)
}
