-- settings is saving the settings. this should ONLY ever have one row

-- +goose Up
CREATE TABLE settings(
    settings_complete BOOLEAN NOT NULL DEFAULT false,
    only_activate_self BOOLEAN NOT NULL DEFAULT false
);

-- +goose Down
DROP TABLE settings;