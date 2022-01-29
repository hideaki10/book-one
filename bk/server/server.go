package server

import (
	"bookstore/server/middleware"
	"bookstore/store"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type BookStoreServer struct {
	s   store.Store
	srv *http.Server
}

func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	srv := &BookStoreServer{
		s: s,
		srv: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/book", srv.createBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", srv.updateBookHandler).Methods("POST")
	router.HandleFunc("/book", srv.getBookHandler).Methods("GET")
	router.HandleFunc("/book", srv.getAllBookHandler).Methods("GET")
	router.HandleFunc("/book/{id}", srv.deleteBookHandler).Methods("DELETE")

	srv.srv.Handler = middleware.logging(middleware.Validating(router))

	return srv
}

func (bs *BookStoreServer) createBookHandler(w http.ResponseWriter, req *http.Request) {
	dec := json.NewDecoder(req.Body)
	var book store.Book
	if err := dec.Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := bs.s.Create(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (bs *BookStoreServer) updateBookHandler(w http.ResponseWriter, req *http.Request) {
	id, ok := mux.Vars(req)["id"]
	if !ok {
		http.Error(w, "no id found in request", http.StatusBadRequest)
		return
	}
	book, err := bs.s.Get(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response(w, book)
}
func (bs *BookStoreServer) getBookHandler(w http.ResponseWriter, req *http.Request) {

}
func (bs *BookStoreServer) getAllBookHandler(w http.ResponseWriter, req *http.Request) {

}
func (bs *BookStoreServer) deleteBookHandler(w http.ResponseWriter, req *http.Request) {

}
func response(w http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}
