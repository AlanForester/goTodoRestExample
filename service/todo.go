package service

import (
	"context"

	"github.com/AlexCollin/goTodoRestExample/model"
	"github.com/AlexCollin/goTodoRestExample/repo"
)

func Close(ctx context.Context) {
	repo.Close(ctx)
}

func GetAll(ctx context.Context) ([]model.Todo, error) {
	return repo.GetAll(ctx)
}

func Insert(ctx context.Context, todo *model.Todo) (int, error) {
	return repo.Insert(ctx, todo)
}

func Get(ctx context.Context, id int) (model.Todo, error) {
	return repo.Get(ctx, id)
}

func Update(ctx context.Context, todo *model.Todo) error {
	return repo.Update(ctx, todo)
}

func Delete(ctx context.Context, id int) error {
	return repo.Delete(ctx, id)
}
