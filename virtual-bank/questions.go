package virtualBank

var Questions = map[string]Question{
	"introductoryOptions": {
		Value:         "What do you want do? \n1. Create an account\n2. Check balance\n3. Deposit money\n4. Withdraw money\n5. Transfer money\n6. Close your account",
		ExampleAnswer: 1,
	},
	"getName": {
		Value:         "What is your name? ",
		ExampleAnswer: "John Smith",
	},
	"getPin": {
		Value:         "What would you like your pin to be? ",
		ExampleAnswer: int(1234),
	},
	"getIncome": {
		Value:         "What is your income? ",
		ExampleAnswer: float64(123456.0),
	},
}
