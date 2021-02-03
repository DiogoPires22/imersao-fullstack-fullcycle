package repository

import (
	"fmt"

	"github.com/DiogoPires22/imersao-go/domain/model"
	"gorm.io/gorm"
)

//BankRepositoryDb is a implementtion of BankRepositoryInterface
type BankRepositoryDb struct {
	Db *gorm.DB
}

//Register is a method used to create a new bank
func (r BankRepositoryDb) Register(bank *model.Bank) error {
	err := r.Db.Create(bank).Error

	if err != nil {
		return err
	}
	return nil
}

//Save is a method used to update a bank
func (r BankRepositoryDb) Save(bank *model.Bank) error {
	err := r.Db.Save(bank).Error

	if err != nil {
		return err
	}
	return nil
}

//FindById is a method to search a bank by id
func (r BankRepositoryDb) FindById(id string) (*model.Bank, error) {
	var bank model.Bank

	r.Db.Preload("Bank").First(bank, "ID =?", id)

	if bank.ID == "" {
		return nil, fmt.Errorf("no bank found")
	}
	return &bank, nil
}
