package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title:filename, Body:body}, nil
}


;; Web server
const lenPath = len("/view/")

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, _ loadPage(title)
	fmt.Fprint(w, "<h1>%s</h1><div>%s</div>", p.Title, (string)p.Body)
}

func main() {
	p1 := &Page{Title: "MattPage", Body: []byte("This is my go program")}
 	p1.save()
	
	p2, _ := loadPage("MattPage")	
	fmt.Println(string(p2.Body))
}
