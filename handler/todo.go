package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/AlexCollin/goTodoRestExample/model"
	"github.com/AlexCollin/goTodoRestExample/repo"
	"github.com/AlexCollin/goTodoRestExample/service"
)

type todoHandler struct {
	repo *repo.Todos
}

func (handler *todoHandler) getAllTodo(w http.ResponseWriter, r *http.Request) {
	ctx := repo.SetRepository(r.Context(), handler.repo)

	todoList, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, todoList)
}

func (handler *todoHandler) insertTodo(w http.ResponseWriter, r *http.Request) {
	ctx := repo.SetRepository(r.Context(), handler.repo)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todo model.Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := service.Insert(ctx, &todo)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, map[string]int{"id": id})
}

func (handler *todoHandler) getTodo(w http.ResponseWriter, r *http.Request) {
	ctx := repo.SetRepository(r.Context(), handler.repo)

	id, err := getTodoIdFromPath(r)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	todo, err := service.Get(ctx, id)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, todo)
}

func (handler *todoHandler) updateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := repo.SetRepository(r.Context(), handler.repo)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var todo model.Todo
	if err := json.Unmarshal(b, &todo); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	todo.ID, err = getTodoIdFromPath(r)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = service.Update(ctx, &todo)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOk(w, todo)
}

func (handler *todoHandler) deleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := repo.SetRepository(r.Context(), handler.repo)

	id, err := getTodoIdFromPath(r)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := service.Delete(ctx, id); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func getTodoIdFromPath(r *http.Request) (int, error) {
	pathSplit := strings.Split(r.URL.Path, "/")
	if len(pathSplit) < 3 {
		return 0, errors.New("Url Param '/todo/{id}' is missing ")
	}

	id, err := strconv.Atoi(pathSplit[2])
	if err != nil {
		return 0, errors.New("Url Param '/todo/{id}' must be integer ")
	}

	return id, nil
}

func responseOk(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
