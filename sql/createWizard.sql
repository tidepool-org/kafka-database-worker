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

    bg_input             DOUBLE PRECISION  NULL,
    carb_input           DOUBLE PRECISION  NULL,
    insulin_carb_input   DOUBLE PRECISION  NULL,

    bolus                Text Null,
    units                Text Null,

    recommended          jsonb NULL
);

SELECT create_hypertable('wizard', 'time');

