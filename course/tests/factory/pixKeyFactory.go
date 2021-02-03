package factory

import (
	"github.com/DiogoPires22/imersao-go/domain/model"
	faker2 "github.com/bxcodec/faker/v3"
	"syreclabs.com/go/faker"
)

const (
	CPFFormatPattern string = `([\d]{3})([\d]{3})([\d]{3})([\d]{2})`
)

func PixKeyEmailActive() *model.PixKey {
	return pixKeyEmail(model.KEY_ACTIVE)
}

func PixKeyEmailInactive() *model.PixKey {
	return pixKeyEmail(model.KEY_INACTIVE)
}

func pixKeyEmail(status string) *model.PixKey {
	account := ValidAccount()
	return &model.PixKey{
		Kind:      model.KIND_EMAIL,
		Account:   account,
		AccountID: account.ID,
		Key:       faker2.Email(),
		Status:    status,
	}
}

func PixKeyCpfActive() *model.PixKey {
	return pixKeyCpf(model.KEY_ACTIVE)
}

func PixKeyCpfInactive() *model.PixKey {
	return pixKeyCpf(model.KEY_INACTIVE)
}

func pixKeyCpf(status string) *model.PixKey {
	account := ValidAccount()
	return &model.PixKey{
		Kind:      model.KIND_CPF,
		Account:   account,
		AccountID: account.ID,
		Key:       faker.Number().Number(11), //TODO: refactor after
		Status:    status,
	}
}

func PixKeyWithAccount(account *model.Account) *model.PixKey {
	return &model.PixKey{
		Kind:      model.KIND_CPF,
		Account:   account,
		AccountID: account.ID,
		Key:       faker.Number().Number(11), //TODO: refactor after
		Status:    model.KEY_ACTIVE,
	}
}
