package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Route handlers
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/register", RegisterHandler).Methods("GET", "POST")
	router.HandleFunc("/login", LoginHandler).Methods("GET", "POST")
	router.HandleFunc("/logout", LogoutHandler).Methods("GET")
	router.HandleFunc("/post/new", CreatePostHandler).Methods("GET", "POST")
	router.HandleFunc("/post/{id}", ViewPostHandler).Methods("GET")
	router.HandleFunc("/post/{id}/edit", EditPostHandler).Methods("GET", "POST")
	router.HandleFunc("/post/{id}/delete", DeletePostHandler).Methods("POST")
	router.HandleFunc("/post/{id}/comment", CommentHandler).Methods("POST")

	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
