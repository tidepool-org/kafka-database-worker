CREATE TABLE pump_settings (
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

    active_schedule      Text Null,
    basal_schedules      jsonb Null,
    bg_target            jsonb Null,
    carb_ratio           jsonb Null,
    insulin_sensitivity  jsonb Null,
    units                jsonb Null
);

SELECT create_hypertable('pump_settings', 'time');

