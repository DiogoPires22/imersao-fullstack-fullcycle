package usecase

import (
	"errors"

	"github.com/DiogoPires22/imersao-go/domain/model"
)

type PixUseCase struct {
	PixKeyRepository  model.PixKeyRepositoryInterface
	AccountRepository model.AccountRepositoryInterface
}

//RegisterKey is a PixKey use case method used to create a pixkey
func (usecase *PixUseCase) RegisterKey(key string, kind string, accountID string) (*model.PixKey, error) {

	account, err := usecase.AccountRepository.FindById(accountID)

	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, kind, key)

	if err != nil {
		return nil, err
	}

	usecase.PixKeyRepository.Register(pixKey)

	if pixKey.ID == "" {
		return nil, errors.New("unable to create a PixKey")
	}

	return pixKey, nil

}

//FindKey is a PixKey use case method used to find a key by kind
func (usecase *PixUseCase) FindKey(kind string, key string) (*model.PixKey, error) {

	pixKey, err := usecase.PixKeyRepository.FindKeyByKind(key, kind)

	if err != nil {
		return nil, err
	}

	return pixKey, nil
}
