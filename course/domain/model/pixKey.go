package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)


const(
	KIND_EMAIL = "email"
	KIND_CPF = "cpf"
	KEY_ACTIVE = "active"
	KEY_INACTIVE = "inactive"
)

type PixKeyRepositoryInterface interface {
	Register(pixKey *PixKey) (error)
	Save(pixKey *PixKey) error
	FindKeyByKind(key string, kind string) (*PixKey, error)
}

type PixKey struct {
	Base      `valid:"required"`
	Kind      string   `json:"kind" valid:"notnull"`
	Key       string   `json:"key" valid:"notnull"`
	Account   *Account `valid:"-"`
	AccountID string   `gorm:"column:account_id;type:uuid;not null" valid:"-"`
	Status    string   `json:"status" valid:"notnull"`
}

func (pixKey *PixKey) isValid() error {
	_, err := govalidator.ValidateStruct(pixKey)

	if pixKey.Kind != KIND_EMAIL && pixKey.Kind != KIND_CPF {
		return errors.New("invalid type of key")
	}

	if pixKey.Status != KEY_ACTIVE && pixKey.Status != KEY_INACTIVE {
		return errors.New("invalid status")
	}

	if err != nil {
		return err
	}

	return nil
}

func NewPixKey(account *Account, kind string, key string) (*PixKey, error) {
	pixKey := PixKey{
		Kind:      kind,
		Key:       key,
		Account:   account,
		AccountID: account.ID,
		Status:    "active",
	}

	pixKey.ID = uuid.NewV4().String()
	pixKey.CreatedAt = time.Now()
	pixKey.UpdatedAt = time.Now()

	e := pixKey.isValid()

	if e != nil {
		return nil, e
	}

	return &pixKey, nil
}
