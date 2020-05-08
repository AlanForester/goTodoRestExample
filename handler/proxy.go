package handler

import (
	"encoding/json"
	"github.com/AlexCollin/goTodoRestExample/model"
	"io/ioutil"
	"net/http"
)

type proxyHandler struct{}

func (handler *proxyHandler) check(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var proxy model.Proxy
	if err := json.Unmarshal(b, &proxy); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	responseOk(w, map[string]bool{"status": proxy.Check()})
}
