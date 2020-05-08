package handler

import (
	"github.com/AlexCollin/goTodoRestExample/repo"
	"net/http"
)

func SetUpRouting(repo *repo.Todos) *http.ServeMux {
	todoHandler := &todoHandler{
		repo: repo,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/todo/", AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.getTodo(w, r)
		case http.MethodPut:
			todoHandler.updateTodo(w, r)
		case http.MethodDelete:
			todoHandler.deleteTodo(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	}))
	mux.HandleFunc("/todo", AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoHandler.getAllTodo(w, r)
		case http.MethodPost:
			todoHandler.insertTodo(w, r)
		default:
			responseError(w, http.StatusNotFound, "")
		}
	}))

	return mux
}
