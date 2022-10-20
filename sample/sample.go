package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
)

// github.com/gorilla/mux adalah framework untuk menjalankan routing
func main() {

	// fmt.Println("hallo") //memiliki fungsi seperti console.log/ menampilkan value
	route := mux.NewRouter()

	//Routing folder public yang berisikan css
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	//Routing(Query Param)
	//var.HandleFunc()
	route.HandleFunc("/", index).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/contact", formcontact).Methods("POST")
	route.HandleFunc("/project", project).Methods("GET")
	route.HandleFunc("/project", formproject).Methods("POST")
	// w http.ResponseWriter = menampilkan di browser
	// r *http.Request = mengambil data
	//byte = menampilkan string

	fmt.Println("server Running on Port 2000")
	http.ListenAndServe("localhost:2000", route) //memberikan port localhost
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8 ") //sebagai struktur bagian head
	var tmpl, error = template.ParseFiles("index.html")         //untuk mengambil file html
	//tmpl, error = untuk menampung template dan erornya
	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(error.Error()))
		return
	}

	//untuk mengambil index
	name, _ := strconv.Atoi(mux.Vars(r)["name"])
	fmt.Println(name)

	//array string, map[string] = objectnya
	response := map[string]interface{}{
		"Title":   "Hi,",
		"Id":      "id",
		"Content": "Lorem ipsum dolor sit, amet consectetur adipisicing elit. In libero quibusdam pariatur non modi temporibus quas fugiat numquam autem ullam? Amet dolore minima explicabo nemo nostrum enim dicta error et. Lorem, ipsum dolor sit amet consectetur adipisicing elit. Sed quidem hic et qui iste nostrum eligendi, eius nam! Aliquam, voluptate. Deserunt iste, voluptas est quod doloremque veniam at autem commodi.",
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, response)
	// w.Write([]byte("Hello World"))
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("contact.html")

	if error != nil {
		w.Write([]byte(error.Error()))
		return
	} //jika eror tidak = nil maka tampilkan text eror
	//untuk menampilkan var yang mengandung file html
	tmpl.Execute(w, error)
	//w.Write([]byte("Hello Gaess"))
}
func formcontact(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	//untuk mengambil data dari html
	fmt.Println("Title: " + r.PostForm.Get("input-name"))      //Get berdasarkan dari name input
	fmt.Println("Email: " + r.PostForm.Get("input-email"))     //Get berdasarkan dari name input
	fmt.Println("Phone: " + r.PostForm.Get("input-phone"))     //Get berdasarkan dari name input
	fmt.Println("Subject: " + r.PostForm.Get("input-subject")) //Get berdasarkan dari name input
	fmt.Println("message: " + r.PostForm.Get("input-message")) //Get berdasarkan dari name input

	//untuk mengarahkan ketika sudah di submit
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func project(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	var tmpl, error = template.ParseFiles("project.html")

	if error != nil {
		w.Write([]byte(error.Error()))
		return
	}

	tmpl.Execute(w, nil)
	//w.Write([]byte("My Project"))
}
func formproject(w http.ResponseWriter, r *http.Request) {
	error := r.ParseForm()
	if error != nil {
		log.Fatal(error)
	}

	//untuk mengambil data dari input html
	fmt.Println("Title: " + r.PostForm.Get("input-title"))                                 //Get berdasarkan name input
	fmt.Println("Start: " + r.PostForm.Get("input-start"))                                 //Get berdasarkan name input
	fmt.Println("End: " + r.PostForm.Get("input-end"))                                     //Get berdasarkan name input
	fmt.Println("Content: " + r.PostForm.Get("input-content"))                             //Get berdasarkan name input
	fmt.Println("File: " + r.PostForm.Get("input-file"))                                   //Get berdasarkan name input
	fmt.Println("Technology: " + r.PostForm.Get("Node-Js, Next-Js, TypeScript, React-Js")) //Get berdasarkan name input

	//untuk mengarahkan ketika sudah di submit
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
