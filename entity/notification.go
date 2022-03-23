package entity

type Notification struct {
	// gorm.Model
	IdNotif    int    `json:"id_notif"`
	Pesan      string `json:"pesan"`
	Status     int    `json:"status"`
	IdPeternak int    `json:"id_peternak"`
	IdCat      int    `json:"id_cat"`
	TglNotif   string `json:"tgl_notif"`
}
