package filter

import "github.com/InTeam-Russia/go-backend-template/internal/events"

type MockFilterRepository struct{}

func (r *MockFilterRepository) GetSports() ([]string, error) {
	set := make(map[string]struct{})
	for _, event := range events.MockEvents {
		set[event.Sport] = struct{}{}
	}
	return setToSlice(set), nil
}

func (r *MockFilterRepository) GetTeams(sport *string) ([]string, error) {
	set := make(map[string]struct{})
	for _, event := range events.MockEvents {
		if sport == nil || event.Sport == *sport {
			set[event.Team] = struct{}{}
		}
	}
	return setToSlice(set), nil
}

func (r *MockFilterRepository) GetCompetitionTypes(team *string) ([]string, error) {
	set := make(map[string]struct{})
	for _, event := range events.MockEvents {
		if team == nil || event.Team == *team {
			set[event.CompetitionType] = struct{}{}
		}
	}
	return setToSlice(set), nil
}

func (r *MockFilterRepository) GetGenderAndAgeGroups(team *string) ([]string, error) {
	set := make(map[string]struct{})
	for _, event := range events.MockEvents {
		if team == nil || event.Team == *team {
			for _, group := range event.GenderAndAgeGroup {
				set[group] = struct{}{}
			}
		}
	}
	return setToSlice(set), nil
}

func (r *MockFilterRepository) GetProgramsDisciplines(team *string) ([]string, error) {
	set := make(map[string]struct{})
	for _, event := range events.MockEvents {
		if team == nil || event.Team == *team {
			for _, discipline := range event.ProgramDiscipline {
				set[discipline] = struct{}{}
			}
		}
	}
	return setToSlice(set), nil
}

func (r *MockFilterRepository) GetCountries() ([]string, error) {
	set := make(map[string]struct{})
	for _, event := range events.MockEvents {
		set[event.Venue.Country] = struct{}{}
	}
	return setToSlice(set), nil
}

func (r *MockFilterRepository) GetRegions(country *string) ([]string, error) {
	set := make(map[string]struct{})
	for _, event := range events.MockEvents {
		if country == nil || event.Venue.Country == *country {
			if event.Venue.Region != nil {
				set[*event.Venue.Region] = struct{}{}
			}
		}
	}
	return setToSlice(set), nil
}

func (r *MockFilterRepository) GetCities(country *string, region *string) ([]string, error) {
	set := make(map[string]struct{})
	for _, event := range events.MockEvents {
		if (country == nil || event.Venue.Country == *country) &&
			(region == nil || (event.Venue.Region != nil && *event.Venue.Region == *region)) {
			set[event.Venue.City] = struct{}{}
		}
	}
	return setToSlice(set), nil
}

func setToSlice(set map[string]struct{}) []string {
	result := make([]string, 0, len(set))
	for key := range set {
		result = append(result, key)
	}
	return result
}
