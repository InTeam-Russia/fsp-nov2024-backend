package event

import (
	"strings"
	"time"

	"github.com/InTeam-Russia/go-backend-template/internal/helpers"
)

var MockEvents = []EventOut{
	{
		EkpId:             2152500017017152,
		Sport:             "АВИАМОДЕЛЬНЫЙ СПОРТ",
		Team:              "Основной состав",
		CompetitionType:   "ЧЕМПИОНАТ ЦЕНТРАЛЬНОГО ФЕДЕРАЛЬНОГО ОКРУГА",
		GenderAndAgeGroup: []string{"женщины", "мужчины от 14 лет и старше"},
		ProgramDiscipline: []string{"КЛАСС F-1D", "дисциплина F-1DХУЙ ЗНАЕТ"},
		Venue: Venue{
			Country: "РОССИЯ",
			Region:  nil,
			City:    "Корнер-кейсово",
		},
		MembersCount: 25,
		StartDate:    time.Date(2024, 2, 16, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 2, 19, 0, 0, 0, 0, time.UTC),
	},
	{
		EkpId:             2152020021017156,
		Sport:             "АВИАМОДЕЛЬНЫЙ СПОРТ",
		Team:              "Основной состав",
		CompetitionType:   "ВСЕРОССИЙСКИЕ СОРЕВНОВАНИЯ",
		GenderAndAgeGroup: []string{"женщины", "мужчины от 14 лет и старше"},
		ProgramDiscipline: []string{"КЛАСС F-1E", "дисциплина F-1Е"},
		Venue: Venue{
			Country: "РОССИЯ",
			Region:  helpers.StrPtr("РЕСПУБЛИКА БАШКОРТОСТАН"),
			City:    "г. Уфа",
		},
		MembersCount: 40,
		StartDate:    time.Date(2024, 2, 16, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 2, 19, 0, 0, 0, 0, time.UTC),
	},
	{
		EkpId:             2152360017017163,
		Sport:             "АВИАМОДЕЛЬНЫЙ СПОРТ",
		Team:              "Основной состав",
		CompetitionType:   "ЧЕМПИОНАТ ЦЕНТРАЛЬНОГО ФЕДЕРАЛЬНОГО ОКРУГА",
		GenderAndAgeGroup: []string{"женщины", "мужчины от 14 лет и старше"},
		ProgramDiscipline: []string{"КЛАСС F-1E", "дисциплина F-1E"},
		Venue: Venue{
			Country: "РОССИЯ",
			Region:  helpers.StrPtr("ВОРОНЕЖСКАЯ ОБЛАСТЬ"),
			City:    "г. Воронеж",
		},
		MembersCount: 25,
		StartDate:    time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 2, 25, 0, 0, 0, 0, time.UTC),
	},
	{
		EkpId:             2152580021017165,
		Sport:             "АВИАМОДЕЛЬНЫЙ СПОРТ",
		Team:              "Основной состав",
		CompetitionType:   "ВСЕРОССИЙСКИЕ СОРЕВНОВАНИЯ",
		GenderAndAgeGroup: []string{"женщины", "мужчины от 14 лет и старше"},
		ProgramDiscipline: []string{"КЛАСС F-2B", "дисциплина F-2В"},
		Venue: Venue{
			Country: "РОССИЯ",
			Region:  helpers.StrPtr("ПЕНЗЕНСКАЯ ОБЛАСТЬ"),
			City:    "г. Пенза",
		},
		MembersCount: 40,
		StartDate:    time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 2, 25, 0, 0, 0, 0, time.UTC),
	},
	{
		EkpId:             2152670017022956,
		Sport:             "АВИАМОДЕЛЬНЫЙ СПОРТ",
		Team:              "Основной состав",
		CompetitionType:   "ЧЕМПИОНАТ ЦЕНТРАЛЬНОГО ФЕДЕРАЛЬНОГО ОКРУГА",
		GenderAndAgeGroup: []string{"женщины", "мужчины от 14 лет и старше"},
		ProgramDiscipline: []string{"КЛАСС F-2D", "дисциплина F-2D"},
		Venue: Venue{
			Country: "РОССИЯ",
			Region:  helpers.StrPtr("СМОЛЕНСКАЯ ОБЛАСТЬ"),
			City:    "г. Смоленск",
		},
		MembersCount: 30,
		StartDate:    time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 2, 26, 0, 0, 0, 0, time.UTC),
	},
	{
		EkpId:             2152640021022918,
		Sport:             "АВИАМОДЕЛЬНЫЙ СПОРТ",
		Team:              "Основной состав",
		CompetitionType:   "ВСЕРОССИЙСКИЕ СОРЕВНОВАНИЯ",
		GenderAndAgeGroup: []string{"женщины", "мужчины от 14 лет и старше"},
		ProgramDiscipline: []string{"КЛАСС F-2D", "дисциплина F-2D"},
		Venue: Venue{
			Country: "РОССИЯ",
			Region:  helpers.StrPtr("САРАТОВСКАЯ ОБЛАСТЬ"),
			City:    "Дубки поселок",
		},
		MembersCount: 60,
		StartDate:    time.Date(2024, 3, 22, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 3, 24, 0, 0, 0, 0, time.UTC),
	},
	{
		EkpId:             2152350021022921,
		Sport:             "АВИАМОДЕЛЬНЫЙ СПОРТ",
		Team:              "Основной состав",
		CompetitionType:   "ВСЕРОССИЙСКИЕ СОРЕВНОВАНИЯ",
		GenderAndAgeGroup: []string{"женщины", "мужчины от 14 лет и старше"},
		ProgramDiscipline: []string{"КЛАСС F-2D", "дисциплина F-2D"},
		Venue: Venue{
			Country: "РОССИЯ",
			Region:  helpers.StrPtr("ВОЛОГОДСКАЯ ОБЛАСТЬ"),
			City:    "г. Вологда",
		},
		MembersCount: 60,
		StartDate:    time.Date(2024, 3, 29, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		EkpId:             2152630019022857,
		Sport:             "АВИАМОДЕЛЬНЫЙ СПОРТ",
		Team:              "Основной состав",
		CompetitionType:   "ЧЕМПИОНАТ РОССИИ",
		GenderAndAgeGroup: []string{"женщины", "мужчины"},
		ProgramDiscipline: []string{"КЛАСС F-1E", "дисциплина F-1E"},
		Venue: Venue{
			Country: "РОССИЯ",
			Region:  helpers.StrPtr("САМАРСКАЯ ОБЛАСТЬ"),
			City:    "г. Тольятти",
		},
		MembersCount: 50,
		StartDate:    time.Date(2024, 3, 30, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		EkpId:             2152520023022234,
		Sport:             "АВИАМОДЕЛЬНЫЙ СПОРТ",
		Team:              "Основной состав",
		CompetitionType:   "ЧЕМПИОНАТ МОСКОВСКОЙ ОБЛАСТИ",
		GenderAndAgeGroup: []string{"женщины", "мужчины от 14 лет и старше"},
		ProgramDiscipline: []string{"КЛАСС F-2E", "дисциплина F-2E"},
		Venue: Venue{
			Country: "РОССИЯ",
			Region:  helpers.StrPtr("МОСКОВСКАЯ ОБЛАСТЬ"),
			City:    "г. Подольск",
		},
		MembersCount: 30,
		StartDate:    time.Date(2024, 4, 5, 0, 0, 0, 0, time.UTC),
		EndDate:      time.Date(2024, 4, 7, 0, 0, 0, 0, time.UTC),
	},
}

type MockEventRepository struct {
	events []EventOut
}

func NewMockEventRepository() EventRepository {
	return &MockEventRepository{events: MockEvents}
}

func matchesFilter(value *string, target string) bool {
	return value == nil || strings.EqualFold(*value, target)
}

func (repo *MockEventRepository) GetEventsByFilter(filter *EventFilter) ([]EventOut, error) {
	filteredEvents := make([]EventOut, 0)

	for _, event := range repo.events {
		if !matchesFilter(filter.Sport, event.Sport) {
			continue
		}
		if !matchesFilter(filter.Team, event.Team) {
			continue
		}
		if !matchesFilter(filter.CompetitionType, event.CompetitionType) {
			continue
		}
		if filter.GenderAndAgeGroup != nil && !helpers.Contains(event.GenderAndAgeGroup, *filter.GenderAndAgeGroup) {
			continue
		}
		if filter.ProgramDiscipline != nil && !helpers.Contains(event.ProgramDiscipline, *filter.ProgramDiscipline) {
			continue
		}
		if !matchesFilter(filter.Country, event.Venue.Country) {
			continue
		}
		if filter.Region != nil && (event.Venue.Region == nil || *filter.Region != *event.Venue.Region) {
			continue
		}
		if !matchesFilter(filter.City, event.Venue.City) {
			continue
		}
		if filter.MinMembersCount != nil && event.MembersCount < *filter.MinMembersCount {
			continue
		}
		if filter.MaxMembersCount != nil && event.MembersCount > *filter.MaxMembersCount {
			continue
		}
		if filter.StartDate != nil && event.EndDate.Before(*filter.StartDate) {
			continue
		}
		if filter.EndDate != nil && event.StartDate.After(*filter.EndDate) {
			continue
		}
		filteredEvents = append(filteredEvents, event)
	}

	return filteredEvents, nil
}

func (repo *MockEventRepository) GetIndicatorsByFilter(filter *EventFilter) (*Indicators, error) {
	events, err := repo.GetEventsByFilter(filter)
	if err != nil {
		return nil, err
	}
	return calculateIndicators(events), nil
}
