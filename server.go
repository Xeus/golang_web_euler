/**
 * learning golang by implementing net/http, project euler, and asserts/testing
 *
 * see:
 * http://golang.org/doc/articles/wiki/
 * http://golangtutorials.blogspot.com/2011/10/gotest-unit-testing-and-benchmarking-go.html
 * https://github.com/coocood/assrt
 * https://github.com/eddie/goalg
 * https://github.com/JanLaussmann/Project-Euler-Golang
 */

package main

import (
 	"fmt"
 	"strconv"
	"html/template"
	"io/ioutil"
	"net/http"
	"regexp"
	"encoding/json"
	"./euler"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// }

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
    if err != nil {
        http.Redirect(w, r, "/edit/" + title, http.StatusFound)
        return
    }
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
    body := r.FormValue("body")
    p := &Page{Title: title, Body: []byte(body)}
    err := p.save()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    http.Redirect(w, r, "/view/" + title, http.StatusFound)
}

var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl + ".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			fn(w, r, "FrontPage")
		} else {
			m := validPath.FindStringSubmatch(r.URL.Path)
	        if m == nil {
	            http.NotFound(w, r)
	            return
	        }
	        fn(w, r, m[2])
    	}
	}
}

type Player struct {  // NBA player
	Name string
	Team string
	Position string
}

// see http://nesv.blogspot.com/2012/09/super-easy-json-http-responses-in-go.html
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	dirk := Player{"Dirk Nowitzki", "Dallas Mavericks", "F"}
	vince := Player{"Vince Carter", "Dallas Mavericks", "F"}
	mavericks := []Player{dirk, vince}  // slice of player structs
	mavs, err := json.Marshal(mavericks)  // convert slice to json
	if err != nil {
		http.NotFound(w, r)
	    return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(mavs))
}

var eulerPath = regexp.MustCompile("^/euler/([0-9]+)[/]*([0-9]*)[/]*$")

type EulerResult struct {
	ProblemNum int
	Result int
}

func eulerHandler(w http.ResponseWriter, r *http.Request) {
	result := 0

	m := eulerPath.FindStringSubmatch(r.URL.Path)
    if m == nil {
        http.NotFound(w, r)
        return
    }

	problemNum, err := strconv.Atoi(m[1])
	if err != nil {
		http.NotFound(w, r)
    	return
	}

	maxNum := 10
	if m[2] != "" {
		userMaxNum, err := strconv.Atoi(m[2])
		if err != nil {
			http.NotFound(w, r)
	    	return
		}
		maxNum = userMaxNum
	} else {
		maxNum = 10
	}

	switch problemNum {
	case 1:
		result = euler.Problem1(maxNum)
		break
	default:
		result = euler.Problem1(10)
	}
	eulerResult := EulerResult{problemNum, result}
	
	answer, err := json.Marshal(eulerResult)  // convert slice to json
	if err != nil {
		http.NotFound(w, r)
	    return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(answer))
}

func main() {
	http.HandleFunc("/", makeHandler(viewHandler))
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/mavericks/", jsonHandler)  // trying out json
	http.HandleFunc("/euler/", eulerHandler)
	http.ListenAndServe(":8080", nil)
}
