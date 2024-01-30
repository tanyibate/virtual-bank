package virtualBank

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func addUserToDB(account *Account, db *sql.DB) (int, error) {
	var accountNumber int
	row := db.QueryRow("INSERT INTO accounts(name,pin,account_number,balance,overdraft_limit) VALUES($1, $2, $3, $4, $5) RETURNING account_number", account.name, account.pin, account.bankAccountNumber, account.balance, account.overdraftLimit)

	if err := row.Scan(&accountNumber); err != nil {

		return accountNumber, fmt.Errorf("addUserToDB %v %v", account.bankAccountNumber, err)

	}

	return accountNumber, nil

}

type GetUser func(int, *sql.DB) (*Account, error)

func GetUserFromDB(accountNumber int, db *sql.DB) (*Account, error) {
	var account *Account = new(Account)
	row := db.QueryRow("SELECT * FROM accounts WHERE account_number=$1", accountNumber)

	if err := row.Scan(&account.name, &account.pin, &account.bankAccountNumber, &account.balance, &account.overdraftLimit); err != nil {
		if err == sql.ErrNoRows {
			return account, nil
		}
		return nil, fmt.Errorf("getUserFromDB %v %v", accountNumber, err)

	}

	return account, nil

}

type UpdateBalance func(float64, int, *sql.DB) (float64, error)

func updateUserBalanceInDB(amount float64, accountNumber int, db *sql.DB) (float64, error) {
	var balance float64
	row := db.QueryRow("UPDATE accounts SET balance=$1 WHERE account_number=$2 RETURNING balance", amount, accountNumber)
	if err := row.Scan(&balance); err != nil {
		return balance, fmt.Errorf("updateUserBalanceInDB %v %v", accountNumber, err)

	}
	return balance, nil
}
