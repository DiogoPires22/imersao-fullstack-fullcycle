package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type AccountRepositoryInterface interface {
	Register(account *Account) error
	Save(account *Account) error
	FindById(id string) (*Account, error)
}

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `json:"owner_name" gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"notnull"`
	Number    string    `json:"number" gorm:"column:number;type:varchar(20);not null" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"`
}

func (account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(account)

	if err != nil {
		return err
	}

	return nil
}

func NewAccount(bank *Bank, owner_name string, number string) (*Account, error) {

	account := Account{
		OwnerName: owner_name,
		Number:    number,
		Bank:      bank,
		BankID:    bank.ID,
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()
	account.UpdatedAt = time.Now()

	err := account.isValid()

	if err != nil {
		return nil, err
	}

	return &account, nil
}
