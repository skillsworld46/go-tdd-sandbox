package blogposts_test

import (
	"testing"
	"testing/fstest"

	"github.com/skillsworld46/go-tdd-sandbox/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Hi")},
		"hello-world2.md": {Data: []byte("Hello")},
	}

	posts := blogposts.NewPostsFromFs(fs)

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, want %d posts", len(posts), len(fs))
	}
}
