package factory

import (
	"github.com/DiogoPires22/imersao-go/domain/model"
	"math/rand"
)

func ValidTransaction() *model.Transaction {
	account := ValidAccount()
	pixKey := PixKeyEmailActive()
	return &model.Transaction{
		AccountFrom:   account,
		AccountFromID: account.ID,
		PixKeyTo:      pixKey,
		PixKeyIdTo:    pixKey.ID,
		Amount:        rand.Float64(),
		Status:        model.TransactionPending,
	}
}
