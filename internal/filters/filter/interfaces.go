package filter

type FilterRepository interface {
	GetSports() ([]string, error)
	GetTeams(sport *string) ([]string, error)
	GetCompetitionTypes(team *string) ([]string, error)
	GetGenderAndAgeGroups(team *string) ([]string, error)
	GetProgramsDisciplines(team *string) ([]string, error)
	GetCountries() ([]string, error)
	GetRegions(country *string) ([]string, error)
	GetCities(country *string, region *string) ([]string, error)
}
