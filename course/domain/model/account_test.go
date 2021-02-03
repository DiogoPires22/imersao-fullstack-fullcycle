package model_test

import (
	"github.com/DiogoPires22/imersao-go/domain/model"
	"github.com/DiogoPires22/imersao-go/tests/factory"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestModel_NewAccount(t *testing.T) {
	account := factory.ValidAccount()
	_, err :=  model.NewAccount(account.Bank, "", account.OwnerName)
	require.NotNil(t, err)
}
