package repository

import (
	"fmt"

	"github.com/DiogoPires22/imersao-go/domain/model"
	"gorm.io/gorm"
)

//AccountRepositoryDb is a implementtion of AccountRepositoryInterface
type AccountRepositoryDb struct {
	db *gorm.DB
}

//Register is a method used to create a new account
func (r AccountRepositoryDb) Register(account *model.Account) error {
	err := r.db.Create(account).Error

	if err != nil {
		return err
	}
	return nil
}

//Save is a method used to update a account
func (r AccountRepositoryDb) Save(account *model.Account) error {
	err := r.db.Save(account).Error

	if err != nil {
		return err
	}
	return nil
}

//FindById is a method to search a account by id
func (r AccountRepositoryDb) FindById(id string) (*model.Account, error) {
	var account model.Account

	r.db.Preload("Bank").First(account, "ID =?", id)

	if account.ID == "" {
		return nil, fmt.Errorf("no key found")
	}
	return &account, nil
}
