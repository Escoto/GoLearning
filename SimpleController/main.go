package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

type mypage struct {
	Title string
	Body  template.HTML
}

// Loads a page for use
func loadPage(title string, r *http.Request) (*mypage, error) {

	x := fmt.Sprintf("./web/%s.html", title)
	body, err := ioutil.ReadFile(x)
	if err != nil {
		return nil, err
	}

	return &mypage{Title: title, Body: template.HTML(body)}, nil
}

// dynamically load webpage
func viewHandler(w http.ResponseWriter, r *http.Request) {

	title := r.URL.Path[len("/"):]
	p, err := loadPage(title, r)

	if err != nil && title != "home" {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	t, _ := template.ParseFiles("./views/view.html")
	t.Execute(w, p)
}

// Shows a particular page
func testHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./views/testview.html")
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/test", testHandler)
	http.ListenAndServe(":8080", nil)
}
