//website_golang
package main

import (
	"html/template"
	"log"
	"net/http"
)

// Şablonları önceden yükleyin
var templates = template.Must(template.ParseGlob("templates/*.html"))

// Ana sayfa için handler
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler çağrıldı")
	renderTemplate(w, "index.html")
}


//Produced By K. Umut Araz
// Hakkında sayfası için handler
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("aboutHandler çağrıldı")
	renderTemplate(w, "about.html")
}

// Şablonları render eden yardımcı fonksiyon
func renderTemplate(w http.ResponseWriter, tmpl string) {
	log.Printf("%s şablonu render ediliyor", tmpl)
	err := templates.ExecuteTemplate(w, tmpl, nil)
	if err != nil {
		log.Printf("Şablon render hatası: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	// Yönlendirme
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)

	// Sunucuyu başlat
	log.Println("Sunucu http://localhost:8080 adresinde çalışıyor...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
