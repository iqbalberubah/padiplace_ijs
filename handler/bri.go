package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	d "padiplace_ijs/database"
	e "padiplace_ijs/entity"
	"strings"
	"time"
)

func GetToken(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// var user e.User
	// var body map[string]interface{}
	// buffer, _ := ioutil.ReadAll(r.Body)
	// if err := json.Unmarshal(buffer, &body); err != nil {
	// 	json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
	// 	return
	// }
	// if result := d.DB.Table("user").Where("username = ? OR tlp_user = ? OR email_user = ?", body["username"], body["username"], body["username"]).First(&user); result.Error != nil {
	// 	json.NewEncoder(w).Encode(e.ErrorResponse{404, "User tidak ditemukan"})
	// } else {
	// 	if user.IsActive == 1 {
	// 		match := PasswordVerify(body["password1_user"].(string), user.Password1User)
	// 		if match {
	// 			json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", user})
	// 		} else {
	// 			json.NewEncoder(w).Encode(e.ErrorResponse{404, "Password tidak sesuai"})
	// 		}
	// 	} else {
	// 		json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", user})
	// 	}
	// }

	var credential e.Credential
	if result := d.DB.Table("bri_credential").Where("name = ?", "access").First(&credential); result.Error != nil {
		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Credential tidak ditemukan"})
	}

	date, _ := time.Parse("yyyy-MM-dd", credential.Date)
	if date.Before(time.Now()) || credential.Value3 == "" {
		apiUrl := "https://sandbox.partner.api.bri.co.id/oauth/client_credential/accesstoken?grant_type=client_credentials"
		data := url.Values{}
		data.Set("client_id", "sZiRlD0cxfr4njgRV080GybGfQcmAXck")
		data.Set("client_secret", "ThFGrh7t31zDPF3L")

		u, _ := url.ParseRequestURI(apiUrl)
		urlStr := u.String()

		client := &http.Client{}
		r, err := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
		if err != nil {

		}
		r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

		resp, _ := client.Do(r)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			var body map[string]interface{}
			buffer, _ := ioutil.ReadAll(resp.Body)
			if err := json.Unmarshal(buffer, &body); err != nil {
				json.NewEncoder(w).Encode(e.ErrorResponse{404, "Response body tidak sesuai"})
				return
			}
			time.Now()
			credential.Date =
				d.DB.Table("bri_credential").Where("name = ?", "access").Save(&credential)
			json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", body})
		} else {
			json.NewEncoder(w).Encode(e.ErrorResponse{404, "Error"})
		}
	}
}
