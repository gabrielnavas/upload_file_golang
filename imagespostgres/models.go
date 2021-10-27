package imagespostgres

import "bytes"

type Product struct {
	ID   int
	Name string
	File *bytes.Buffer
}

func NewProduct(name string) *Product {
	return &Product{
		ID:   0,
		Name: name,
		File: bytes.NewBuffer(nil),
	}
}
