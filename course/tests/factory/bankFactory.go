package factory

import (
"github.com/DiogoPires22/imersao-go/domain/model"
"syreclabs.com/go/faker"
)

func ValidBank() *model.Bank{
	return &model.Bank{
		Name: faker.Company().Name(),
		Code: faker.Number().Number(20),
	}
}

