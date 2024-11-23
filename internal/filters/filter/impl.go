package filter

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type PgFilterRepository struct {
	db     *pgxpool.Pool
	logger *zap.Logger
}

func NewPgFilterRepository(db *pgxpool.Pool, logger *zap.Logger) FilterRepository {
	return &PgFilterRepository{
		db:     db,
		logger: logger,
	}
}

func (r *PgFilterRepository) GetSports() ([]string, error) {
	r.logger.Debug("Getting sports from database")
	rows, err := r.db.Query(context.Background(), "SELECT name FROM sports")
	if err != nil {
		r.logger.Error("Failed to get sports", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	sports := make([]string, 0)
	for rows.Next() {
		var sport string
		if err := rows.Scan(&sport); err != nil {
			r.logger.Error("Failed to scan sport", zap.Error(err))
			return nil, err
		}
		sports = append(sports, sport)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Row iteration error", zap.Error(err))
		return nil, err
	}

	return sports, nil
}

func (r *PgFilterRepository) GetTeams(sport *string) ([]string, error) {
	r.logger.Debug("Getting teams", zap.Stringp("sport", sport))

	var query string
	var args []interface{}

	if sport == nil {
		query = "SELECT name FROM teams"
	} else {
		query = "SELECT name FROM teams WHERE sport_name = $1"
		args = append(args, *sport)
	}

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("Failed to get teams", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	teams := make([]string, 0)
	for rows.Next() {
		var team string
		if err := rows.Scan(&team); err != nil {
			r.logger.Error("Failed to scan team", zap.Error(err))
			return nil, err
		}
		teams = append(teams, team)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Row iteration error", zap.Error(err))
		return nil, err
	}

	return teams, nil
}

func (r *PgFilterRepository) GetCompetitionTypes(team *string) ([]string, error) {
	r.logger.Debug("Getting competition types", zap.Stringp("team", team))

	var query string
	var args []interface{}

	if team == nil {
		query = "SELECT DISTINCT name FROM competition_types"
	} else {
		query = "SELECT DISTINCT name FROM competition_types WHERE team_name = $1"
		args = append(args, *team)
	}

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("Failed to get competition types", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	competitionTypes := make([]string, 0)
	for rows.Next() {
		var competitionType string
		if err := rows.Scan(&competitionType); err != nil {
			r.logger.Error("Failed to scan competition type", zap.Error(err))
			return nil, err
		}
		competitionTypes = append(competitionTypes, competitionType)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Row iteration error", zap.Error(err))
		return nil, err
	}

	return competitionTypes, nil
}

func (r *PgFilterRepository) GetGenderAndAgeGroups(team *string) ([]string, error) {
	r.logger.Debug("Getting gender and age groups", zap.Stringp("team", team))

	var query string
	var args []interface{}

	if team == nil {
		query = "SELECT DISTINCT name FROM gender_and_age_groups"
	} else {
		query = `
			SELECT DISTINCT gag.name
			FROM teams_gender_and_age_groups tgag
			JOIN gender_and_age_groups gag ON tgag.gender_and_age_group_name = gag.name
			WHERE tgag.team_name = $1`
		args = append(args, *team)
	}

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("Failed to get gender and age groups", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	groups := make([]string, 0)
	for rows.Next() {
		var group string
		if err := rows.Scan(&group); err != nil {
			r.logger.Error("Failed to scan gender and age group", zap.Error(err))
			return nil, err
		}
		groups = append(groups, group)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Row iteration error", zap.Error(err))
		return nil, err
	}

	return groups, nil
}

func (r *PgFilterRepository) GetProgramsDisciplines(team *string) ([]string, error) {
	r.logger.Debug("Getting programs and disciplines", zap.Stringp("team", team))

	var query string
	var args []interface{}

	if team == nil {
		query = "SELECT DISTINCT name FROM programs_disciplines"
	} else {
		query = `
			SELECT DISTINCT pd.name
			FROM teams_programs_disciplines tpd
			JOIN programs_disciplines pd ON tpd.program_discipline_name = pd.name
			WHERE tpd.team_name = $1`
		args = append(args, *team)
	}

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("Failed to get programs and disciplines", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	programs := make([]string, 0)
	for rows.Next() {
		var program string
		if err := rows.Scan(&program); err != nil {
			r.logger.Error("Failed to scan program or discipline", zap.Error(err))
			return nil, err
		}
		programs = append(programs, program)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Row iteration error", zap.Error(err))
		return nil, err
	}

	return programs, nil
}

func (r *PgFilterRepository) GetCountries() ([]string, error) {
	r.logger.Debug("Getting countries from database")
	rows, err := r.db.Query(context.Background(), "SELECT country FROM countries")
	if err != nil {
		r.logger.Error("Failed to get countries", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	countries := make([]string, 0)
	for rows.Next() {
		var country string
		if err := rows.Scan(&country); err != nil {
			r.logger.Error("Failed to scan country", zap.Error(err))
			return nil, err
		}
		countries = append(countries, country)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Row iteration error", zap.Error(err))
		return nil, err
	}

	return countries, nil
}

func (r *PgFilterRepository) GetRegions(country *string) ([]string, error) {
	r.logger.Debug("Getting regions", zap.Stringp("country", country))

	var query string
	var args []interface{}

	if country == nil {
		query = "SELECT region FROM regions"
	} else {
		query = "SELECT region FROM regions WHERE country = $1"
		args = append(args, *country)
	}

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("Failed to get regions", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	regions := make([]string, 0)
	for rows.Next() {
		var region string
		if err := rows.Scan(&region); err != nil {
			r.logger.Error("Failed to scan region", zap.Error(err))
			return nil, err
		}
		regions = append(regions, region)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Row iteration error", zap.Error(err))
		return nil, err
	}

	return regions, nil
}

func (r *PgFilterRepository) GetCities(country *string, region *string) ([]string, error) {
	r.logger.Debug("Getting cities", zap.Stringp("country", country), zap.Stringp("region", region))

	var query string
	var args []interface{}

	if country == nil && region == nil {
		query = "SELECT city FROM cities"
	} else if country != nil && region == nil {
		query = "SELECT city FROM cities WHERE country = $1"
		args = append(args, *country)
	} else if country == nil && region != nil {
		query = "SELECT city FROM cities WHERE region = $1"
		args = append(args, *region)
	} else {
		query = "SELECT city FROM cities WHERE country = $1 AND region = $2"
		args = append(args, *country, *region)
	}

	rows, err := r.db.Query(context.Background(), query, args...)
	if err != nil {
		r.logger.Error("Failed to get cities", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	cities := make([]string, 0)
	for rows.Next() {
		var city string
		if err := rows.Scan(&city); err != nil {
			r.logger.Error("Failed to scan city", zap.Error(err))
			return nil, err
		}
		cities = append(cities, city)
	}

	if err := rows.Err(); err != nil {
		r.logger.Error("Row iteration error", zap.Error(err))
		return nil, err
	}

	return cities, nil
}
