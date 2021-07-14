package types

type ImageHero struct {
	Url string `json:"url"`
}

type ConnectionsHero struct {
	GroupAffiliation string `json:"group-affiliation"`
	Relatives        string `json:"relatives"`
}

type WorkHero struct {
	Occupation string `json:"occupation"`
	Base       string `json:"base"`
}

type AppearanceHero struct {
	Gender    string   `json:"gender"`
	Race      string   `json:"race"`
	Height    []string `json:"height"`
	Weight    []string `json:"weight"`
	EyeColor  string   `json:"eyecolor"`
	HairColor string   `json:"haircolor"`
}

type BiographyHero struct {
	FullName        string   `json:"full-name"`
	AlterEgos       string   `json:"alter-egos"`
	Aliases         []string `json:"aliases"`
	PlaceOfBirth    string   `json:"place-of-birth"`
	FirstAppearance string   `json:"first-appearance"`
	Publisher       string   `json:"publisher"`
	Alignment       string   `json:"alignment"`
}

type PowerstatsHero struct {
	Intelligence string `json:"intelligence"`
	Strength     string `json:"strength"`
	Speed        string `json:"speed"`
	Durability   string `json:"durability"`
	Power        string `json:"power"`
	Combat       string `json:"combat"`
}

type Hero struct {
	Id          string          `json:"id"`
	Name        string          `json:"name"`
	Powerstats  PowerstatsHero  `json:"powerstats"`
	Biography   BiographyHero   `json:"biography"`
	Appearance  AppearanceHero  `json:"appearance"`
	Work        WorkHero        `json:"work"`
	Connections ConnectionsHero `json:"connections"`
	Image       ImageHero       `json:"image"`
}
