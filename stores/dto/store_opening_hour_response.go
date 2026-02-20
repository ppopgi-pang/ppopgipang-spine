package dto

type StoreOpeningHourResponse struct {
	ID        int64  `json:"id"`
	DayOfWeek *int8  `json:"day_of_week"`
	OpenTime  string `json:"open_time"`
	CloseTime string `json:"close_time"`
	IsClosed  bool   `json:"is_closed"`
}
