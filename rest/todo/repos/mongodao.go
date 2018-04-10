package repos

import (
	"log"

	"gopkg.in/mgo.v2"
)

type MongoDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

func (t *MongoDAO) connect() {
	session, err := mgo.Dial(t.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(t.Database)
}
func init() {
	var dao = MongoDAO{}
	dao.Server = "localhost"
	dao.Database = "test"
	dao.connect()
	// CreateTodo(Todo{Name: " Enhancing repo"})
	// CreateTodo(Todo{Name: " Adding CRUD "})

}
