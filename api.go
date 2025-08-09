package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiServer struct {
	listenPort string
	store      Storage
}

func NewApiServer(listenPort string, store Storage) *ApiServer {
	return &ApiServer{
		listenPort: listenPort,
		store:      store,
	}
}
func (s *ApiServer) Run() {
	router := mux.NewRouter()

	accountRouter := router.PathPrefix("/accounts").Subrouter()
	accountRouter.HandleFunc("", makeHttpHandleFunc(s.handleCreateAccount)).Methods("POST")
	accountRouter.HandleFunc("", makeHttpHandleFunc(s.handleGetAccount)).Methods("GET")
	accountRouter.HandleFunc("/{id}", makeHttpHandleFunc(s.handleGetAccountById)).Methods("GET")
	accountRouter.HandleFunc("", makeHttpHandleFunc(s.handleDeleteAccount)).Methods("DELETE")

	// Catch-all route for invalid paths
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("⚠️ Invalid route accessed: %s %s", r.Method, r.URL.Path)
		WriteJson(w, http.StatusBadRequest, ApiError{Error: "Invalid route"})
	})

	log.Printf("🚀 Server is running on port %s", s.listenPort)
	http.ListenAndServe(s.listenPort, router)
}

func (s *ApiServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	createAccReq := new(CreateAccountRequest)
	if err := json.NewDecoder(r.Body).Decode(createAccReq); err != nil {
		return fmt.Errorf("bad request - invalid request body")
	}
	defer r.Body.Close()

	acc := NewAccount(createAccReq.FirstName, createAccReq.LastName)
	if err := s.store.CreateAccount(acc); err != nil {
		return err
	}
	return WriteJson(w, http.StatusCreated, acc)
}

func (s *ApiServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	accounts, err := s.store.GetAccount()
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, accounts)
}
func (s *ApiServer) handleGetAccountById(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return fmt.Errorf("bad request - unparseable id")
	}

	acc, err := s.store.GetAccountById(id)
	if err != nil {
		return err
	}

	return WriteJson(w, http.StatusOK, acc)
}
func (s *ApiServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("DELETE method called on /accounts endpoint")
	return nil
}
