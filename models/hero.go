package models

import (
	"github.com/lib/pq"
)

type Hero struct {
	Id                string         `json:"id"`
	Name              string         `json:"name"`
	FullName          string         `json:"fullName"`
	Intelligence      string         `json:"intelligence"`
	Power             string         `json:"power"`
	Occupation        string         `json:"occupation"`
	ImageUrl          string         `json:"imageUrl"`
	GroupAffiliation  pq.StringArray `gorm:"type:text[]" json:"groupAffiliation"`
	NumberOfRelatives int            `json:"numberOfRelatives"`
	Alignment         string         `json:"alignment"`
}
