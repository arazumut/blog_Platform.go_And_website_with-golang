package main

type User struct {
	ID       int
	Username string
	Password string
}

type Post struct {
	ID      int
	Title   string
	Content string
	Author  User
}

type Comment struct {
	ID      int
	Content string
	Author  User
	PostID  int
}
