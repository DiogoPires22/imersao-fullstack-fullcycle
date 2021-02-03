package model_test

import (
	"testing"

	"github.com/DiogoPires22/imersao-go/tests/factory"

	uuid "github.com/satori/go.uuid"

	"github.com/DiogoPires22/imersao-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestNewTransaction(t *testing.T) {
	bank := factory.ValidBank()

	account := factory.AccountWithBank(bank)

	pixKey := factory.PixKeyEmailInactive()

	require.NotEqual(t, account.ID, pixKey.Account.ID)

	amount := 3.10
	statusTransaction := "pending"
	transaction, err := model.NewTransaction(account, amount, pixKey, "My description")
	//
	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(transaction.ID))
	require.Equal(t, transaction.Amount, amount)
	require.Equal(t, transaction.Status, statusTransaction)
	require.Equal(t, transaction.Description, "My description")
	require.Empty(t, transaction.CancelDescription)

	pixKeySameAccount := factory.PixKeyWithAccount(account)

	_, err = model.NewTransaction(account, amount, pixKeySameAccount, "My description")
	require.NotNil(t, err)

	_, err = model.NewTransaction(account, 0, pixKey, "My description")
	require.NotNil(t, err)

}

func TestModel_ChangeStatusOfATransaction(t *testing.T) {
	transaction := factory.ValidTransaction()

	_ = transaction.Complete()
	require.Equal(t, transaction.Status, model.TransactionCompleted)

	_ = transaction.Cancel("Error")
	require.Equal(t, transaction.Status, model.TransactionCancelled)
	require.Equal(t, transaction.CancelDescription, "Error")

}
