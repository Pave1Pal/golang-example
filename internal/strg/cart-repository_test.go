package strg

import (
	"database/sql"
	"example.com/internal/domain/entity"
	"github.com/google/uuid"
	"log"
	"testing"
)

func TestCreateCart(t *testing.T) {
	db := createDBConnection()
	defer db.Close()

	prodId := insertOneProduct(db)

	cartPrice := int64(3000)
	products := []entity.Product{{Id: prodId}}

	cart := entity.Cart{
		Price:    cartPrice,
		Products: products,
	}

	cartRepo := CartRepository{db}

	created, err := cartRepo.Creat(&cart)
	if err != nil {
		log.Fatal("")
	}

	stmt, err := db.Prepare("SELECT product_id FROM cart_product_merge WHERE product_id = $1")
	if err != nil {
		log.Fatal("make statement error")
	}
	row := stmt.QueryRow(&prodId)
	var createdProductId uuid.UUID
	if err = row.Scan(&createdProductId); err != nil {
		log.Fatal("bind created product id error")
	}

	if created.Price != cartPrice {
		t.Error("expected price: ", cartPrice, "actual price: ", cart.Price)
	}
	if created.Id == uuid.Nil {
		t.Error("cart id was not create")
	}
	if createdProductId != prodId {
		t.Error("expected product id: ", prodId, "actual product id: ", createdProductId)
	}

	clearCartTable(db)
	clearProudctTable(db)
}

func clearCartTable(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err.Error())
	}

	delFromMerge, err := tx.Prepare("DELETE FROM cart_product_merge")
	if _, err := delFromMerge.Exec(); err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}

	deleteCart, err := tx.Prepare("DELETE FROM cart")
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}
	if _, err := deleteCart.Exec(); err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}
	tx.Commit()
}

func insertCart(db *sql.DB) *entity.Cart {
	cartId := uuid.New()
	cartPrice := int64(3000)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err.Error())
	}
	createCartStmt, err := tx.Prepare("INSERT INTO cart(id, price) VALUES ($1, $2) RETURNING id, price")
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}
	defer createCartStmt.Close()

	row := createCartStmt.QueryRow(cartId, cartPrice)

	var created entity.Cart

	err = row.Scan(&created.Id, &created.Price)
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}

	productId := insertOneProduct(db)

	toMergeTableStmt, err := tx.Prepare(
		"INSERT INTO cart_product_merge(cart_id, product_id) VALUES ($1, $2)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}
	_, err = toMergeTableStmt.Exec(&created.Id, &productId)
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())

	}
	created.Products = []entity.Product{{Id: productId}}
	tx.Commit()
	return &created
}
