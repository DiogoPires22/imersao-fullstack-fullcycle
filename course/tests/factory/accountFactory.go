package factory

import (
	"github.com/DiogoPires22/imersao-go/domain/model"
	"syreclabs.com/go/faker"
)

func ValidAccount() *model.Account{
	bank := ValidBank()
	return &model.Account{
		OwnerName: faker.Company().Name(),
		Number: faker.Number().Number(20),
		Bank: bank,
		BankID: bank.ID,
	}
}
