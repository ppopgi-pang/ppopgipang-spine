package dto

type StoreOpeningHourResponse struct {
	ID        int64
	DayOfWeek *int8
	OpenTime  string
	CloseTime string
	IsClosed  bool
}
