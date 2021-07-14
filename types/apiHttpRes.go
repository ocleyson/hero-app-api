package types

type GetHeroRes struct {
	Response string `json:"response"`
	Hero
}

type SearchHeroesRes struct {
	Response   string `json:"response"`
	ResultsFor string `json:"results-for"`
	Results    []Hero `json:"results"`
}
