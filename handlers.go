package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	store     = sessions.NewCookieStore([]byte("something-very-secret"))
	templates = template.Must(template.ParseGlob("templates/*.html"))
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Home page logic
	templates.ExecuteTemplate(w, "index.html", nil)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Registration logic
	if r.Method == http.MethodPost {
		// Handle registration
	}
	templates.ExecuteTemplate(w, "register.html", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Login logic
	if r.Method == http.MethodPost {
		// Handle login
	}
	templates.ExecuteTemplate(w, "login.html", nil)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Logout logic
	session, _ := store.Get(r, "session-name")
	session.Values["authenticated"] = false
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// Create new post logic
	if r.Method == http.MethodPost {
		// Handle new post
	}
	templates.ExecuteTemplate(w, "post.html", nil)
}

func ViewPostHandler(w http.ResponseWriter, r *http.Request) {
	// View post logic
	vars := mux.Vars(r)
	id := vars["id"]
	// Retrieve post by id
	templates.ExecuteTemplate(w, "post.html", id)
}

func EditPostHandler(w http.ResponseWriter, r *http.Request) {
	// Edit post logic
	if r.Method == http.MethodPost {
		// Handle post update
	}
	templates.ExecuteTemplate(w, "edit.html", nil)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	// Delete post logic
	vars := mux.Vars(r)
	id := vars["id"]
	// Delete post by id
	http.Redirect(w, r, "/", http.StatusFound)
}

func CommentHandler(w http.ResponseWriter, r *http.Request)
