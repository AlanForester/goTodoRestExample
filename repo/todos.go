package repo

import (
	"database/sql"

	"github.com/AlexCollin/goTodoRestExample/model"

	_ "github.com/lib/pq"
)

type Todos struct {
	DB *sql.DB
}

func (p *Todos) Close() {
	p.DB.Close()
}

func (p *Todos) GetAll() ([]model.Todo, error) {
	query := `
        SELECT *
        FROM todos
        ORDER BY id;
    `

	rows, err := p.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var todoList []model.Todo
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.UserID); err != nil {
			return nil, err
		}
		todoList = append(todoList, t)
	}

	return todoList, nil
}

func (p *Todos) Insert(todo *model.Todo) (int, error) {
	query := `
        INSERT INTO todos (id, title, user_id)
        VALUES (nextval('todo_id'), $1, $2)
        RETURNING id;
    `

	rows, err := p.DB.Query(query, todo.Title, todo.UserID)
	if err != nil {
		return -1, err
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return -1, err
		}
	}

	return id, nil
}

func (p *Todos) Get(id int) (model.Todo, error) {
	query := `
        SELECT *
        FROM todos
        WHERE id = $1;
    `

	rows, err := p.DB.Query(query, id)
	if err != nil {
		return model.Todo{}, err
	}

	var todo model.Todo
	for rows.Next() {
		var t model.Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.UserID); err != nil {
			return model.Todo{}, err
		}
		todo = t
	}

	return todo, nil
}

func (p *Todos) Update(todo *model.Todo) error {
	query := `
        UPDATE todos SET title = $2, user_id = $3 
         WHERE id = $1 
        RETURNING id;
    `

	_, err := p.DB.Query(query, todo.ID, todo.Title, todo.UserID)
	if err != nil {
		return err
	}

	return nil
}

func (p *Todos) Delete(id int) error {
	query := `
        DELETE FROM todos
        WHERE id = $1;
    `

	if _, err := p.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}

func ConnectPostgres() (*Todos, error) {
	connStr := "postgres://postgres:postgres@db/todo?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Todos{db}, nil
}
