package views

import (
	"net/http"
	"text/template"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	activeSession := GetActiveSession(r)

	if activeSession == "" {
		tmpl := template.Must(template.ParseFiles("./templates/login.html"))
		data := map[string]interface{}{}
		data["Title"] = "Register | POS_SYSTEM"
		tmpl.Execute(w, data)

	} else {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	}
}

func GetActiveSession(r *http.Request) string {
	key, err := r.Cookie("session")
	if err == nil && key != nil {
		return key.Value
	}
	return ""
}
