package tmpl

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title   string
	Heading string
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	title := "Home"
	heading := "Movies"
	p := Page{title, heading}

	t, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	title := "Contact"
	heading := "Contact"
	p := Page{title, heading}

	t, err := template.ParseFiles("./web/templates/contact.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, p)
	if err != nil {
		panic(err)
	}
}
