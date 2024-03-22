package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gosimple/slug"
	"github.com/willTomasini/api/pkg/recipes"
	"net/http"

	"github.com/gorilla/mux"
)

var ()

func main() {
	store := recipes.NewMemStore()
	recipesHandler := NewRecipesHandler(store)
	home := homeHandler{}

	router := mux.NewRouter()

	router.HandleFunc("/", home.ServeHTTP)

	r := router.PathPrefix("/recipes").Subrouter()

	r.HandleFunc("/", recipesHandler.ListRecipes).Methods(http.MethodGet)
	r.HandleFunc("/", recipesHandler.CreateRecipe).Methods(http.MethodPost)
	r.HandleFunc("/{id}", recipesHandler.GetRecipe).Methods(http.MethodGet)
	r.HandleFunc("/{id}", recipesHandler.UpdateRecipe).Methods(http.MethodPut)
	r.HandleFunc("/{id}", recipesHandler.DeleteRecipe).Methods(http.MethodDelete)

	//u := router.PathPrefix("/users").Subrouter()

	http.ListenAndServe(":8010", router)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Why tho\n -Ryan"))
}

type recipeStore interface {
	Add(name string, recipe recipes.Recipe) error
	Get(name string) (recipes.Recipe, error)
	List() (map[string]recipes.Recipe, error)
	Update(name string, recipe recipes.Recipe) error
	Remove(name string) error
}

type RecipesHandler struct {
	store recipeStore
}

func NewRecipesHandler(s recipeStore) *RecipesHandler {
	return &RecipesHandler{
		store: s,
	}
}

func (h RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe recipes.Recipe

	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	resourceID := slug.Make(recipe.Name)

	if err := h.store.Add(resourceID, recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {
	recipeList, err := h.store.List()

	jsonBytes, err := json.Marshal(recipeList)

	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	recipe, err := h.store.Get(id)
	if err != nil {
		if errors.Is(err, recipes.ErrNotFound) {
			NotFoundHandler(w, r)
			return
		}
		InternalServerErrorHandler(w, r)
		return
	}
	jsonBytes, err := json.Marshal(recipe)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func (h RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var recipe recipes.Recipe
	if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	if err := h.store.Update(id, recipe); err != nil {
		if errors.Is(err, recipes.NotFoundErr) {
			NotFoundHandler(w, r)
			return
		}

		InternalServerErrorHandler(w, r)
		return
	}

	jsonBytes, err := json.Marshal(recipe)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Println("done with update")
	fmt.Println(string(jsonBytes))
	w.Write(jsonBytes)
}

func (h RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := h.store.Remove(id); err != nil {
		InternalServerErrorHandler(w, r)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 Internal Server Error"))
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 Not Found"))
}
