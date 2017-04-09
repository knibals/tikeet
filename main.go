package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var port = os.Getenv("PORT")

func main() {
	if port == "" {
		log.Fatal("La variable \"PORT\" ne peut pas être vide...")
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintln(w, "<h1>Aucun ticket disponible à la vente</h1>")
	})
	log.Printf("Démarrage du serveur sur le port %s\n", port)
	log.Fatal(http.ListenAndServe("127.0.0.1:"+port, nil))
}
