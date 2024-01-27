package virtualBank

import (
	"errors"
	"math/rand"
	"time"
)

var overdraftLimits = map[float64]float64{
	0:      0,
	10000:  100,
	50000:  1000,
	100000: 2000,
}

type Account struct {
	name              string
	bankAccountNumber int
	pin               int
	balance           float64
	overdraftLimit    float64
}

func generateBankAccountNumber() int {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	val := r.Intn(89999999) + 10000000
	return val
}

func (account *Account) NewAccount(name string, pin int, income float64) {

	account.bankAccountNumber = generateBankAccountNumber()
	account.name = name
	account.pin = pin

	for key, value := range overdraftLimits {
		if key <= income {
			account.overdraftLimit = value
		}
	}

}

func (account *Account) WithdrawMoney(amount float64) (float64, error) {
	if account.overdraftLimit+account.balance >= amount {
		account.balance -= amount
		return account.balance, nil

	}
	return account.balance, errors.New("balance is not sufficient to cover transaction")

}

func (account *Account) DepositMoney(amount float64) float64 {
	account.balance += amount
	return account.balance

}
