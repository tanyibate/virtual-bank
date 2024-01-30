package virtualBank

import (
	"database/sql"
	"example.com/virtual-bank/database"
	"fmt"
	"log"
)

var accounts map[int]*Account = make(map[int]*Account)
var session int

func StartBank(db *sql.DB) {
	var actionChoice int

	fmt.Println("What do you want do?")
	fmt.Println("1. Create an account")
	fmt.Println("2. Check balance")
	fmt.Println("3. Deposit money")
	fmt.Println("4. Withdraw money")
	fmt.Println("5. Transfer money")
	fmt.Println("6. Close your account")

	fmt.Scan(&actionChoice)

	if actionChoice == 1 {
		newAccount := createAccount(db)
		accounts[newAccount.bankAccountNumber] = &newAccount
		session = newAccount.bankAccountNumber
		StartBank(db)
	}

	if session == 0 {
		checkIfLoggedIn()
	}

	// Protected Options
	account, err := GetUserFromDB(session, db)
	if err != nil {
		log.Fatal(err)

	}

	if actionChoice == 2 {
		fmt.Printf("Your balance is £%v \n", account.balance)

	} else if actionChoice == 3 {
		depositMoney(account.bankAccountNumber, db)
	}
	StartBank(db)

}

func createAccount(db *sql.DB) Account {
	var name string
	var pin int
	var income float64

	getValueFromUser(&name, "What is your name? ")
	getValueFromUser(&pin, "What would you like your pin to be? ")
	getValueFromUser(&income, "What is your income? ")

	var account = new(Account)
	account.NewAccount(name, pin, income)

	_, err := addUserToDB(account, db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Successfully created a bank account for %v with number: %v and overdraft limit of £%v\n", account.name, account.bankAccountNumber, account.overdraftLimit)

	return *account

}

func withdrawMoney(amount float64, account Account) {
	balance, err := account.WithdrawMoney(amount)
	if err != nil {
		fmt.Printf("Your new balance is £%v\n", balance)
	} else {
		fmt.Println(err)
	}

}

func depositMoney(accountNumber int, db *sql.DB) {
	fmt.Println("How much do you want to deposit?")
	var amount float64
	fmt.Scan(&amount)
	balance, err := updateUserBalanceInDB(amount, accountNumber, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your new balance is £%v\n", balance)
}

func getValueFromUser(field any, question string) {
	fmt.Print(question)
	_, err := fmt.Scan(field)
	if err != nil {
		fmt.Println("You entered it wrong please try again")
		getValueFromUser(field, question)

	}

}

func checkIfLoggedIn() {
	var bankAccountNumber int
	var pin int
	if session != 0 {
		return
	}
	getValueFromUser(&bankAccountNumber, "What is your bank account number?")
	getValueFromUser(&pin, "What is your pin?")

	account, err := GetUserFromDB(bankAccountNumber, database.DB)
	if err != nil {
		log.Fatal(err)
	}
	if account.pin != pin {
		fmt.Println("Incorrect pin please try again")
		checkIfLoggedIn()
	}

}
