package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var tmpl = make(map[string]*template.Template)
var port = os.Getenv("PORT")

type Page struct {
	Title, Content string
}

func init() {
	tmpl["home"] = template.Must(template.ParseFiles("templates/index.html", "templates/layout.html"))
	tmpl["page"] = template.Must(template.ParseFiles("templates/page.html", "templates/layout.html"))
}

func main() {
	if port == "" {
		log.Fatal("La variable \"PORT\" ne peut pas être vide...")
	}
	http.HandleFunc("/", homepage)
	http.HandleFunc("/parseform", parseForm)
	log.Printf("Démarrage du serveur sur le port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homepage(w http.ResponseWriter, r *http.Request) {
	page := Page{
		Title:   "Tikeeet, votre partenaire events",
		Content: "Dynamique content from a database or another content repository...",
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := tmpl["home"].ExecuteTemplate(w, "layout", page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func parseForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("Erreur de traitement du formulaire: %v\n", err)
	}
	event := Event{
		about:       r.FormValue("event_name"),
		eventStatus: r.FormValue("event_status"),
		location:    r.FormValue("event_location"),
	}
	log.Printf("%v\n", event)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
