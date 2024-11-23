package event

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PostgresEventRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewPgEventRepository(db *pgxpool.Pool, logger *zap.Logger) EventRepository {
	return &PostgresEventRepository{db, logger}
}

func (r *PostgresEventRepository) GetIndicatorsByFilter(filter *EventFilter) (*Indicators, error) {
	r.logger.Debug("Fetching indicators by filter", zap.Any("filter", filter))

	events, err := r.GetEventsByFilter(filter)
	if err != nil {
		r.logger.Error("Failed to fetch events by filter", zap.Any("filter", filter), zap.Error(err))
		return nil, err
	}

	indicators := calculateIndicators(events)
	r.logger.Debug("Calculated indicators", zap.Any("indicators", indicators))
	return indicators, nil
}

func (r *PostgresEventRepository) GetEventsByFilter(filter *EventFilter) ([]EventOut, error) {
	r.logger.Debug("Fetching events by filter", zap.Any("filter", filter))

	query := `
	SELECT
    competitions.ekp_id, -- Явное указание таблицы
    members_count,
    start_date,
    end_date,
    country,
    region,
    city,
    sport_name,
    team_name,
    ARRAY_REMOVE(ARRAY_AGG(DISTINCT gender_and_age_groups.name), NULL) AS gender_and_age_groups,
    ARRAY_REMOVE(ARRAY_AGG(DISTINCT programs_disciplines.name), NULL) AS programs_disciplines
	FROM competitions
	LEFT JOIN competitions_gender_and_age_groups
    ON competitions.ekp_id = competitions_gender_and_age_groups.ekp_id
	LEFT JOIN gender_and_age_groups
    ON gender_and_age_groups.name = competitions_gender_and_age_groups.gender_and_age_group_name
	LEFT JOIN competitions_programs_discipline
	  ON competitions.ekp_id = competitions_programs_discipline.ekp_id
	LEFT JOIN programs_disciplines
	  ON programs_disciplines.name = competitions_programs_discipline.program_discipline_name
	GROUP BY
		competitions.ekp_id
	HAVING 1 = 1
	`

	args := []interface{}{}
	argIndex := 1

	if filter.Sport != nil {
		query += fmt.Sprintf(" AND sport_name = $%d", argIndex)
		args = append(args, filter.Sport)
		argIndex++
	}
	if filter.Team != nil {
		query += fmt.Sprintf(" AND team_name = $%d", argIndex)
		args = append(args, filter.Team)
		argIndex++
	}
	if filter.Country != nil {
		query += fmt.Sprintf(" AND country = $%d", argIndex)
		args = append(args, filter.Country)
		argIndex++
	}
	if filter.Region != nil {
		query += fmt.Sprintf(" AND region = $%d", argIndex)
		args = append(args, filter.Region)
		argIndex++
	}
	if filter.City != nil {
		query += fmt.Sprintf(" AND city = $%d", argIndex)
		args = append(args, filter.City)
		argIndex++
	}
	if filter.StartDate != nil && filter.EndDate == nil {
		filter.EndDate = filter.StartDate
	} else if filter.EndDate != nil && filter.StartDate == nil {
		filter.StartDate = filter.EndDate
	}
	if filter.StartDate != nil && filter.EndDate != nil {
		query += fmt.Sprintf(" AND end_date >= $%d AND start_date <= $%d", argIndex, argIndex+1)
		args = append(args, filter.StartDate, filter.EndDate)
		argIndex += 2
	}
	if filter.MinMembersCount != nil {
		query += fmt.Sprintf(" AND members_count >= $%d", argIndex)
		args = append(args, filter.MinMembersCount)
		argIndex++
	}
	if filter.MaxMembersCount != nil {
		query += fmt.Sprintf(" AND members_count <= $%d", argIndex)
		args = append(args, filter.MaxMembersCount)
		argIndex++
	}
	if filter.GenderAndAgeGroup != nil {
		query += fmt.Sprintf(" AND $%d = ANY(ARRAY_AGG(gender_and_age_groups.name))", argIndex)
		args = append(args, filter.GenderAndAgeGroup)
		argIndex++
	}
	if filter.ProgramDiscipline != nil {
		query += fmt.Sprintf(" AND $%d = ANY(ARRAY_AGG(programs_disciplines.name))", argIndex)
		args = append(args, filter.ProgramDiscipline)
	}

	r.logger.Debug("Executing query", zap.String("query", query), zap.Any("args", args))

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("Query execution failed", zap.String("query", query), zap.Any("args", args), zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	events := make([]EventOut, 0)
	for rows.Next() {
		var event EventOut
		if err := rows.Scan(
			&event.EkpId, &event.MembersCount, &event.StartDate,
			&event.EndDate, &event.Venue.Country, &event.Venue.Region,
			&event.Venue.City, &event.Sport, &event.Team,
			&event.GenderAndAgeGroup, &event.ProgramDiscipline,
		); err != nil {
			r.logger.Error("Failed to scan row", zap.Error(err))
			return nil, err
		}
		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Error iterating rows", zap.Error(err))
		return nil, err
	}

	r.logger.Debug("Fetched events successfully", zap.Int("events_count", len(events)))
	return events, nil
}
