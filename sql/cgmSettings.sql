CREATE TABLE wizard (
    time                 TIMESTAMPTZ NOT NULL,

    created_time         TIMESTAMPTZ NULL,
    modified_time        TIMESTAMPTZ NULL,
    device_time          TIMESTAMPTZ NULL,

    device_id            TEXT NULL,
    id                   Text Null,

    timezone             Text Null,
    timezone_offset      BIGINT NULL,
    clock_drift_offset   BIGINT NULL,
    conversion_offset    BIGINT NULL,

    upload_id            Text Null,
    user_id              Text Null,

    revision             BIGINT Null,

    transmitterId        Text NULL,
    units                Text NULL,

    low_alerts           jsonb NULL,
    high_alerts          jsonb NULL,
    rate_of_change_alerts jsonb NULL,
    out_of_range_alerts  jsonb NULL,
);

SELECT create_hypertable('wizard', 'time');

