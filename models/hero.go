package models

type Hero struct {
	Id                string   `json:"id"`
	Name              string   `json:"name"`
	FullName          string   `json:"fullName"`
	Intelligence      string   `json:"intelligence"`
	Power             string   `json:"power"`
	Occupation        string   `json:"occupation"`
	ImageUrl          string   `json:"imageUrl"`
	GroupAffiliation  []string `json:"groupAffiliation"`
	NumberOfRelatives int      `json:"numberOfRelatives"`
	Alignment         string   `json:"alignment"`
}
