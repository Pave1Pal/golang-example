package strg

import (
	"database/sql"
	"errors"
	"example.com/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type IPurchaseRepository interface {
	FindAll() ([]entity.Purchase, error)
	FindById(uuid.UUID) (*entity.Purchase, error)
	Create(*entity.Purchase) (*entity.Purchase, error)
	Update(*entity.Purchase) (*entity.Purchase, error)
	Delete(uuid.UUID) (*uuid.UUID, error)
}

type PurchaseRepository struct {
	db *sql.DB
}

func (p PurchaseRepository) FindAll() ([]entity.Purchase, error) {
	var purchaseList []entity.Purchase

	stmt, err := p.db.Prepare("SELECT id, person, address, date, cart_id FROM purchase")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var purchase entity.Purchase
		err := rows.Scan(
			&purchase.Id,
			&purchase.Person,
			&purchase.Address,
			&purchase.Date,
			&purchase.Cart.Id)
		if err != nil {
			return nil, err
		}
		purchaseList = append(purchaseList, purchase)
	}
	return purchaseList, nil
}

func (p PurchaseRepository) FindById(id uuid.UUID) (*entity.Purchase, error) {
	var purchase entity.Purchase

	stmt, err := p.db.Prepare("SELECT id, person, address, date, cart_id FROM purchase WHERE id=$1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)

	err = row.Scan(
		&purchase.Id,
		&purchase.Person,
		&purchase.Address,
		&purchase.Date,
		&purchase.Cart.Id)

	if err != nil {
		return nil, err
	}

	return &purchase, nil
}

func (p PurchaseRepository) Create(purchase *entity.Purchase) (*entity.Purchase, error) {
	purchase.Id = uuid.New()
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	stmt, err := tx.Prepare("INSERT INTO purchase(id, person, address, date, cart_id) VALUES($1, $2, $3, $4, $5) RETURNING id, person, address, date, cart_id")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	r := stmt.QueryRow(purchase.Id, purchase.Person, purchase.Address, time.Now(), purchase.Cart.Id)

	var created entity.Purchase
	err = r.Scan(&created.Id, &created.Person, &created.Address, &created.Date, &created.Cart.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &created, nil
}

func (p PurchaseRepository) Update(purchase *entity.Purchase) (*entity.Purchase, error) {
	if purchase.Id == uuid.Nil {
		return nil, errors.New("updated purchase does not have id")
	}
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	stmr, err := tx.Prepare(
		"UPDATE purchase " +
			"SET person = $1, address = $2 " +
			"WHERE id = $3 	" +
			"RETURNING id, person, address, date, cart_id")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmr.Close()

	row := stmr.QueryRow(
		purchase.Person,
		purchase.Address,
		purchase.Id)
	var updated entity.Purchase

	err = row.Scan(
		&updated.Id,
		&updated.Person,
		&updated.Address,
		&updated.Date,
		&updated.Cart.Id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &updated, nil
}

func (p PurchaseRepository) Delete(id uuid.UUID) (*uuid.UUID, error) {
	tx, err := p.db.Begin()
	if err != nil {
		return nil, err
	}
	stms, err := tx.Prepare("DELETE FROM purchase WHERE id = $1 RETURNING id")
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stms.Close()

	row := stms.QueryRow(id)
	var deletedId uuid.UUID
	if err = row.Scan(&deletedId); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &deletedId, nil
}

func NewPurchaseRepository(db *sql.DB) IPurchaseRepository {
	return &PurchaseRepository{db}
}
