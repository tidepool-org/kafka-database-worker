CREATE TABLE upload (
    time                 TIMESTAMPTZ NULL,

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

    data_set_type        Text Null,
    data_state           Text Null,

    device_serial_number Text Null,
    state                Text Null,
    version              Text Null
);


