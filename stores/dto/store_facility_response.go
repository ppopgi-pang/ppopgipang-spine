package dto

type StoreFacilityResponse struct {
	MachineCount   *int     `json:"machine_count"`
	PaymentMethods []string `json:"payment_methods"`
}
