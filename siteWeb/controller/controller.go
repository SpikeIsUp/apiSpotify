package controller

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)


func RenderTemplate(w http.ResponseWriter, tmpl string) error {
    path := filepath.Join("template", tmpl)
    log.Println("Rendu:", path)

    t, err := template.ParseFiles(path)
    if err != nil {
        log.Println("Erreur ParseFiles:", err)
        http.Error(w, "Erreur serveur (template introuvable)", http.StatusInternalServerError)
        return err
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")

    if err := t.Execute(w, nil); err != nil {
        log.Println("Erreur Execute:", err)
        http.Error(w, "Erreur serveur (render)", http.StatusInternalServerError)
        return err
    }

    return nil
}


func Home(w http.ResponseWriter, r *http.Request) {
    log.Println("➡️ GET /")
    if err := RenderTemplate(w, "home.html"); err != nil {
        log.Println("Erreur RenderTemplate:", err)
    }
}