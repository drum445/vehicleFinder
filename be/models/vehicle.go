package models

type Vehicle struct {
	ID           int    `json:"id" bson:"_id"`
	Make         string `json:"make" bson:"make"`
	ShortModel   string `json:"shortModel" bson:"shortModel"`
	LongModel    string `json:"longModel" bson:"longModel"`
	Trim         string `json:"trim" bson:"trim"`
	Derivative   string `json:"derivative" bson:"derivative"`
	Introduced   string `json:"introduced" bson:"introduced"`
	Discontinued string `json:"discontinued" bson:"discontinued"`
	Available    string `json:"available" bson:"available"`
	Image        string `json:"image,omitempty" bson:"image,omitempty"`
}

type Vehicles []Vehicle
