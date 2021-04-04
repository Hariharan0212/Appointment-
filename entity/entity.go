package entity

type Appointment struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Year int `json:"year"`
	Month int `json:"month"`
	Day  int `json:"day"`
	Hour int `json:"hour"`
	Minute int `json:"minute"`
	Second int `json:"second"`
	NanoSecond int `json:"nano_second"`
}