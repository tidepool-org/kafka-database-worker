CREATE TABLE bolus (
    time                 TIMESTAMPTZ NOT NULL,

    archived_time        TIMESTAMPTZ NULL,
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

    payload              Text Null,
    origin               Text Null,

    active             boolean DEFAULT TRUE,

    revision             BIGINT Null,

    normal               DOUBLE PRECISION  NULL,

    sub_type            Text Null

);


SELECT create_hypertable('bolus', 'time');

