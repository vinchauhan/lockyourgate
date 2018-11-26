package main

import (
	"fmt"
	"github.com/vinchauhan/lockyourgate/controllers"
	"github.com/vinchauhan/lockyourgate/views"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var homeView *views.View
var loginView *views.View

//var homeView = *views.View

func index(w http.ResponseWriter, r *http.Request) {
	if err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil); err != nil {
		panic(err)
	}
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("/css"))
}

func resourceNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "<h1>Sorry Not Found - Monkey Dispached</h1>")
}

func main() {

	f, err := os.OpenFile("lock-your-gate.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	log.SetOutput(f)
	userC := controllers.NewUser()
	homeView = views.NewView("semantic", "views/home.gohtml")

	r := mux.NewRouter()

	r.HandleFunc("/login", userC.New)
	r.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("./css/"))))
	r.PathPrefix("/images").Handler(http.StripPrefix("/images/", http.FileServer(http.Dir("./images/"))))
	r.HandleFunc("/", index)
	r.NotFoundHandler = http.HandlerFunc(resourceNotFound)
	log.Println("Server started on Port :", 3000)
	if err := http.ListenAndServe(":3000", r); err != nil {
		panic(err)
	}
	log.Println("Server started on Port :", 3000)
}
