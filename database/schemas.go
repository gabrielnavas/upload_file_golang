package database

const createProducts = `
	CREATE TABLE IF NOT EXISTS product (
		id SERIAL PRIMARY KEY,
		name TEXT,
		file BYTEA
	);
`
