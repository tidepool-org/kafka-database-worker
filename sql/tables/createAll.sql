CREATE TABLE cgm_settings (
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

                              transmitter_id       Text NULL,
                              units                Text NULL,

                              low_alerts           jsonb NULL,
                              high_alerts          jsonb NULL,
                              rate_of_change_alerts jsonb NULL,
                              out_of_range_alerts  jsonb NULL
);

SELECT create_hypertable('cgm_settings', 'time');

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

CREATE TABLE bolus (
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

    normal               DOUBLE PRECISION  NULL,

    sub_type            Text Null

);


SELECT create_hypertable('bolus', 'time');

CREATE TABLE cbg (
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

    value               DOUBLE PRECISION  NULL,

    units               Text Null
);

SELECT create_hypertable('cbg', 'time');

CREATE TABLE device_event (
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

    sub_type             Text Null,
    units                Text Null,

    value                DOUBLE PRECISION Null,

    duration             BIGINT Null,
    reason               Text Null,

    prime_target          Text Null,
    volume               DOUBLE PRECISION Null

);

SELECT create_hypertable('device_event', 'time');



CREATE TABLE device_meta (
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

    reason               jsonb Null,
    status               Text Null,
    sub_type             Text Null,

    duration             BIGINT Null
);

SELECT create_hypertable('device_meta', 'time');

CREATE TABLE food (
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

    nutrition            jsonb Null
);

SELECT create_hypertable('food', 'time');
CREATE TABLE physical_activity (
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

    duration            jsonb Null,
    distance            jsonb Null,
    energy              jsonb Null,
    name                Text Null
);

SELECT create_hypertable('physical_activity', 'time');
CREATE TABLE pump_settings (
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

    active_schedule      Text Null,
    basal_schedules      jsonb Null,
    bg_target            jsonb Null,
    carb_ratio           jsonb Null,
    insulin_sensitivity  jsonb Null,
    units                jsonb Null
);

SELECT create_hypertable('pump_settings', 'time');

CREATE TABLE smbg (
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

    value                DOUBLE PRECISION  NULL,

    units                Text Null

);

SELECT create_hypertable('smbg', 'time');

CREATE TABLE upload (
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

    data_set_type        Text Null,
    data_state           Text Null,

    device_serial_number Text Null,
    state                Text Null,
    version              Text Null
);


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

CREATE TABLE users (
                        user_id              TEXT,
                        username             TEXT,
                        authenticated        bool
);

CREATE TABLE clinics (
                         clinic_id     TEXT,
                         name          TEXT,
                         address       TEXT,
                         active        bool
);

CREATE TABLE clinics_clinicians (
                                    clinic_id     TEXT,
                                    clinician_id  TEXT,
                                    active        bool
);

CREATE TABLE clinics_patients (
                                  clinic_id     TEXT,
                                  patient_id    TEXT,
                                  active        bool
);