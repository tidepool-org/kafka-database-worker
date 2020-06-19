CREATE TABLE bolus (
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

    normal               DOUBLE PRECISION  NULL,

    sub_type            Text Null,

    recommended          json NULL
);


SELECT create_hypertable('bolus', 'time');

