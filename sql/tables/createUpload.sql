CREATE TABLE upload (
    time                 TIMESTAMPTZ NULL,

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

    data_set_type        Text Null,
    data_state           Text Null,

    device_manufacturers  Text[],
    device_model          Text Null,
    device_serial_number Text Null,
    device_tags          Text[],

    state                Text Null,
    version              Text Null
);


