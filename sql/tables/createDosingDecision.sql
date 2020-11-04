CREATE TABLE dosing_decision (
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

    insulin_on_board                                   jsonb Null,
    carbohydrates_on_board                             jsonb Null,

    blood_glucose_target_schedule                      jsonb Null,
    blood_glucose_forecast                             jsonb Null,
    blood_glucose_forecast_including_pending_insulin   jsonb Null,

    recommended_basal                                  jsonb Null,
    recommended_bolus                                  jsonb Null,
    units                                              jsonb Null
);

SELECT create_hypertable('dosing_decision', 'time');

