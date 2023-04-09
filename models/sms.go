package models

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// Operator !
type SMS struct {
	Type       string
	SenderName string
	Username   string
	Phone      string
	Message    string
	From       string
	Error      string `sql:"type:text"`
	Response   string `sql:"type:text"`
}

// Sendsms !
func (o *SMS) Sendsms(priority bool) {
	if o.Type == "SEMAPHORE" {
		// log.Println("semaphore send to:", o.Phone)
		// log.Println("Message:", o.Message)
		client := &http.Client{}

		URL := "https://api.semaphore.co/api/v4/messages"
		if priority {
			URL = "http://api.semaphore.co/api/v4/priority"
		}

		v := url.Values{}
		v.Set("message", o.Message)

		if !strings.HasPrefix(o.Phone, "+") {
			if strings.HasPrefix(o.Phone, "09") {
				o.Phone = strings.TrimPrefix(o.Phone, "0")
				o.Phone = "+63" + o.Phone

			} else {
				log.Println("ERROR: SMS.Sendsms Unknown Cellphone number Format", o.Phone)
				return
			}
		}

		v.Set("number", o.Phone)
		v.Set("apikey", o.Username) //"dfbf5f79576de26479d6949a986a8465")
		v.Set("sendername", o.SenderName)

		req, err := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		if err != nil {
			log.Println("ERROR: SMS.Sendsms.http.NewRequest", err)
			o.Error = err.Error()
		}

		resp, err := client.Do(req)
		if err != nil {
			log.Println("ERROR: SMS.Sendsms.Client.Do", err)
			o.Error = err.Error()
			return
		}

		bodyText, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("ERROR: SMS.Sendsms.ioutil.ReadAll", err)
			o.Error = err.Error()
		}

		s := string(bodyText)
		log.Println("INFO: Message Sent:", s)

		o.Response += s
		return
	}
}
