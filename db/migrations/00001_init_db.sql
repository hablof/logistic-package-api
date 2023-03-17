-- +goose Up
CREATE TABLE package (
  package_id BIGSERIAL PRIMARY KEY,
  title      VARCHAR(32) NOT NULL,
  material   VARCHAR(32) NOT NULL,
  max_volume REAL NOT NULL,
  reusable   BOOL DEFAULT FALSE NOT NULL,
  created_at TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TYPE event_type_enum AS ENUM ('Created', 'Updated', 'Removed');

CREATE TYPE event_status_enum AS ENUM ('Unlocked', 'Locked');

CREATE TABLE package_event (
  package_event_id BIGSERIAL PRIMARY KEY,
  package_id       BIGINT,
  event_type       event_type_enum NOT NULL,
  event_status     event_status_enum DEFAULT 'Unlocked',
  payload          JSONB,
  created_at       TIMESTAMP DEFAULT now()
);

-- +goose Down
DROP TABLE package;

DROP TABLE package_event;

DROP TYPE event_type_enum;

DROP TYPE event_status_enum;
