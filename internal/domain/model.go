package domain

type Patient struct {
	Id            string `json:"id"`
	Name          string `json:"name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	DNI           string `json:"dni"  binding:"required"`
	DischargeDate string `json:"discharge_date" `
}
