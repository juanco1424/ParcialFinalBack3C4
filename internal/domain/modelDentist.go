package domain

type Dentist struct {
	ID           int    `json:"id"`
	LastName     string `json:"lastName"`
	Name         string `json:"Name"`
	Registration string `json:"Registration"`
}
