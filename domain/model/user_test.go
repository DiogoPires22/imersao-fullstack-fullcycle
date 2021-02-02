package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/DiogoPires22/imersao-go/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {
	name := "Diogo Pires Caires Silva"
	email := "diogo.123@aaa.com"
	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotEmpty(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser("", "")
	require.NotNil(t, err)
}
