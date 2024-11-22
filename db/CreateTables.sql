CREATE TYPE user_role AS ENUM ('ADMIN', 'USER');

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    username TEXT UNIQUE NOT NULL,
    role user_role NOT NULL,
    password_hash BYTEA NOT NULL,
    password_salt BYTEA NOT NULL
);

CREATE TABLE sports (
    name TEXT PRIMARY KEY
);

CREATE TABLE teams (
    name TEXT NOT NULL,
    sport_name TEXT NOT NULL,
    PRIMARY KEY (name, sport_name),
    FOREIGN KEY (sport_name) REFERENCES sports(name)
);

CREATE TABLE gender_and_age_groups (
    name TEXT PRIMARY KEY
);

CREATE TABLE programs_disciplines (
    name TEXT PRIMARY KEY
);

CREATE TABLE countries (
    country TEXT PRIMARY KEY
);

CREATE TABLE regions (
    region TEXT PRIMARY KEY,
    country TEXT NOT NULL,
    FOREIGN KEY (country) REFERENCES countries(country)
);

CREATE TABLE cities (
    city TEXT PRIMARY KEY,
    region TEXT,
    country TEXT NOT NULL,
    FOREIGN KEY (region) REFERENCES regions(region),
    FOREIGN KEY (country) REFERENCES countries(country)
);

CREATE TABLE competitions (
    ekp_id BIGINT PRIMARY KEY,
    members_count INT,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL,

    country TEXT NOT NULL,
    region TEXT,
    city TEXT NOT NULL,
    sport_name TEXT NOT NULL,
    team_name TEXT NOT NULL,

    FOREIGN KEY (country) REFERENCES countries(country),
    FOREIGN KEY (region) REFERENCES regions(region),
    FOREIGN KEY (city) REFERENCES cities(city),
    FOREIGN KEY (sport_name) REFERENCES sports(name),
    FOREIGN KEY (sport_name, team_name) REFERENCES teams(sport_name, name)
);

CREATE TABLE competitions_programs_discipline (
    program_discipline_name TEXT NOT NULL,
    ekp_id BIGINT NOT NULL,

    PRIMARY KEY (program_discipline_name, ekp_id),
    FOREIGN KEY (program_discipline_name) REFERENCES programs_disciplines(name),
    FOREIGN KEY (ekp_id) REFERENCES competitions(ekp_id)
);

CREATE TABLE competitions_gender_and_age_groups (
    gender_and_age_group_name TEXT NOT NULL,
    ekp_id BIGINT NOT NULL,

    PRIMARY KEY (gender_and_age_group_name, ekp_id),
    FOREIGN KEY (gender_and_age_group_name) REFERENCES gender_and_age_groups(name),
    FOREIGN KEY (ekp_id) REFERENCES competitions(ekp_id)
);

CREATE TABLE competition_types (
    sport_name TEXT NOT NULL,
    team_name TEXT NOT NULL,
    name TEXT NOT NULL,

    PRIMARY KEY (name, sport_name, team_name),
    FOREIGN KEY (sport_name) REFERENCES sports(name),
    FOREIGN KEY (sport_name, team_name) REFERENCES teams(sport_name, name)
);

CREATE TABLE teams_gender_and_age_groups (
    gender_and_age_group_name TEXT NOT NULL,
    sport_name TEXT NOT NULL,
    team_name TEXT NOT NULL,

    PRIMARY KEY (gender_and_age_group_name, sport_name, team_name),
    FOREIGN KEY (gender_and_age_group_name) REFERENCES gender_and_age_groups(name),
    FOREIGN KEY (sport_name) REFERENCES sports(name),
    FOREIGN KEY (sport_name, team_name) REFERENCES teams(sport_name, name)
);

CREATE TABLE teams_programs_disciplines (
    program_discipline_name TEXT NOT NULL,
    sport_name TEXT NOT NULL,
    team_name TEXT NOT NULL,

    PRIMARY KEY (program_discipline_name, sport_name, team_name),
    FOREIGN KEY (program_discipline_name) REFERENCES programs_disciplines(name),
    FOREIGN KEY (sport_name) REFERENCES sports(name),
    FOREIGN KEY (sport_name, team_name) REFERENCES teams(sport_name, name)
);
