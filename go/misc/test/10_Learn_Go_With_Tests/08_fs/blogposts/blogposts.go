package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}

	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, f fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(f.Name())
	if err != nil {
		return Post{}, err
	}

	defer postFile.Close()

	return newPost(postFile)
}

const (
	TitleName       = "Title: "
	DescriptionName = "Description: "
	TagsName        = "Tags: "
)

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	title := readLine(scanner, TitleName)
	description := readLine(scanner, DescriptionName)
	tags := readTags(scanner)
	body := readBody(scanner)

	post := Post{
		Title:       title,
		Description: description,
		Tags:        tags,
		Body:        body,
	}
	return post, nil
}

func readLine(scanner *bufio.Scanner, tagName string) string {
	scanner.Scan()
	text := scanner.Text()
	return strings.TrimPrefix(text, tagName)
}

func readTags(scanner *bufio.Scanner) []string {
	tags := readLine(scanner, TagsName)
	return strings.Split(tags, ", ")
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan()
	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
