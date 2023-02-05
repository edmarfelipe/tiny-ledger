package entity

import (
	"time"

	"github.com/google/uuid"
)

func NewPerson(name string, doc string, birth time.Time) *Person {
	return &Person{
		ID:    uuid.New().String(),
		Name:  name,
		Doc:   doc,
		Birth: birth,
	}
}

type Person struct {
	ID    string    `json:"id"`
	Name  string    `json:"name"`
	Doc   string    `json:"doc"`
	Birth time.Time `json:"birth"`
}

func (p *Person) Create() {
	p.ID = uuid.New().String()
}
