package repository

import (
	"fmt"

	"github.com/DiogoPires22/imersao-go/domain/model"
	"gorm.io/gorm"
)

//PixKeyRepositoryDb is a implementtion of PixKeyRepositoryInterface
type PixKeyRepositoryDb struct {
	Db *gorm.DB
}

//Register is a method used to create a new pixkey
func (r PixKeyRepositoryDb) Register(account *model.Account) error {
	err := r.Db.Create(account).Error

	if err != nil {
		return err
	}
	return nil
}

//Save is a method used to update a pixkey
func (r PixKeyRepositoryDb) Save(account *model.Account) error {
	err := r.Db.Save(account).Error

	if err != nil {
		return err
	}
	return nil
}

//FindKeyByKind is a method to search a PixKey key in a kind
func (r PixKeyRepositoryDb) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	var pixKey model.PixKey

	r.Db.Preload("Account.Bank").First(pixKey, "kind =? AND key =?", kind, key)

	if pixKey.ID == "" {
		return nil, fmt.Errorf("no pixkey found")
	}
	return &pixKey, nil
}
