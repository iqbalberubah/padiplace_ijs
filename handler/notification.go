package handler

// import (
// 	"encoding/json"
// 	"io/ioutil"
// 	"net/http"
// 	d "padiplace_ijs/database"
// 	e "padiplace_ijs/entity"
// )

// func GetNotification(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var notification []e.Notification
// 	var body map[string]interface{}
// 	buffer, _ := ioutil.ReadAll(r.Body)
// 	if err := json.Unmarshal(buffer, &body); err != nil {
// 		json.NewEncoder(w).Encode(e.ErrorResponse{404, "Request body tidak sesuai"})
// 		return
// 	}
// 	d.DB.Table("notifikasi_fcm").Where("id_kios = ?", body["id_kios"]).Find(&notification)
// 	json.NewEncoder(w).Encode(e.SuccesResponse{0, "Succes", notification})
// }
