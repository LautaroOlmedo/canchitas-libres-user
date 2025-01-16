package domain

type Field struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
}

func NewField(name string, number int) Field {
	return Field{
		Name:   name,
		Number: number,
	}
}
