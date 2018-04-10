package repos

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// DTO declaration
type Todo struct {
	Id        int       `json:id bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Completed bool      `json:"completed" bson:"completed"`
	Due       time.Time `json:"due" bson:"due"`
}
type Todos []Todo

var todos Todos
var currentID int

// var db *mgo.Database
var todo Todo

// type TodoDAO struct {
// 	Server   string
// 	Database string
// }

const (
	COLLECTION = "todo"
)

/* CreateToDo */
func CreateTodo(t Todo) Todo {
	// currentID++
	// t.Id = currentID

	// todos = append(todos, t)
	var err error
	t.Id, err = getNextSequence("todoId")
	err = db.C(COLLECTION).Insert(t)
	if err != nil {
		log.Fatal(err)
	}
	return t

}
func FindAllTodo() []Todo {
	err := db.C(COLLECTION).Find(bson.M{}).All(&todos)
	if err != nil {
		log.Fatal(err)
	}
	return todos
}
func FindTodo(id int) Todo {
	// stringid := strconv.Itoa(id)

	// err := db.C(COLLECTION).FindId(bson.ObjectIdHex(stringid)).One(&todo)
	err := db.C(COLLECTION).Find(bson.M{"id": id}).One(&todo)
	if err != nil {
		log.Fatal(err)
	}
	// for _, t := range todos {
	// 	if t.Id == id {
	// 		return t
	// 	}
	// }
	return todo
}
func DeleteTodo(id int) error {

	err := db.C(COLLECTION).Remove(bson.M{"id": id})
	// Find(bson.M{"id": id}).One(&todo)

	return fmt.Errorf("Could not remove Todo with ID: %d %s", id, err)
}

func UpdateTodo(t Todo) error {

	// todos = append(todos, t)
	err := db.C(COLLECTION).Find(bson.M{"id": t.Id}).One(&todo)

	if err != nil {
		// return t, fmt.Errorf("Could not find Todo with ID: %d", t.Id)
		return fmt.Errorf("Could not find Todo with ID: %d", t.Id)
	}
	err = db.C(COLLECTION).Update(bson.M{"id": todo.Id}, &t)

	return err

}
