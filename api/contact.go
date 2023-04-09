package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dafalo/Mobile-Tracker-App/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func AddContact(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := models.Contact{}
	name := r.FormValue("name")
	number := r.FormValue("number")
	relation := r.FormValue("relation")
	address := r.FormValue("address")
	users, _ := r.Cookie("id")

	contact.Name = name
	contact.Number = number
	contact.Relationship = relation
	contact.Address = address
	contact.UserID = users.Value

	db.Save(&contact)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetContact(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := []models.Contact{}
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

func EditContact(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := models.Contact{}
	id := r.FormValue("id")
	name := r.FormValue("name")
	number := r.FormValue("number")
	relation := r.FormValue("relation")
	address := r.FormValue("address")

	db.Where("id = ?", id).Find(&contact)

	contact.Name = name
	contact.Number = number
	contact.Relationship = relation
	contact.Address = address

	db.Save(&contact)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func DeleteContact(w http.ResponseWriter, r *http.Request) {

	db := GormDB()
	id, _ := strconv.Atoi(r.FormValue("id"))
	contact := models.Contact{}

	db.Where("id", id).Statement.Delete(&contact)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GormDB() *gorm.DB {
	dsn := "root:a@tcp(127.0.0.1:3306)/mobile_tracker?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Faied to Connect to the Database ", err)
	}

	return db
}

func ReturnJSON(w http.ResponseWriter, r *http.Request, v interface{}) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		response := map[string]interface{}{
			"status":    "error",
			"error_msg": fmt.Sprintf("unable to encode JSON. %s", err),
		}
		b, _ = json.MarshalIndent(response, "", "  ")
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
		return
	}
	w.Write(b)
}
