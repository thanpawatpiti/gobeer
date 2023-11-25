package main

import "github.com/thanpawatpiti/gobeer/conf/database"

func main() {
	database.InitMariaDB()
	database.InitMongoDB()
}
