package repo

import (
	"context"

	"github.com/AlexCollin/goTodoRestExample/model"
)

const keyRepository = "Repository"

type Repository interface {
	Close()
	GetAll() ([]model.Todo, error)
	Insert(todo *model.Todo) (int, error)
	Get(id int) (model.Todo, error)
	Update(todo *model.Todo) error
	Delete(id int) error
}

func SetRepository(ctx context.Context, repository Repository) context.Context {
	return context.WithValue(ctx, keyRepository, repository)
}

func Close(ctx context.Context) {
	getRepository(ctx).Close()
}

func GetAll(ctx context.Context) ([]model.Todo, error) {
	return getRepository(ctx).GetAll()
}

func Insert(ctx context.Context, todo *model.Todo) (int, error) {
	return getRepository(ctx).Insert(todo)
}

func Get(ctx context.Context, id int) (model.Todo, error) {
	return getRepository(ctx).Get(id)
}

func Update(ctx context.Context, todo *model.Todo) error {
	return getRepository(ctx).Update(todo)
}

func Delete(ctx context.Context, id int) error {
	return getRepository(ctx).Delete(id)
}

func getRepository(ctx context.Context) Repository {
	return ctx.Value(keyRepository).(Repository)
}
