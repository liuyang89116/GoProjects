package server

import (
	"Bookstore/server/middleware"
	"Bookstore/store"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type BookStoreServer struct {
	store  store.Store
	server *http.Server
}

func NewBookStoreServer(addr string, s store.Store) *BookStoreServer {
	server := &BookStoreServer{
		store: s,
		server: &http.Server{
			Addr: addr,
		},
	}

	router := mux.NewRouter()
	router.HandleFunc("/book", server.createBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", server.updateBookHandler).Methods("POST")
	router.HandleFunc("/book/{id}", server.getBookHandler).Methods("GET")
	router.HandleFunc("/book", server.getAllBookHandler).Methods("GET")
	router.HandleFunc("/book/{id}", server.delBookHandler).Methods("DELETE")

	server.server.Handler = middleware.Logging(middleware.Validating(router))

	return server
}

func (bss *BookStoreServer) ListenAndServe() (<-chan error, error) {
	var err error
	errChan := make(chan error)
	go func() {
		bss.server.ListenAndServe()
		errChan <- err
	}()

	select {
	case err = <-errChan:
		return nil, err
	case <-time.After(time.Second):
		return errChan, nil
	}
}

func (bss *BookStoreServer) Shutdown(ctx context.Context) error {
	return bss.server.Shutdown(ctx)
}

func (bss *BookStoreServer) createBookHandler(writer http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var book store.Book
	if err := decoder.Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if err := bss.store.Create(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (bss *BookStoreServer) updateBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "no id found in request", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(request.Body)
	var book store.Book
	if err := decoder.Decode(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	book.Id = id
	if err := bss.store.Update(&book); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func (bss *BookStoreServer) getBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "no id found in request", http.StatusBadRequest)
		return
	}

	book, err := bss.store.Get(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response(writer, book)
}

func (bss *BookStoreServer) getAllBookHandler(writer http.ResponseWriter, request *http.Request) {
	books, err := bss.store.GetAll()
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	response(writer, books)
}

func (bss *BookStoreServer) delBookHandler(writer http.ResponseWriter, request *http.Request) {
	id, ok := mux.Vars(request)["id"]
	if !ok {
		http.Error(writer, "no id found in request", http.StatusBadRequest)
		return
	}

	err := bss.store.Delete(id)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
}

func response(writer http.ResponseWriter, v interface{}) {
	data, err := json.Marshal(v)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(data)
}
