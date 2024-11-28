package main

import (
	"fmt"
)

type author struct {
	firstName	string
	lastName	string
	bio	string
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type blogPost struct {
	title	string
	content	string
	author
}

func (b blogPost) details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.author.fullName())
	fmt.Println("Bio: ", b.author.bio)
}

func main() {
	author1 := author{
		"Naveen",
		"Ramanthan",
		"Golang Enthusiast",
	}
	
	blogPost1 := blogPost{
		"Inheritence in Go",
		"Go supports composition instead of inhertience",
		author1,
	}
	
	blogPost2 := blogPost{
		"Struct instead of Classes in Go",
		"Go does not support classes but methods can be added to structs",
		author1,
	}
	
	blogPost3 := blogPost{
		"Concurrency",
		"Go is a concurrent languagre and not a parallel one",
		author1,
	}
	
	w := website{
		blogPosts: []blogPost{blogPost1, blogPost2, blogPost3},
	}
	w.contents()
	
}

type website struct {
	blogPosts []blogPost
}

func (w website) contents() {
	fmt.Println("Contents of Website\n")
	for _, v := range w.blogPosts {
		v.details()
		fmt.Println()
	}
}