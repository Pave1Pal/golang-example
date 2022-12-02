package strg

import (
	"database/sql"
	"log"
	"testing"

	"example.com/internal/domain/entity"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func createDBConnection() *sql.DB {
	db, err := sql.Open("postgres", "user=postgres password=postgres host=localhost dbname=t_golang_db sslmode=disable")
	if err != nil {
		log.Fatal("test: can not connect to DB")
	}
	if err := db.Ping(); err != nil {
		log.Fatal("test: can not connect to DB")
	}
	return db
}

func insertOneProduct(db *sql.DB) uuid.UUID {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("test: " + err.Error())
	}
	stmt, err := tx.Prepare("INSERT INTO product(id, name, price) VALUES ($1, $2, $3)")
	if err != nil {
		tx.Rollback()
		log.Fatal("test: " + err.Error())
	}
	createdID := uuid.New()
	if _, err := stmt.Exec(createdID, "test", int64(30000)); err != nil {
		tx.Rollback()
		log.Fatal("test: " + err.Error())
	}
	tx.Commit()
	return createdID
}

func insertProducts(db *sql.DB) []uuid.UUID {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("test: " + err.Error())
	}
	var ids []uuid.UUID
	for i := 0; i < 3; i++ {
		stmt, err := tx.Prepare("INSERT INTO product(id, name, price) VALUES ($1, $2, $3)")
		if err != nil {
			tx.Rollback()
			log.Fatal("test: " + err.Error())
		}
		createdID := uuid.New()
		if _, err := stmt.Exec(createdID, "test", int64(3000)); err != nil {
			tx.Rollback()
			log.Fatal("test: " + err.Error())
		}
		ids = append(ids, createdID)
	}
	tx.Commit()
	return ids
}

func clearProudctTable(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("test: " + err.Error())
	}
	if _, err := tx.Exec("DELETE FROM product"); err != nil {
		tx.Rollback()
		log.Fatal("test: " + err.Error())
	}
	tx.Commit()
}

func TestFindAllProducts(t *testing.T) {
	var prdRepo IProductRepository

	db := createDBConnection()
	clearProudctTable(db)
	insertProducts(db)

	prdRepo = NewProductRepository(db)
	prds, err := prdRepo.FindAll()
	if err != nil {
		log.Fatal("test: " + err.Error())
	}
	prdSize := len(prds)
	if prdSize != 3 {
		t.Error("Expected 3, got", prdSize)
	}
	clearProudctTable(db)
}

func TestFindProductByID(t *testing.T) {
	var prdRepo IProductRepository
	db := createDBConnection()
	defer db.Close()
	clearProudctTable(db)
	createdID := insertOneProduct(db)
	prdRepo = NewProductRepository(db)

	p, err := prdRepo.FindById(createdID)
	if err != nil {
		log.Fatal("test: " + err.Error())
	}
	if p.Id != createdID {
		t.Error("actual id", p.Id, "expected id", createdID)
	}
	clearProudctTable(db)
}

func TestDelete(t *testing.T) {
	var prdRepo IProductRepository
	db := createDBConnection()
	defer db.Close()
	clearProudctTable(db)
	createdID := insertOneProduct(db)

	prdRepo = NewProductRepository(db)

	deletedID, err := prdRepo.Delete(createdID)
	if err != nil {
		log.Fatal("test: " + err.Error())
	}

	if *deletedID != createdID {
		t.Error("expected:", createdID, "actual:", deletedID)
	}
	clearProudctTable(db)
}

func TestUpdate(t *testing.T) {
	var rep IProductRepository
	db := createDBConnection()
	defer db.Close()
	clearProudctTable(db)
	targetID := insertOneProduct(db)
	rep = NewProductRepository(db)

	newName := "updated"
	newPrice := int64(4000)

	prd := entity.Product{
		Id:    targetID,
		Name:  newName,
		Price: newPrice,
	}

	updatedPrd, err := rep.Update(&prd)
	if err != nil {
		log.Fatal("test: " + err.Error())
	}

	actualName := updatedPrd.Name
	actualPrice := updatedPrd.Price

	if actualName != newName ||
		actualPrice != newPrice {
		t.Error(
			"actual Name:", actualName,
			"expected Name:", newName,
			"actual Price:", actualPrice,
			"expected Price:", newPrice,
		)
	}
	clearProudctTable(db)
}
