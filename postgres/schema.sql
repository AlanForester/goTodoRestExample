DROP TABLE IF EXISTS todos;
CREATE SEQUENCE todo_id START 1;
CREATE TABLE todos (
  ID serial PRIMARY KEY,
  TITLE TEXT NOT NULL,
  USER_ID INTEGER
);