package virtualBank

import (
	"database/sql"
	"fmt"
	"log"
)

var accounts map[int]*Account = make(map[int]*Account)
var session int

// Monkey hoisted functions
var Scan = fmt.Scan
var GetValueFromUser = getValueFromUser

func StartBank(db *sql.DB, userService UserService) {
	var actionChoice int

	GetValueFromUser(&actionChoice, "What do you want do? \n1. Create an account\n2. Check balance\n3. Deposit money\n4. Withdraw money\n5. Transfer money\n6. Close your account")

	if actionChoice == 1 {
		newAccount := createAccount(db, userService)
		accounts[newAccount.bankAccountNumber] = &newAccount
		session = newAccount.bankAccountNumber
		StartBank(db, userService)
	}

	if session == 0 {
		checkIfLoggedIn(db, userService)
	}

	// Protected Options
	account, err := userService.GetAccount(session, db)
	if err != nil {
		log.Fatal(err)

	}

	if actionChoice == 2 {
		fmt.Printf("Your balance is £%v \n", account.balance)

	} else if actionChoice == 3 {
		depositMoney(account.bankAccountNumber, db, userService)
	}
	StartBank(db, userService)

}

func createAccount(db *sql.DB, userService UserService) Account {
	var name string
	var pin int
	var income float64

	GetValueFromUser(&name, "What is your name? ")
	GetValueFromUser(&pin, "What would you like your pin to be? ")
	GetValueFromUser(&income, "What is your income? ")

	var account = new(Account)
	account.NewAccount(name, pin, income)

	_, err := userService.AddAccount(account, db)
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

func depositMoney(accountNumber int, db *sql.DB, userService UserService) {
	var amount float64
	GetValueFromUser(&amount, "How much do you want to deposit?")
	balance, err := userService.UpdateBalance(amount, accountNumber, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Your new balance is £%v\n", balance)
}

func checkIfLoggedIn(db *sql.DB, userService UserService) {
	var bankAccountNumber int
	var pin int
	if session != 0 {
		return
	}
	GetValueFromUser(&bankAccountNumber, "What is your bank account number?")
	GetValueFromUser(&pin, "What is your pin?")

	account, err := userService.GetAccount(bankAccountNumber, db)
	if err != nil {
		log.Fatal(err)
	}
	if account.pin != pin {
		fmt.Println("Incorrect pin please try again")
		checkIfLoggedIn(db, userService)
	}

}

func getValueFromUser(field any, question string) {
	fmt.Println(question)
	_, err := Scan(field)
	if err != nil {
		fmt.Println("You entered it wrong please try again")
		getValueFromUser(field, question)

	}

}
