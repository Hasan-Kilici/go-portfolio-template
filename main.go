package main

import (
	"fmt"
	"net/http"
	"html/template"
	"time"
	"path"
)

func handle(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseGlob("templates/*")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := ""
	if r.URL.Path == "/" {
		name = "index.html"
	} else if r.URL.Path == "/hakkimda"{
    name = "hakkimda.html"
    } else if r.URL.Path == "/sosyal-medya" {
    name = "sosyal-medya.html"
    } else {
		name = path.Base(r.URL.Path)
	}

	data := struct{
		Time time.Time
    Author string
    Hakkimda string
    Discord string
    Github string
    PP string
  }{
		Time: time.Now(),
    Author: "Sitenin sahibi",
    Hakkimda: `Lorem Ipsum
"Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit..."
"There is no one who loves pain itself, who seeks after it and wants to have it, simply because it is pain..."`,
  Discord: "https://discord.gg/MQFct2YGb6",
  Github: "https://github.com/hasan-kilici",
  PP: "/static/32.png",
	}

	if err := tmpl.ExecuteTemplate(w, name, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error", err)
	}
}

func main() {
	fmt.Println("http server up!")
	http.Handle(
		"/static/",
		 http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)
	http.HandleFunc("/", handle)
	http.ListenAndServe(":0", nil)
}
