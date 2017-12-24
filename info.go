package main

import (
	"html/template"
	"log"
	"net/http"
)

type infoModel struct {
	Template string
	Sidebar  sidebar
	Coincap  coincap
}

func info(w http.ResponseWriter, r *http.Request) {
	sb, err := sidebarValues()
	if err != nil {
		log.Print(err)
	}

	c, err := parseCoincap()
	if err != nil {
		log.Print(err)
	}

	model := infoModel{"info", sb, c}

	t, err := template.ParseFiles(
		"static/templates/layout.html",
		"static/templates/head.html",
		"static/templates/sidebar.html",
		"static/templates/info.html",
		"static/templates/info.js",
	)
	if err != nil {
		log.Print(err)
	}

	if err := t.Execute(w, model); err != nil {
		log.Print(err)
	}
}