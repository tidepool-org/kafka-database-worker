CREATE TABLE device_event (
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

    payload              jsonb Null,
    origin               jsonb Null,

    active             boolean DEFAULT TRUE,

    revision             BIGINT Null,

    sub_type             Text Null,
    units                Text Null,

    value                DOUBLE PRECISION Null,

    duration             BIGINT Null,
    reason               jsonb Null,

    prime_target          Text Null,
    volume               DOUBLE PRECISION Null

);

SELECT create_hypertable('device_event', 'time');



