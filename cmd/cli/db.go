package main

import (
	"bm-support/config/database"
	"bm-support/src/models"

	"github.com/fatih/color"
)

func migrateDB() {
	color.Yellow("DB start migration")
	db := database.OpenConnectionToDB()
	db.AutoMigrate(models.User{}, models.Token{})
	database.Connection = database.ReturnConnection(db)
	database.CloseConnection()
	color.Green("DB migration finished")
}

func seedByModel(model string) {
	db := database.OpenConnectionToDB()
	if model == "superadmin" {
		var user models.User
		// user.FirstName = models.JSONB{
		// 	"ar": "محمود",
		// 	"en": "Mahmoud",
		// }
		// user.LastName = models.JSONB{
		// 	"ar": "أحمد",
		// 	"en": "Ahmed",
		// }
		user.Email = "admin@gmail.com"
		user.SetPassword("123456")
		db.Create(&user)
	}
	database.Connection = database.ReturnConnection(db)
	database.CloseConnection()
}
