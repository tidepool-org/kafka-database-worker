CREATE TABLE basal (
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

    delivery_type      TEXT NULL,
    duration           BIGINT,
    expected_duration  BIGINT,
    rate               DOUBLE PRECISION  NULL,
    percent            DOUBLE PRECISION  NULL,
    schedule_name      TEXT NULL
);

SELECT create_hypertable('basal', 'time');

