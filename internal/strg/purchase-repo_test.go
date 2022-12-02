package strg

import (
	"database/sql"
	"log"
	"testing"
	"time"

	"example.com/internal/domain/entity"
	"github.com/google/uuid"
)

func insertPurchase(db *sql.DB) *entity.Purchase {

	purchasePerson := "test"
	purchaseAddress := "test"
	createdTime := time.Now()
	insertedCart := insertCart(db)

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err.Error())
	}
	s, err := tx.Prepare("INSERT INTO purchase(id, person, address, date, cart_id) values($1, $2, $3, $4, $5)")
	if err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}
	defer s.Close()

	id := uuid.New()
	if _, err := s.Exec(id, purchasePerson, purchaseAddress, createdTime, insertedCart.Id); err != nil {
		tx.Rollback()
		log.Fatal(err.Error())
	}
	tx.Commit()
	return &entity.Purchase{
		Id:      id,
		Person:  purchasePerson,
		Address: purchaseAddress,
		Date:    createdTime,
		Cart:    entity.Cart{Id: insertedCart.Id},
	}
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
	clearCartTable(db)
}

func TestFindAllPurchases(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	clearPurchaseTable(db)

	id1 := insertPurchase(db).Id
	id2 := insertPurchase(db).Id

	purchaseRps := NewPurchaseRepository(db)

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
		t.Error("find less than 2 purchase")
	}

	clearPurchaseTable(db)
	clearCartTable(db)
	clearProudctTable(db)
}

func TestFindPurchaseByID(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	created := insertPurchase(db)

	var p IPurchaseRepository = PurchaseRepository{db}
	found, err := p.FindById(created.Id)
	if err != nil {
		log.Fatal(err)
	}

	if found.Id != created.Id {
		t.Error("expected", created, "actual", found.Id)
	}
	if found.Cart.Id == uuid.Nil {
		t.Error("product id is nil")
	}

	clearPurchaseTable(db)

}

func TestCreatePurchase(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	clearPurchaseTable(db)

	var p IPurchaseRepository = PurchaseRepository{db}

	createCartId := insertCart(db).Id
	createAddress := "Cool"
	createPerson := "Cool Person"
	createdDate := time.Now()

	purchase := entity.Purchase{
		Address: createAddress,
		Person:  createPerson,
		Date:    createdDate,
		Cart:    entity.Cart{Id: createCartId},
	}

	createdPurchase, err := p.Create(&purchase)
	if err != nil {
		log.Fatal(err.Error())
	}

	if createdPurchase.Id == uuid.Nil {
		t.Error("id for purchase does not created")
	}
	if createdPurchase.Address != createAddress {
		t.Error("expected address: ", createAddress, "actual address: ", createdPurchase.Address)
	}
	if createdPurchase.Date.Day() != createdDate.Day() ||
		createdPurchase.Date.Month() != createdDate.Month() ||
		createdPurchase.Date.Year() != createdDate.Year() {
		t.Error("expected date: ", createdDate, "actual date: ", createdPurchase.Date)
	}

	clearPurchaseTable(db)
	clearCartTable(db)
	clearProudctTable(db)
}

func TestUpdatePurchase(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	created := insertPurchase(db)
	createdID := created.Id

	var repo IPurchaseRepository = PurchaseRepository{db}
	person := "new person"
	address := "new address"

	updatePurchase := &entity.Purchase{
		Id:      createdID,
		Person:  person,
		Address: address}

	updated, err := repo.Update(updatePurchase)
	if err != nil {
		t.Error("update return error")
	}
	if updated.Id != createdID {
		t.Error("expected id: ", createdID, "actual: ", updated.Id)
	}
	if updated.Person != person {
		t.Error("expected person: ", person, "actual: ", updated.Person)
	}
	if updated.Address != address {
		t.Error("expected address: ", address, "actual address: ", updated.Address)
	}
	if updated.Cart.Id != created.Cart.Id {
		t.Error("expected cart id: ", created.Cart.Id, "actual cart id", updated.Cart.Id)
	}
	clearPurchaseTable(db)
	clearCartTable(db)
	clearProudctTable(db)
}

func TestDeletePurchase(t *testing.T) {
	db := createDBConnection()
	defer db.Close()
	createdID := insertPurchase(db).Id

	var pr IPurchaseRepository = NewPurchaseRepository(db)

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
