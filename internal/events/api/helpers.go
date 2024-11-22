package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/InTeam-Russia/go-backend-template/internal/events/event"
)

func BuildEventFilter(r *http.Request) (*event.EventFilter, error) {
	query := r.URL.Query()
	filter := &event.EventFilter{}

	if sport := query.Get("sport"); sport != "" {
		filter.Sport = &sport
	}
	if team := query.Get("team"); team != "" {
		filter.Team = &team
	}
	if competitionType := query.Get("competitionType"); competitionType != "" {
		filter.CompetitionType = &competitionType
	}
	if genderAndAgeGroup := query.Get("genderAndAgeGroup"); genderAndAgeGroup != "" {
		filter.GenderAndAgeGroup = &genderAndAgeGroup
	}
	if programDiscipline := query.Get("programDiscipline"); programDiscipline != "" {
		filter.ProgramDiscipline = &programDiscipline
	}
	if country := query.Get("country"); country != "" {
		filter.Country = &country
	}
	if region := query.Get("region"); region != "" {
		filter.Region = &region
	}
	if city := query.Get("city"); city != "" {
		filter.City = &city
	}

	if minMembersCount := query.Get("minMembersCount"); minMembersCount != "" {
		if val, err := strconv.Atoi(minMembersCount); err == nil {
			filter.MinMembersCount = &val
		} else {
			return nil, err
		}
	}
	if maxMembersCount := query.Get("maxMembersCount"); maxMembersCount != "" {
		if val, err := strconv.Atoi(maxMembersCount); err == nil {
			filter.MaxMembersCount = &val
		} else {
			return nil, err
		}
	}

	if startDate := query.Get("startDate"); startDate != "" {
		if ts, err := strconv.ParseInt(startDate, 10, 64); err == nil {
			t := time.Unix(ts, 0)
			filter.StartDate = &t
		} else {
			return nil, err
		}
	}
	if endDate := query.Get("endDate"); endDate != "" {
		if ts, err := strconv.ParseInt(endDate, 10, 64); err == nil {
			t := time.Unix(ts, 0)
			filter.EndDate = &t
		} else {
			return nil, err
		}
	}

	return filter, nil
}
