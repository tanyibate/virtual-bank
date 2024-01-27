package virtualBank

import "fmt"

var accounts map[int]*Account = make(map[int]*Account)
var session int

func StartBank() {
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
		newAccount := createAccount()
		accounts[newAccount.bankAccountNumber] = &newAccount
		session = newAccount.bankAccountNumber
		StartBank()
	}

	if session == 0 {
		checkIfLoggedIn()
	}

	// Protected Options
	account, ok := accounts[session]
	if ok {
		fmt.Printf("Welcome back %v \n", account.name)

	} else {
		fmt.Println("There was an issue with accessing your bank account")
	}

	if actionChoice == 2 {
		fmt.Printf("Your balance is £%v \n", account.balance)

	} else if actionChoice == 3 {
		depositMoney(account)
	}
	StartBank()

}

func createAccount() Account {
	var name string
	var pin int
	var income float64

	getValueFromUser(&name, "What is your name? ")
	getValueFromUser(&pin, "What would you like your pin to be? ")
	getValueFromUser(&income, "What is your income? ")

	var account = new(Account)
	account.NewAccount(name, pin, income)

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

func depositMoney(account *Account) {
	fmt.Println("How much do you want to deposit?")
	var amount float64
	fmt.Scan(&amount)
	balance := account.DepositMoney(amount)
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
	if accounts[bankAccountNumber].pin != pin {
		fmt.Println("Incorrect pin please try again")
		checkIfLoggedIn()
	}

}
