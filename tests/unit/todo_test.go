package unit

import (
	"github.com/AlexCollin/goTodoRestExample/repo"
	"github.com/AlexCollin/goTodoRestExample/tests"
	"reflect"
	"testing"

	"github.com/AlexCollin/goTodoRestExample/model"
	_ "github.com/lib/pq"
)

func TestPostgres_GetAll(t *testing.T) {
	todos := &repo.Todos{tests.Setup()}
	defer todos.Close()

	todo := &model.Todo{
		Title: "title1",
		Token: "123",
	}

	_, err := todos.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	got, err := todos.GetAll()
	if err != nil {
		t.Fatal(err)
	}

	want := []*model.Todo{
		{
			ID:    1,
			Title: "title1",
			Token: "123",
		},
	}

	if equal(got, want) {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
}

func TestPostgres_Insert(t *testing.T) {
	todos := &repo.Todos{tests.Setup()}
	defer todos.Close()

	todo := &model.Todo{
		Title: "title1",
		Token: "123",
	}

	got, err := todos.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	want := 1

	if got != want {
		t.Fatal(err)
	}
}

func TestPostgres_Get(t *testing.T) {
	todos := &repo.Todos{tests.Setup()}
	defer todos.Close()

	todo := &model.Todo{
		Title: "title1",
		Token: "123",
	}

	id, err := todos.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	got, err := todos.Get(id)
	if err != nil {
		t.Fatal(err)
	}

	want := &model.Todo{
		ID:    1,
		Title: "title1",
		Token: "123",
	}

	if equal(got, want) {
		t.Fatalf("Want: %v, Got: %v", want, got)
	}
}

func TestPostgres_Update(t *testing.T) {
	todos := &repo.Todos{tests.Setup()}
	defer todos.Close()

	todo := &model.Todo{
		Title: "title1",
		Token: "123",
	}

	_, err := todos.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	err = todos.Update(todo)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPostgres_Delete(t *testing.T) {
	todos := &repo.Todos{tests.Setup()}
	defer todos.Close()

	todo := &model.Todo{
		Title: "title1",
		Token: "123",
	}

	id, err := todos.Insert(todo)
	if err != nil {
		t.Fatal(err)
	}

	err = todos.Delete(id)
	if err != nil {
		t.Fatal(err)
	}

	_, err = todos.Get(id)
	if err != nil {
		t.Fatal(err)
	}

}

func equal(got interface{}, want interface{}) bool {
	return reflect.DeepEqual(got, want)
}
