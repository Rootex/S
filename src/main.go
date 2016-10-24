package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	fmt.Printf("Writing %s to file", p.Title)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	page1 := &Page{Title: "Testpage", Body: []byte("Data sample page")}
	page1.save()
	load1, _ := loadPage("Testpage")
	fmt.Println(string(load1.Body))
}
