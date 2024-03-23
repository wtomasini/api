package main

import (
	"github.com/willTomasini/api/pkg/recipes"
	"github.com/willTomasini/api/pkg/users"
	"net/http"

	"github.com/gorilla/mux"
)

var ()

func main() {
	recipeStore := recipes.NewMemStore()
	recipesHandler := NewRecipesHandler(recipeStore)

	userStore := users.NewMemStore()
	usersHandler := NewUsersHandler(userStore)

	home := homeHandler{}

	router := mux.NewRouter()

	router.HandleFunc("/", home.ServeHTTP)

	r := router.PathPrefix("/recipes").Subrouter()

	r.HandleFunc("/", recipesHandler.ListRecipes).Methods(http.MethodGet)
	r.HandleFunc("/", recipesHandler.CreateRecipe).Methods(http.MethodPost)
	r.HandleFunc("/{id}", recipesHandler.GetRecipe).Methods(http.MethodGet)
	r.HandleFunc("/{id}", recipesHandler.UpdateRecipe).Methods(http.MethodPut)
	r.HandleFunc("/{id}", recipesHandler.DeleteRecipe).Methods(http.MethodDelete)

	u := router.PathPrefix("/users").Subrouter()

	u.HandleFunc("/", usersHandler.ListUsers).Methods(http.MethodGet)
	u.HandleFunc("/", usersHandler.CreateUser).Methods(http.MethodPost)
	u.HandleFunc("/{id}", usersHandler.GetUser).Methods(http.MethodGet)
	u.HandleFunc("/{id}", usersHandler.UpdateUser).Methods(http.MethodPut)
	u.HandleFunc("/{id}", usersHandler.DeleteUser).Methods(http.MethodDelete)

	http.ListenAndServe(":8010", router)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Why tho\n -Ryan"))
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
