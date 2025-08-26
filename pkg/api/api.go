package api

import (
	"comments-service/pkg/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// Программный интерфейс сервера
type API struct {
	db     storage.Interface
	router *mux.Router
}

// Конструктор объекта API
func New(db storage.Interface) *API {
	api := API{
		db: db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

// Регистрация обработчиков API
func (a *API) endpoints() {
	a.router.Use(requestIdMiddleware)
	// получить комментарии по n новости
	a.router.HandleFunc("/{n}", a.commentsHandler).Methods(http.MethodGet, http.MethodOptions)
	// добавить комментарий к новости или комментарию
	a.router.HandleFunc("/addComment", a.addCommentHandler).Methods(http.MethodPost, http.MethodOptions)
	a.router.Use(loggingMiddleware)
}

// Получение маршрутизатора запросов
func (a *API) Router() *mux.Router {
	return a.router
}

// Получение всех комментариев к новости
func (a *API) commentsHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["n"]
	n, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	comments, err := a.db.Comments(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

// Добавление комментария к новости или комментарию
func (a *API) addCommentHandler(w http.ResponseWriter, r *http.Request) {
	var p storage.Comment
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = a.db.AddComment(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
