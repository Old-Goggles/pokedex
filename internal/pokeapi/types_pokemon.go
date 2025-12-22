package pokeapi

type Pokemon struct {
	Name           string        `json:"name"`
	BaseExperience int           `json:"base_experience"`
	Height         int           `json:"height"`
	Weight         int           `json:"weight"`
	Stats          []PokemonStat `json:"stats"`
	Types          []PokemonType `json:"types"`
}

type PokemonStat struct {
	BaseStat int          `json:"base_stat"`
	Stat     NamedAPIStat `json:"stat"`
}

type NamedAPIStat struct {
	Name string `json:"name"`
}

type PokemonType struct {
	Type NamedAPIResource `json:"type"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
}
