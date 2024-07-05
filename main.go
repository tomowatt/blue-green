package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var release string

type BlueGreenData struct {
	Title   string
	Text    string
	Color   string
	Release string
}

func InitData() *BlueGreenData {
	title := os.Getenv("TITLE")
	text := os.Getenv("TEXT")
	color := os.Getenv("COLOR")

	if title == "" {
		fmt.Println("\"TITLE\" not set, defaulting to: \"blue\"")
		title = "blue"
	}

	if text == "" {
		fmt.Println("\"TEXT\" not set, defaulting to: \"blue\"")
		text = "blue"
	}

	if color == "" {
		fmt.Println("\"COLOR\" not set, defaulting to: \"blue\"")
		color = "blue"
	}

	data := BlueGreenData{
		Title:   title,
		Text:    text,
		Color:   color,
		Release: release,
	}

	return &data
}

func healthz(w http.ResponseWriter, r *http.Request) {
	response := "Healthy"
	w.WriteHeader(200)
	w.Write([]byte(response))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("\"PORT\" not set, defaulting to: \"8888\"")
		port = "8888"
	}

	default_tmpl := os.Getenv("TEMPLATE")
	if default_tmpl == "" {
		fmt.Println("\"TEMPLATE\" not set, defaulting to: \"template/index.html\"")
		default_tmpl = "template/index.html"
	}

	data := InitData()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles(default_tmpl)
		if err != nil {
			fmt.Println(err)
			return
		}

		tmpl.Execute(w, data)
	})

	http.HandleFunc("/healthz", healthz)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
