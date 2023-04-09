package api

import (
	"fmt"
	"net/http"

	"github.com/dafalo/Mobile-Tracker-App/models"
	"github.com/google/uuid"
)

func AddUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := models.User{}
	name := r.FormValue("name")
	number := r.FormValue("number")
	code := r.FormValue("code")

	contact.Name = name
	contact.Number = number
	contact.ENC = code

	db.Save(&contact)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetUser(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := []models.User{}
	db.Find(&contact)

	data := map[string]interface{}{
		"status": "ok",
		"item":   contact,
	}
	ReturnJSON(w, r, data)

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func Login(w http.ResponseWriter, r *http.Request) {

	db := GormDB()

	contact := models.User{}
	code := r.FormValue("code")
	newSession := uuid.NewString()

	db.Where("enc = ?", code).Find(&contact)

	http.SetCookie(w, &http.Cookie{
		Path:  "/",
		Name:  "session",
		Value: newSession,
	})

	http.SetCookie(w, &http.Cookie{
		Path:  "/",
		Name:  "id",
		Value: fmt.Sprint(contact.ID),
	})

	sqlDB, _ := db.DB()
	sqlDB.Close()

}

func GetActiveSession(r *http.Request) string {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	return ""
}
