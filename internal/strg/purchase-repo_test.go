package strg

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"example.com/internal/domain/entity"
	"github.com/google/uuid"
)

func insertOnePurchase(db *sql.DB) uuid.UUID {
	createdProductID := insertOneProduct(db)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err.Error())
	}
	s, err := tx.Prepare("INSERT INTO purchase(id, person, address, date, product_id) values($1, $2, $3, $4, $5)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}
	defer s.Close()

	id := uuid.New()
	if _, err := s.Exec(id, "test", "test", time.Now(), createdProductID); err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}
	tx.Commit()
	return id
}

func clearPurchaseTable(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("test: " + err.Error())
	}
	if _, err := tx.Exec("DELETE FROM purchase"); err != nil {
		tx.Rollback()
		log.Fatal("test: " + err.Error())
	}
	tx.Commit()
	clearProudctTable(db)
}

func TestFindAllPurchases(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	clearPurchaseTable(db)

	id1 := insertOnePurchase(db)
	id2 := insertOnePurchase(db)

	var purchaseRps IPurchaseRepository = PurchaseRepository{db: db}

	ps, err := purchaseRps.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	match := 0
	for _, p := range ps {

		if p.Id == id1 || p.Id == id2 {
			match++
		}
	}

	if match != 2 {
		t.Error("find less than 2 pu")
	}

	clearPurchaseTable(db)
}

func TestFindPurchaseByID(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	createdID := insertOnePurchase(db)

	var p IPurchaseRepository = PurchaseRepository{db}
	found, err := p.FindById(createdID)
	if err != nil {
		log.Fatal(err)
	}

	if found.Id != createdID {
		t.Error("expected", createdID, "actual", found.Id)
	}
	if found.Product.Id == uuid.Nil {
		t.Error("product id is nil")
	}

	clearPurchaseTable(db)

}

func TestCreatePurchase(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	clearPurchaseTable(db)

	var p IPurchaseRepository = PurchaseRepository{db}

	createdProductID := insertOneProduct(db)
	createdID := uuid.New()
	createAddress := "Cool"
	createPerson := "Cool Person"
	createdDate := time.Now()

	purchase := entity.Purchase{
		Id:      createdID,
		Address: createAddress,
		Person:  createPerson,
		Date:    createdDate,
		Product: entity.Product{Id: createdProductID},
	}

	createdPurchase, err := p.Create(&purchase)
	if err != nil {
		log.Fatal(err.Error())
	}

	if createdPurchase.Id != createdID ||
		createdPurchase.Address != createAddress ||
		createdPurchase.Date.Day() != createdDate.Day() ||
		createdPurchase.Date.Month() != createdDate.Month() ||
		createdPurchase.Date.Year() != createdDate.Year() ||
		createdPurchase.Person != createPerson ||
		createdPurchase.Product.Id != createdProductID {
		t.Error(
			"expected Id", createdID, "actual id", createdPurchase.Id,
			"expected product id", createdProductID, "actual product id", createdPurchase.Product.Id,
			"expected date", createdDate, "eactual date", createdPurchase.Date,
		)
	}
	clearPurchaseTable(db)
}

func TestUpdatePurchase(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	createdID := insertOnePurchase(db)

	var repo IPurchaseRepository = PurchaseRepository{db}
	person := "new person"
	address := "new address"
	newProdID := insertOneProduct(db)

	updatePurchase := &entity.Purchase{
		Id:      createdID,
		Person:  person,
		Address: address,
		Product: entity.Product{Id: newProdID}}

	updated, err := repo.Update(updatePurchase)
	if err != nil {
		t.Error("update return error")
	}

	if updated.Id != createdID || updated.Person != person || updated.Address != address || updated.Product.Id != newProdID {
		t.Error("expected:", updatePurchase, "actual:", updated)
	}
	clearPurchaseTable(db)
}

func TestDeletePurchase(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	createdID := insertOnePurchase(db)

	var pr IPurchaseRepository = PurchaseRepository{db: db}

	deletedID, err := pr.Delete(createdID)
	if err != nil {
		log.Fatal("test: " + err.Error())
	}

	if *deletedID != createdID {
		t.Error(
			"expected deleted id:", createdID,
			"actual id:", *deletedID)
	}
	clearPurchaseTable(db)
}
