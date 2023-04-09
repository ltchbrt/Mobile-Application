package api

import (
	"net/http"
	"strings"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "contact") {
		AddContact(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_contact") {
		GetContact(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_contact") {
		EditContact(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_contact") {
		DeleteContact(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "office") {
		AddOffice(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_office") {
		GetOffice(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "edit_office") {
		EditOffice(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "delete_office") {
		DeleteOffice(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "user") {
		AddUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "get_user") {
		GetUser(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "login") {
		Login(w, r)
		return
	}

	if strings.HasPrefix(r.URL.Path, "sms") {
		SendSMS(w, r)
		return
	}

}
