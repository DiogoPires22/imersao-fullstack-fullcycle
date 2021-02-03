package factory

import (
	"github.com/DiogoPires22/imersao-go/domain/model"
	uuid "github.com/satori/go.uuid"
	"syreclabs.com/go/faker"
)

func ValidAccount() *model.Account {
	bank := ValidBank()
	return &model.Account{
		OwnerName: faker.Company().Name(),
		Number:    faker.Number().Number(20),
		Bank:      bank,
		BankID:    bank.ID,
	}
}

func AccountWithBank(bank *model.Bank) *model.Account {
	account := &model.Account{
		OwnerName: faker.Company().Name(),
		Number:    faker.Number().Number(20),
		Bank:      bank,
		BankID:    bank.ID,
	}
	account.ID = uuid.NewV4().String()
	return account
}
