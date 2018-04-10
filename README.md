# gobootcamp
How to run:
cd rest
go build restmain.go
./restmain


Create:
curl -sSX POST -d '{"name":"adding sequences","completed" : false,"due":"2018-01-30T00:00:00Z"}' http://localhost:8080/todos

GetAll:
curl http://localhost:8080/todos

Get:
curl http://localhost:8080/todos/1

Update:
curl -sSX PUT -d '{"id":1,"name":"adding sequences","completed" : false,"due":"2018-01-30T00:00:00Z"}' http://localhost:8080/todos
Delete:
curl -sSX DELETE http://localhost:8080/todos/3