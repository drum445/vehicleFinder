package models

// need to declare the db property where the column name in the db
// differens from our objects attribute
type Vehicle struct {
	ID           int    `json:"id" db:"vehicle_id"`
	Make         string `json:"make"`
	ShortModel   string `json:"shortModel" db:"short_model"`
	LongModel    string `json:"longModel"  db:"long_model"`
	Trim         string `json:"trim"`
	Derivative   string `json:"derivative"`
	Introduced   string `json:"introduced"`
	Discontinued string `json:"discontinued"`
	Available    string `json:"available"`
	Image        string `json:"image,omitempty"`
}

type Vehicles []Vehicle
