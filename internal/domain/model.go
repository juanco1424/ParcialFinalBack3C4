package domain

type Patient struct {
	Id            int    `json:"id"`
	Name          string `json:"name" binding:"required"`
	LastName      string `json:"last_name" binding:"required"`
	Address       string `json:"address" binding:"required"`
	DNI           string `json:"dni"  binding:"required"`
	DischargeDate string `json:"dischargeDate" `
}

type Dentist struct {
	ID           int    `json:"id"`
	LastName     string `json:"lastName"`
	Name         string `json:"Name"`
	Registration string `json:"Registration"`
}

type Appointment struct {
	ID          int     `json:"id"`
	Patient     Patient `json:"patient"`
	Dentist     Dentist `json:"dentist"`
	Date        string  `json:"date"`
	Hour        string  `json:"hour"`
	Description string  `json:"description"`
}
