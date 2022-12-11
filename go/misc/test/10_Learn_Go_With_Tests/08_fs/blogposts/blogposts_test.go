package blogposts_test

import (
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "myapp/blogposts"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md": {Data: []byte(`Title: Post 1
Description: Description 1
Tags: tdd, go
---
Hello
World`)},
		"hello-world2.md": {Data: []byte(`Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
B
L
M`)},
	}

	got, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(got) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(got), len(fs))
	}

	want := blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tags:        []string{"tdd", "go"},
		Body: `Hello
World`,
	}

	assertPost(t, got, want)
}

func assertPost(t testing.TB, got []blogposts.Post, want blogposts.Post) {
	t.Helper()
	for _, post := range got {
		if reflect.DeepEqual(post, want) {
			return
		}
	}
	t.Errorf("got %+v, want %+v", got, want)
}
