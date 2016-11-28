package main

import (
    "html/template"
    "net/http"
    "path"
)

var ROOT_DIR = path.Dir(".")
var TEMPLATES_DIR = path.Join(ROOT_DIR, "templates")

type Context struct {
    Title string
}

func init() {
    http.HandleFunc("/", homeHandler)

    // Mandatory root-based resources
    serveSingle("/favicon.ico", "images/favicon.ico")

    // Normal resources
    http.Handle("/static", http.FileServer(http.Dir("./static/")))
}

func main() {
    http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fp := path.Join(TEMPLATES_DIR, "home.html")
    tmpl, err := template.ParseFiles(fp)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    context := Context{Title:"Krystian Hanek website"}
    if err := tmpl.Execute(w, context); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func serveSingle(pattern string, filename string) {
    http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, path.Join(TEMPLATES_DIR, filename))
    })
}
