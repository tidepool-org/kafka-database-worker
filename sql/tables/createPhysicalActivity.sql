CREATE TABLE physical_activity (
    time                 TIMESTAMPTZ NOT NULL,

    archived_time        TIMESTAMPTZ NULL,
    created_time         TIMESTAMPTZ NULL,
    modified_time        TIMESTAMPTZ NULL,
    device_time          TIMESTAMPTZ NULL,

    device_id            TEXT NULL,
    id                   Text Null,
    internal_mongo_id    Text Null,
    guid                 Text Null,

    timezone             Text Null,
    timezone_offset      BIGINT NULL,
    clock_drift_offset   BIGINT NULL,
    conversion_offset    BIGINT NULL,

    upload_id            Text Null,
    user_id              Text Null,

    payload              jsonb Null,
    origin               jsonb Null,
    annotations          jsonb Null,

    active             boolean DEFAULT TRUE,

    revision             BIGINT Null,

    duration            jsonb Null,
    distance            jsonb Null,
    energy              jsonb Null,
    name                Text Null
);

SELECT create_hypertable('physical_activity', 'time');
