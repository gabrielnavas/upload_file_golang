package imagespostgres

import (
	"bytes"
	"database/sql"
	"io"
)

type Repository interface {
	Store(product *Product) (int, error)
	Get(id int) (*Product, error)
}

type RepositoryPostgres struct {
	db *sql.DB
}

func (p *Product) SetFile(src io.Reader) error {
	if _, err := io.Copy(p.File, src); err != nil {
		return err
	}
	return nil
}

func (p *Product) SetFileFromBuffer(src []byte) {
	p.File = bytes.NewBuffer(src)
}

func (repo *RepositoryPostgres) Store(product *Product) (int, error) {
	const sqlInsertProduct = `
		INSERT INTO product (name, file)
		VALUES ($1, $2)
		RETURNING id;
	`
	row := repo.db.QueryRow(sqlInsertProduct, product.Name, product.File.Bytes())
	if row.Err() != nil {
		return 0, row.Err()
	}
	var id int
	row.Scan(&id)
	return id, nil
}

func (repo *RepositoryPostgres) Get(id int) (*Product, error) {
	const sqlInsertProduct = `
		SELECT id, name, file 
		FROM product
		WHERE id = $1;
	`
	row := repo.db.QueryRow(sqlInsertProduct, id)
	if row.Err() != nil {
		return nil, row.Err()
	}

	product := NewProduct("")
	var fileBuffer []byte
	err := row.Scan(&product.ID, &product.Name, &fileBuffer)
	if err != nil {
		return nil, err
	}
	product.SetFileFromBuffer(fileBuffer)
	return product, nil
}
