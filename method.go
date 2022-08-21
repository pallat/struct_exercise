package structexercise

import "fmt"

type Book struct {
	Name   string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("%s by %s", b.Name, b.Author)
}
