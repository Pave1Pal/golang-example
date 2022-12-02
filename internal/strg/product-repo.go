package strg

import (
	"database/sql"
	"errors"

	"example.com/internal/domain/entity"
	"github.com/google/uuid"
)

type IProductRepository interface {
	FindAll() ([]entity.Product, error)
	FindById(uuid.UUID) (*entity.Product, error)
	Create(*entity.Product) (*entity.Product, error)
	Update(*entity.Product) (*entity.Product, error)
	Delete(id uuid.UUID) (*uuid.UUID, error)
}

type ProductRepository struct {
	db *sql.DB
}

func (p ProductRepository) FindAll() ([]entity.Product, error) {
	var products []entity.Product
	stmt, err := p.db.Prepare("SELECT prd.id, prd.name, prd.price FROM product prd")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var product entity.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (p ProductRepository) FindById(id uuid.UUID) (*entity.Product, error) {
	product := entity.Product{}
	stmt, err := p.db.Prepare(
		"SELECT prd.id, prd.name, prd.price FROM product prd WHERE prd.id=$1")
	if err != nil {
		return nil, err
	}

	if err := stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price); err != nil {
		return nil, err
	}
	return &product, nil
}

func (p ProductRepository) Create(product *entity.Product) (*entity.Product, error) {
	var created entity.Product

	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO product(id, name, price) VALUES ($1, $2, $3) RETURNING id, name, price")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(uuid.New(), product.Name, product.Price)
	if err = row.Scan(&created.Id, &created.Name, &created.Price); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return &created, nil
}

func (p ProductRepository) Delete(id uuid.UUID) (*uuid.UUID, error) {
	var deletedID uuid.UUID

	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("DELETE FROM product WHERE id=$1 RETURNING id")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	if err := row.Scan(&deletedID); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &id, nil
}

func (p ProductRepository) Update(product *entity.Product) (*entity.Product, error) {
	if product.Id == uuid.Nil {
		return nil, errors.New("updated product does not have id")
	}
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare("UPDATE product SET price=$1, name=$2 WHERE id=$3 RETURNING id, name, price")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()
	var updated entity.Product

	row := stmt.QueryRow(product.Price, product.Name, product.Id)
	if err := row.Scan(&updated.Id, &updated.Name, &updated.Price); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &updated, nil
}

func NewProductRepository(db *sql.DB) IProductRepository {
	return &ProductRepository{db: db}
}
