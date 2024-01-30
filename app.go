package main

import virtualBank "example.com/virtual-bank/virtual-bank"
import database "example.com/virtual-bank/database"

func main() {
	database.InitDB()
	var userService virtualBank.BaseUserService
	virtualBank.StartBank(database.DB, &userService)

}
