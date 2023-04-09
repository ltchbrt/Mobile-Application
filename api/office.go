package api

import (
	"net/http"
	"strconv"

	"github.com/dafalo/Mobile-Tracker-App/models"
)

func AddOffice(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := models.Office{}
	name := r.FormValue("name")
	number := r.FormValue("number")
	address := r.FormValue("address")
	users, _ := r.Cookie("id")

	contact.Name = name
	contact.Number = number
	contact.Address = address
	contact.UserID = users.Value

	db.Save(&contact)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetOffice(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := []models.Office{}
	users, _ := r.Cookie("id")
	db.Where("user_id = ?", users.Value).Order("name").Find(&contact)

	data := map[string]interface{}{
		"status": "ok",
		"item":   contact,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func EditOffice(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := models.Office{}
	id := r.FormValue("id")
	name := r.FormValue("name")
	number := r.FormValue("number")
	address := r.FormValue("address")

	db.Where("id = ?", id).Find(&contact)

	contact.Name = name
	contact.Number = number
	contact.Address = address

	db.Save(&contact)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func DeleteOffice(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	contact := models.Office{}

	db.Where("id", id).Statement.Delete(&contact)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}
