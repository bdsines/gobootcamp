package repos

import (
	"time"
)

type Todo struct {
	Id        int       `json:id bson:"id"`
	Name      string    `json:"name" bson:"name"`
	Completed bool      `json:"completed" bson:"completed"`
	Due       time.Time `json:"due" bson:"due"`
}
type Todos []Todo
