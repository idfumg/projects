package main

import (
	"fmt"
	"os"

	"myapp/blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("./blogposts"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Files were read: %d\n", len(posts))
}
