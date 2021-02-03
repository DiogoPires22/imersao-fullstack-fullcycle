package model_test

import (
	"github.com/DiogoPires22/imersao-go/tests/factory"
	"testing"

	"github.com/DiogoPires22/imersao-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	account := factory.ValidAccount()

	pixKey := factory.PixKeyEmailActive()

	_, err := model.NewPixKey(pixKey.Account, pixKey.Kind, pixKey.Key)
	require.Nil(t, err)

	_, err = model.NewPixKey(account, "nome", "-")
	require.NotNil(t, err)
}
