package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type BankRepositoryInterface interface {
	Register(bank *Bank) error
	Save(Bank *Bank) error
	FindById(id string) (*Account, error)
}

type Bank struct {
	Base     `valid:"required"`
	Code     string     `json:"code" gorm:"column:code;type:varchar(20)" valid:"notnull"`
	Name     string     `json:"name" gorm:"column:name;type:varchar(255)" valid:"notnull"`
	Accounts []*Account `gorm:"ForeignKey:BankID" valid:"-"`
}

func (bank *Bank) isValid() error {
	_, err := govalidator.ValidateStruct(bank)

	if err != nil {
		return err
	}

	return nil
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4().String()
	bank.CreatedAt = time.Now()
	bank.UpdatedAt = time.Now()

	e := bank.isValid()

	if e != nil {
		return nil, e
	}

	return &bank, nil
}
