package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", index).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/contact", formcontact).Methods("POST")
	route.HandleFunc("/project", project).Methods("GET")
	route.HandleFunc("/project", formproject).Methods("POST")

	fmt.Println("server running on Port 2000")
	http.ListenAndServe("localhost:2000", route)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("index.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// name, _ := strconv.Atoi(mux.Vars(r)["name"])
	// fmt.Println(name)

	response := map[string]interface{}{
		"Title":   "Hi, Welcome to my hut!",
		"Id":      "id",
		"Content": "Lorem ipsum dolor sit, amet consectetur adipisicing elit. In libero quibusdam pariatur non modi temporibus quas fugiat numquam autem ullam? Amet dolore minima explicabo nemo nostrum enim dicta error et. Lorem, ipsum dolor sit amet consectetur adipisicing elit. Sed quidem hic et qui iste nostrum eligendi, eius nam! Aliquam, voluptate. Deserunt iste, voluptas est quod doloremque veniam at autem commodi.",
	}
	tmpl.Execute(w, response)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("contact.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, "")
}
func formcontact(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Title: " + r.PostForm.Get("input-name"))
	fmt.Println("Email: " + r.PostForm.Get("input-email"))
	fmt.Println("Phone: " + r.PostForm.Get("input-phone"))
	fmt.Println("Subject: " + r.PostForm.Get("input-subject"))
	fmt.Println("message: " + r.PostForm.Get("input-message"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	var tmpl, err = template.ParseFiles("project.html")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	tmpl.Execute(w, "")
}
func formproject(w http.ResponseWriter, r *http.Request) {
	error := r.ParseForm()
	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Title: " + r.PostForm.Get("input-title"))
	fmt.Println("Start: " + r.PostForm.Get("input-start"))
	fmt.Println("End: " + r.PostForm.Get("input-end"))
	fmt.Println("Content: " + r.PostForm.Get("input-content"))
	fmt.Println("File: " + r.PostForm.Get("input-file"))
	fmt.Println("Technology: " + r.PostForm.Get("Node-Js, Next-Js, TypeScript, React-Js"))

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
