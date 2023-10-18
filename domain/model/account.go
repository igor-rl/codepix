package model

import (
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Account struct {
	Base      `valid:"required"`
	OwnerName string    `gorm:"column:owner_name;type:varchar(255);not null" valid:"notnull"`
	Bank      *Bank     `valid:"-"`
	BankID    string    `gorm:"column:bank_id;type:uuid;not null" valid:"-"`
	Number    string    `json:"number" gorm:"type:varchar(20)" valid:"notnull"`
	PixKeys   []*PixKey `gorm:"ForeignKey:AccountID" valid:"-"`
}

func (Account *Account) isValid() error {
	_, err := govalidator.ValidateStruct(Account)
	if err != nil {
		return err
	}
	return nil
}

func NewAccount(bank *Bank, number string, owner_name string) (*Account, error) {
	Account := Account{
		OwnerName: owner_name,
		Bank:      bank,
		Number:    number,
	}
	Account.ID = uuid.NewV4().String()
	Account.CreatedAt = time.Now()
	err := Account.isValid()
	if err != nil {
		return nil, err
	}
	return &Account, nil
}
