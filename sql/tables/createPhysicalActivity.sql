CREATE TABLE physical_activity (
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

    duration            jsonb Null,
    distance            jsonb Null,
    energy              jsonb Null,
    name                Text Null
);

SELECT create_hypertable('physical_activity', 'time');
