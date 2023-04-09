package api

import (
	"net/http"

	"github.com/dafalo/Mobile-Tracker-App/models"
)

func SendSMS(w http.ResponseWriter, r *http.Request) {
	db := GormDB()

	contact := []models.Contact{}
	office := []models.Office{}

	users, _ := r.Cookie("id")
	location := r.FormValue("location")
	time := r.FormValue("time")

	db.Preload("User").Where("user_id = ?", users.Value).Order("name").Find(&contact)
	db.Preload("User").Where("user_id = ?", users.Value).Order("name").Find(&office)

	for i := range contact {
		operator := models.SMS{}
		operator.Type = "SEMAPHORE"
		operator.SenderName = "ShareCafe"
		operator.Username = "dfbf5f79576de26479d6949a986a8465"

		operator.Message = contact[i].Name + " " + "Mobile Tracker's panic button has been triggered by" + " " + contact[i].User.Name + " " + "at" + " " + location + " " + time + "." + " " + "Please send assistance immediately."

		operator.Phone = contact[i].Number
		operator.Sendsms(false)

		println("Contacts", operator.Message)

		data := map[string]interface{}{
			"status": "ok",
			"item":   operator.Message,
		}
		ReturnJSON(w, r, data)
	}

	for j := range office {
		operator := models.SMS{}
		operator.Type = "SEMAPHORE"
		operator.SenderName = "ShareCafe"
		operator.Username = "dfbf5f79576de26479d6949a986a8465"

		operator.Message = office[j].Name + " " + "Mobile Tracker's panic button has been triggered by" + " " + contact[j].User.Name + " " + "at" + " " + location + " " + time + "." + " " + "Please send assistance immediately."

		operator.Phone = contact[j].Number
		operator.Sendsms(false)
		println("Offices", operator.Message)

	}

}
