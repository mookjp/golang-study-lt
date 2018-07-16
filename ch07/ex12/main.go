package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/create", db.create)
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	page, err := loadPage("list.html")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	// render
	listTemplate := template.New("list")
	template.Must(listTemplate.Parse(string(page)))
	listTemplate.Execute(w, db)
}

func loadPage(pageName string) ([]byte, error) {
	body, err := ioutil.ReadFile(fmt.Sprintf("./%v", pageName))
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
	}
}

// Create
func (db *database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := (*db)[item]; ok {
		w.WriteHeader(http.StatusConflict) // 409
		fmt.Fprintf(w, "conflict; item: %q\n", item)
		return
	} else {
		parsed, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprintf(w, "price should be a form to be converted to float")
			return
		}
		(*db)[item] = dollars(parsed)
		w.WriteHeader(http.StatusCreated) // 201
		fmt.Fprintf(w, "created; item: %q\n", item)
	}
}

// Update
func (db *database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if _, ok := (*db)[item]; ok {
		parsed, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // 400
			fmt.Fprintf(w, "price should be a form to be converted to float")
			return
		}
		(*db)[item] = dollars(parsed)
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "updated; item: %q\n", item)
	} else {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "no such item. item: %q\n", item)
		return
	}
}

// Delete
func (db *database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if _, ok := (*db)[item]; ok {
		delete(*db, item)
		w.WriteHeader(http.StatusOK) // 200
		fmt.Fprintf(w, "delete; item: %q\n", item)
		return
	} else {
		w.WriteHeader(http.StatusBadRequest) // 400
		fmt.Fprintf(w, "no such item. item: %q\n", item)
		return
	}
}
