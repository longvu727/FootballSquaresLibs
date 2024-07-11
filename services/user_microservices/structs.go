package userservices

type User struct {
	UserID     int    `json:"user_id"`
	UserGUID   string `json:"user_guid"`
	IP         string `json:"ip"`
	DeviceName string `json:"device_name"`
	UserName   string `json:"user_name"`
	Alias      string `json:"alias"`
}
