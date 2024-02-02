package virtualBank

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

type MockUserService struct{}

func (m *MockUserService) AddAccount(account *Account, db *sql.DB) (int, error) {
	return 0, nil
}
func (m *MockUserService) GetAccount(accountNumber int, db *sql.DB) (*Account, error) {
	var testAccount = Account{
		name:              "John Smith",
		pin:               1234,
		overdraftLimit:    2000,
		balance:           0,
		bankAccountNumber: 12345678,
	}
	return &testAccount, nil
}

func (m *MockUserService) UpdateBalance(amount float64, accountNumber int, db *sql.DB) (float64, error) {
	return 0, nil
}

func mockScan(input ...interface{}) (n int, err error) {
	var typeOfInput = reflect.TypeOf(input[0]).String()
	if typeOfInput == "*string" {
		*input[0].(*string) = "Test"
	} else if typeOfInput == "*int" {
		*input[0].(*int) = 1234
	}
	return 0, nil
}

func mockGetValueFromUser(input any, question string) {

	for _, q := range Questions {
		if q.Value == question {
			var typeOfInput = reflect.TypeOf(input).String()

			switch typeOfInput {
			case "*int":
				*input.(*int) = q.ExampleAnswer.(int)
			case "*string":
				*input.(*string) = q.ExampleAnswer.(string)
			case "*float64":
				*input.(*float64) = q.ExampleAnswer.(float64)
			}

			break

		}
	}

}

var db, _, _ = sqlmock.New()

func TestCreateAccount(t *testing.T) {
	Scan = mockScan
	GetValueFromUser = mockGetValueFromUser

	var mockUserService MockUserService

	account := createAccount(db, &mockUserService)

	var testAccount = Account{
		name:           "John Smith",
		pin:            1234,
		overdraftLimit: 2000,
		balance:        0,
	}

	if account.name != testAccount.name || account.pin != testAccount.pin {
		t.Error("Fail")
	}

}
