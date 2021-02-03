package repository

import (
	"fmt"

	"github.com/DiogoPires22/imersao-go/domain/model"
	"gorm.io/gorm"
)

type TransactionRepositoryDb struct {
	Db *gorm.DB
}

//Register is a method used to create a new pixkey
func (r TransactionRepositoryDb) Register(transaction *model.Transaction) error {
	err := r.Db.Create(transaction).Error

	if err != nil {
		return err
	}
	return nil
}

//Save is a method used to update a transaction
func (r TransactionRepositoryDb) Save(transaction *model.Transaction) error {
	err := r.Db.Save(transaction).Error

	if err != nil {
		return err
	}
	return nil
}

//FindKeyByKind is a method to search a transaction id
func (r TransactionRepositoryDb) FindKeyByKind(id string) (*model.Transaction, error) {
	var transaction model.Transaction

	r.Db.Preload("AccountFrom.Bank").Preload("PixKeyTo.Account").First(&transaction, "id = ?", id)

	if transaction.ID == "" {
		return nil, fmt.Errorf("no transaction found")
	}
	return &transaction, nil
}
