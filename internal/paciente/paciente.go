package paciente

type Pacient struct {
	Name          string `json:"name"`
	LastName      string `json:"last_name"`
	Address       string `json:"address"`
	DNI           string `json:"dni"`
	DischargeDate string `json:"discharge_date"`
}
