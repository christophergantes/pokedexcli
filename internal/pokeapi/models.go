package pokeapi

type (
	Encounter struct {
		MinLevel        int                `json:"min_level"`
		MaxLevel        int                `json:"max_level"`
		ConditionValues []NamedAPIResource `json:"condition_values"`
		Chance          int                `json:"chance"`
		Method          NamedAPIResource   `json:"method"`
	}

	EncounterCondition struct {
		ID     int                `json:"id"`
		Name   string             `json:"name"`
		Names  []string           `json:"names"`
		Values []NamedAPIResource `json:"values"`
	}

	EncounterConditionValue struct {
		ID        int              `json:"id"`
		Name      string           `json:"name"`
		Condition NamedAPIResource `json:"condition"`
		Names     []Name           `json:"names"`
	}

	EncounterMethod struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		URL   string `json:"url"`
		Names []Name `json:"names"`
	}

	EncounterMethodRate struct {
		EncounterMethod NamedAPIResource          `json:"encounter_method"`
		VersionDetails  []EncounterVersionDetails `json:"version_details"`
	}

	EncounterVersionDetails struct {
		Rate    int              `json:"rate"`
		Version NamedAPIResource `json:"version"`
	}

	LocationArea struct {
		ID                   int                   `json:"id"`
		Name                 string                `json:"name"`
		GameIndex            int                   `json:"game_index"`
		EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
		Location             NamedAPIResource      `json:"location"`
		Names                []Name                `json:"names"`
		PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
	}

	NamedAPIResource struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	NamedAPIResourceList struct {
		Count    int                `json:"count"`
		Next     *string            `json:"next"`
		Previous *string            `json:"previous"`
		Results  []NamedAPIResource `json:"results"`
	}

	PokemonEncounter struct {
		Pokemon        NamedAPIResource         `json:"pokemon"`
		VersionDetails []VersionEncounterDetail `json:"version_details"`
	}

	Name struct {
		Name     string           `json:"name"`
		Language NamedAPIResource `json:"language"`
	}

	VersionEncounterDetail struct {
		Version          NamedAPIResource `json:"version"`
		MaxChance        int              `json:"max_chance"`
		EncounterDetails []Encounter      `json:"encounter_details"`
	}
)
