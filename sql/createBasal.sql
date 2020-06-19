CREATE TABLE basal (
    time                 TIMESTAMPTZ NOT NULL,

    created_time         TIMESTAMPTZ NULL,
    modified_time        TIMESTAMPTZ NULL,
    device_time          TIMESTAMPTZ NULL,

    device_id            TEXT NULL,
    id                   Text Null,

    timezone             Text Null,
    timezone_offset      INT NULL,
    clock_drift_offset   INT NULL,
    conversion_offset    INT NULL,

    upload_id            Text Null,
    user_id              Text Null,

    revision             INT Null,

    delivery_type      TEXT NULL,
    duration           INT,
    expected_duration  INT,
    rate               DOUBLE PRECISION  NULL,
    percent            DOUBLE PRECISION  NULL,
    schedule_name      TEXT NULL
);

SELECT create_hypertable('basal', 'time');

