package usecase_test

import (
	"errors"
	"github.com/DiogoPires22/imersao-go/application/usecase"
	"github.com/DiogoPires22/imersao-go/domain/model"
	"github.com/DiogoPires22/imersao-go/tests/factory"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type AccountRepositoryMock struct {
	mock.Mock
}

func (m *AccountRepositoryMock) Register(account *model.Account) error {
	args := m.Called(account)
	return args.Error(1)
}

func (m *AccountRepositoryMock) Save(account *model.Account) error {
	args := m.Called(account)
	return args.Error(1)
}
func (m *AccountRepositoryMock) FindById(id string) (*model.Account, error) {
	args := m.Called(id)

	if args.Error(1)!= nil{
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Account), args.Error(1)
}

type PixKeyRepositoryMock struct {
	mock.Mock
}

func (m *PixKeyRepositoryMock) Register(pixKey *model.PixKey) error {
	args := m.Called(pixKey)
	return args.Error(1)
}

func (m *PixKeyRepositoryMock) Save(pixKey *model.PixKey) error {
	args := m.Called(pixKey)
	return args.Error(1)
}
func (m *PixKeyRepositoryMock) FindKeyByKind(key string, kind string) (*model.PixKey, error) {
	args := m.Called(key, kind)
	return args.Get(0).(*model.PixKey), args.Error(1)
}

func TestPixUseCase_RegisterKey(t *testing.T) {

	t.Run("When account not found, should return a error", func(t *testing.T) {
		expectedAccountId := uuid.NewV4().String()
		expectedErrorMessage := "account not found"
		accountRepositoryMock := AccountRepositoryMock{}
		pixkeyRepositoryMock := PixKeyRepositoryMock{}
		accountRepositoryMock.On("FindById", expectedAccountId).Return(nil, errors.New("account not found"))

		usecase := usecase.PixUseCase{
			&pixkeyRepositoryMock,
			&accountRepositoryMock,
		}

		_, err := usecase.RegisterKey("23423", "", expectedAccountId)

		assert.NotNil(t, err)
		assert.Equal(t, expectedErrorMessage, err.Error())
	})
	
	t.Run("When occur an validation error in pixKey should return a error", func(t *testing.T) {
		var expectedAccount = factory.ValidAccount()
		accountRepositoryMock := AccountRepositoryMock{}
		pixkeyRepositoryMock := PixKeyRepositoryMock{}
		accountRepositoryMock.On("FindById", expectedAccount.ID).Return(expectedAccount, nil)

		useCase := usecase.PixUseCase{
			PixKeyRepository:  &pixkeyRepositoryMock,
			AccountRepository: &accountRepositoryMock,
		}

		_, err := useCase.RegisterKey("23423", "", expectedAccount.ID)

		assert.NotNil(t, err)
		assert.Equal(t, "invalid type of key", err.Error())
	})
}
