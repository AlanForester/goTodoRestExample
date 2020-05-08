package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AlexCollin/goTodoRestExample/handler"
	"github.com/AlexCollin/goTodoRestExample/repo"
)

func main() {
	var postgres *repo.Todos
	var err error

	postgres, err = repo.ConnectPostgres()

	if err != nil {
		panic(err)
	} else if postgres == nil {
		panic("postgres is nil")
	}

	mux := handler.SetUpRouting(postgres)

	fmt.Println("http://localhost:8088")
	log.Fatal(http.ListenAndServe(":8088", mux))
}
