package strg

import (
	"database/sql"
	"example.com/internal/domain/entity"
	"github.com/google/uuid"
)

type ICartRepository interface {
	Creat(cart *entity.Cart) (*entity.Cart, error)
}

type CartRepository struct {
	db *sql.DB
}

func (c CartRepository) Creat(cart *entity.Cart) (*entity.Cart, error) {
	cartId := uuid.New()
	tx, err := c.db.Begin()
	if err != nil {
		return nil, err
	}
	createCartStmt, err := tx.Prepare("INSERT INTO cart(id, price) VALUES ($1, $2) RETURNING id, price")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer createCartStmt.Close()

	row := createCartStmt.QueryRow(cartId, cart.Price)

	var created entity.Cart

	err = row.Scan(&created.Id, &created.Price)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	for _, product := range cart.Products {
		rows, err := tx.Query(
			"INSERT INTO cart_product_merge(cart_id, product_id) VALUES ($1, $2)",
			created.Id, product.Id)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		rows.Close()
	}
	created.Products = cart.Products
	tx.Commit()
	return &created, nil
}

func NewCartRepository(db *sql.DB) ICartRepository {
	return &CartRepository{db: db}
}
