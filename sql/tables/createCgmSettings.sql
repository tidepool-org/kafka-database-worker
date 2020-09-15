CREATE TABLE cgm_settings (
    time                 TIMESTAMPTZ NOT NULL,

    archived_time        TIMESTAMPTZ NULL,
    created_time         TIMESTAMPTZ NULL,
    modified_time        TIMESTAMPTZ NULL,
    device_time          TIMESTAMPTZ NULL,

    device_id            TEXT NULL,
    id                   Text Null,
    guid                 Text Null,

    timezone             Text Null,
    timezone_offset      BIGINT NULL,
    clock_drift_offset   BIGINT NULL,
    conversion_offset    BIGINT NULL,

    upload_id            Text Null,
    user_id              Text Null,

    payload              Text Null,
    origin               Text Null,

    active             boolean DEFAULT TRUE,

    revision             BIGINT Null,

    transmitter_id       Text NULL,
    units                Text NULL,

    low_alerts           jsonb NULL,
    high_alerts          jsonb NULL,
    rate_of_change_alerts jsonb NULL,
    out_of_range_alerts  jsonb NULL
);

SELECT create_hypertable('cgm_settings', 'time');

