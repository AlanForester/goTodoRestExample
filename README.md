# goTodoRestExample

## First implementation - TODO
Clean implementation REST api for TODO application, with postgres

### Usage
**Run api:**
~~~
$ make run
OR
$ make up 
~~~
Service will be running on http://localhost:8088

**Run tests:**
~~~
$ make test
~~~


**Endpoints:**

```$sh
## TODO - GetAll
curl "http://localhost:8088/todo?token=123"

## TODO - Insert
curl -X "POST" "http://localhost:8088/todo?token=123" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
  "title": "My Task1",
  "user_id": 0
}'

## TODO - Get
curl "http://localhost:8088/todo/1?token=123"

## TODO - Update
curl -X "PUT" "http://localhost:8088/todo/1?token=123" \
     -H 'Content-Type: text/plain; charset=utf-8' \
     -d $'{"title":"My Task2","user_id":0}
'

## TODO - Delete
curl -X "DELETE" "http://localhost:8088/todo/1?token=123"

```

## Second implementation - Proxy Checker
Proxy checker for SOCKS4(5) HTTP(S) protocols
### Usage
**Run api:**
~~~
$ make run
OR
$ make up 
~~~
Service will be running on http://localhost:8088

**Endpoints:**
```$sh
## Proxy Check
curl -X "POST" "http://localhost:8088/check" \
     -H 'Content-Type: application/json; charset=utf-8' \
     -d $'{
  "ip": "191.7.209.186",
  "port": "31576"
}'
```

## Third implementation - Struct viewer

### Usage
**Run script:**
~~~
$ go run scripts/test.go
~~~