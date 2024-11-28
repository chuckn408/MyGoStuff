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
	blogPost1.details()
}