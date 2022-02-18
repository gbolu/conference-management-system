-- Add UUID extension
CREATE EXTENSION "uuid-ossp";

-- Set timezone
-- For more information, please visit:
-- https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
SET TIMEZONE="Africa/Lagos";

CREATE TYPE attendee_type AS ENUM ('speaker', 'participant');
CREATE TYPE edit_type AS ENUM ('conference', 'talk', 'attendee');

CREATE TABLE conferences(
  id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL,
  title VARCHAR(255) UNIQUE NOT NULL,
  description VARCHAR(255) NOT NULL,
  start_date TIMESTAMP NOT NULL,
  end_date TIMESTAMP NOT NULL
);

CREATE TABLE talks (
  id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL,
  title VARCHAR(255) UNIQUE NOT NULL,
  description VARCHAR(255) NOT NULL,
  duration INT NOT NULL,
  scheduled_date TIMESTAMP,
  conference_id UUID REFERENCES conferences(id) NOT NULL
);

CREATE TABLE attendees(
  id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL,
  username VARCHAR(255) UNIQUE NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  attendee_type attendee_type DEFAULT 'participant',
  talk_id UUID REFERENCES talks(id) NOT NULL 
);

CREATE TABLE edits (
  id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW (),
  updated_at TIMESTAMP NULL,
  previous_state JSONB NOT NULL,
  current_state JSONB NOT NULL,
  edit_type edit_type NOT NULL,
  edit_target_id UUID NOT NULL
);

CREATE INDEX idx_attendees_username ON attendees (LOWER(username));

CREATE INDEX idx_conferences_title ON conferences (LOWER(title));

CREATE INDEX idx_talks_title ON talks (LOWER(title));
