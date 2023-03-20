-- +goose Up
CREATE INDEX package_id_index ON package (package_id);

CREATE INDEX package_event_id_index ON package_event (package_event_id);
CREATE INDEX event_status_index     ON package_event (event_status);

-- +goose Down
DROP INDEX package_id_index;

DROP INDEX package_event_id_index;
DROP INDEX event_status_index;