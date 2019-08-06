package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/bird", getBirdHandler).Methods("GET")
	r.HandleFunc("/bird", createBirdHandler).Methods("POST")
	staticFileDir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDir))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8001", r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	birds, err := store.GetBirds()

	birdListBytes, err := json.Marshal(birds)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	bird := Bird{}
	err := r.ParseForm()
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")
	err = store.CreateBird(&bird)
	if err != nil {
		fmt.Println(err)
	}
	http.Redirect(w, r, "/assets/", http.StatusFound)
}
