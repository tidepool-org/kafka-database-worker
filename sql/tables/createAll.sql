CREATE TABLE basal
(
    time               TIMESTAMPTZ      NOT NULL,

    archived_time      TIMESTAMPTZ      NULL,
    created_time       TIMESTAMPTZ      NULL,
    modified_time      TIMESTAMPTZ      NULL,
    device_time        TIMESTAMPTZ      NULL,

    device_id          TEXT             NULL,
    id                 Text             Null,
    internal_mongo_id  Text             Null,
    guid               Text             Null,

    timezone           Text             Null,
    timezone_offset    BIGINT           NULL,
    clock_drift_offset BIGINT           NULL,
    conversion_offset  BIGINT           NULL,

    upload_id          Text             Null,
    user_id            Text             Null,

    payload            jsonb            Null,
    origin             jsonb            Null,
    annotations        jsonb            Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT           Null,

    delivery_type      TEXT             NULL,
    duration           BIGINT,
    expected_duration  BIGINT,
    rate               DOUBLE PRECISION NULL,
    percent            DOUBLE PRECISION NULL,
    schedule_name      TEXT             NULL
);

CREATE TABLE blood_ketone
(
    time               TIMESTAMPTZ      NOT NULL,

    archived_time      TIMESTAMPTZ      NULL,
    created_time       TIMESTAMPTZ      NULL,
    modified_time      TIMESTAMPTZ      NULL,
    device_time        TIMESTAMPTZ      NULL,

    device_id          TEXT             NULL,
    id                 Text             Null,
    internal_mongo_id  Text             Null,
    guid               Text             Null,

    timezone           Text             Null,
    timezone_offset    BIGINT           NULL,
    clock_drift_offset BIGINT           NULL,
    conversion_offset  BIGINT           NULL,

    upload_id          Text             Null,
    user_id            Text             Null,

    payload            jsonb            Null,
    origin             jsonb            Null,
    annotations        jsonb            Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT           Null,

    value              DOUBLE PRECISION NULL,
    units              Text             Null

);

SELECT create_hypertable('blood_ketone', 'time');


SELECT create_hypertable('basal', 'time');

CREATE TABLE bolus
(
    time               TIMESTAMPTZ      NOT NULL,

    archived_time      TIMESTAMPTZ      NULL,
    created_time       TIMESTAMPTZ      NULL,
    modified_time      TIMESTAMPTZ      NULL,
    device_time        TIMESTAMPTZ      NULL,

    device_id          TEXT             NULL,
    id                 Text             Null,
    internal_mongo_id  Text             Null,
    guid               Text             Null,

    timezone           Text             Null,
    timezone_offset    BIGINT           NULL,
    clock_drift_offset BIGINT           NULL,
    conversion_offset  BIGINT           NULL,

    upload_id          Text             Null,
    user_id            Text             Null,

    payload            jsonb            Null,
    origin             jsonb            Null,
    annotations        jsonb            Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT           Null,

    normal             DOUBLE PRECISION NULL,
    expected_normal    DOUBLE PRECISION NULL,

    extended           DOUBLE PRECISION NULL,
    expected_extended  DOUBLE PRECISION NULL,

    duration           DOUBLE PRECISION NULL,
    expected_duration  DOUBLE PRECISION NULL,

    sub_type           Text             Null

);


SELECT create_hypertable('bolus', 'time');

CREATE TABLE cbg
(
    time               TIMESTAMPTZ      NOT NULL,

    archived_time      TIMESTAMPTZ      NULL,
    created_time       TIMESTAMPTZ      NULL,
    modified_time      TIMESTAMPTZ      NULL,
    device_time        TIMESTAMPTZ      NULL,

    device_id          TEXT             NULL,
    id                 Text             Null,
    internal_mongo_id  Text             Null,
    guid               Text             Null,

    timezone           Text             Null,
    timezone_offset    BIGINT           NULL,
    clock_drift_offset BIGINT           NULL,
    conversion_offset  BIGINT           NULL,

    upload_id          Text             Null,
    user_id            Text             Null,

    payload            jsonb            Null,
    origin             jsonb            Null,
    annotations        jsonb            Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT           Null,

    value              DOUBLE PRECISION NULL,

    units              Text             Null
);


SELECT create_hypertable('cbg', 'time');

CREATE TABLE cgm_settings
(
    time                  TIMESTAMPTZ NOT NULL,

    archived_time         TIMESTAMPTZ NULL,
    created_time          TIMESTAMPTZ NULL,
    modified_time         TIMESTAMPTZ NULL,
    device_time           TIMESTAMPTZ NULL,

    device_id             TEXT        NULL,
    id                    Text        Null,
    internal_mongo_id     Text        Null,
    guid                  Text        Null,

    timezone              Text        Null,
    timezone_offset       BIGINT      NULL,
    clock_drift_offset    BIGINT      NULL,
    conversion_offset     BIGINT      NULL,

    upload_id             Text        Null,
    user_id               Text        Null,

    payload               jsonb       Null,
    origin                jsonb       Null,
    annotations           jsonb       Null,

    active                boolean DEFAULT TRUE,

    revision              BIGINT      Null,

    transmitter_id        Text        NULL,
    units                 Text        NULL,

    low_alerts            jsonb       NULL,
    high_alerts           jsonb       NULL,
    rate_of_change_alerts jsonb       NULL,
    out_of_range_alerts   jsonb       NULL
);

SELECT create_hypertable('cgm_settings', 'time');

CREATE TABLE clinics
(
    clinic_id TEXT,
    name      TEXT,
    address   TEXT,
    active    bool
);


CREATE TABLE clinics_clinicians
(
    clinic_id    TEXT,
    clinician_id TEXT,
    active       bool
);


CREATE TABLE clinics_patients
(
    clinic_id  TEXT,
    patient_id TEXT,
    active     bool
);


CREATE TABLE device_event
(
    time               TIMESTAMPTZ      NOT NULL,

    archived_time      TIMESTAMPTZ      NULL,
    created_time       TIMESTAMPTZ      NULL,
    modified_time      TIMESTAMPTZ      NULL,
    device_time        TIMESTAMPTZ      NULL,

    device_id          TEXT             NULL,
    id                 Text             Null,
    internal_mongo_id  Text             Null,
    guid               Text             Null,

    timezone           Text             Null,
    timezone_offset    BIGINT           NULL,
    clock_drift_offset BIGINT           NULL,
    conversion_offset  BIGINT           NULL,

    upload_id          Text             Null,
    user_id            Text             Null,

    payload            jsonb            Null,
    origin             jsonb            Null,
    annotations        jsonb            Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT           Null,

    sub_type           Text             Null,
    units              Text             Null,

    value              DOUBLE PRECISION Null,

    duration           BIGINT           Null,
    reason             jsonb            Null,

    prime_target       Text             Null,
    volume             DOUBLE PRECISION Null

);

SELECT create_hypertable('device_event', 'time');



CREATE TABLE device_meta
(
    time               TIMESTAMPTZ NOT NULL,

    archived_time      TIMESTAMPTZ NULL,
    created_time       TIMESTAMPTZ NULL,
    modified_time      TIMESTAMPTZ NULL,
    device_time        TIMESTAMPTZ NULL,

    device_id          TEXT        NULL,
    id                 Text        Null,
    internal_mongo_id  Text        Null,
    guid               Text        Null,

    timezone           Text        Null,
    timezone_offset    BIGINT      NULL,
    clock_drift_offset BIGINT      NULL,
    conversion_offset  BIGINT      NULL,

    upload_id          Text        Null,
    user_id            Text        Null,

    payload            jsonb       Null,
    origin             jsonb       Null,
    annotations        jsonb       Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT      Null,

    reason             jsonb       Null,
    status             Text        Null,
    sub_type           Text        Null,

    duration           BIGINT      Null
);

SELECT create_hypertable('device_meta', 'time');


CREATE TABLE dosing_decision
(
    time                                             TIMESTAMPTZ NOT NULL,

    archived_time                                    TIMESTAMPTZ NULL,
    created_time                                     TIMESTAMPTZ NULL,
    modified_time                                    TIMESTAMPTZ NULL,
    device_time                                      TIMESTAMPTZ NULL,

    device_id                                        TEXT        NULL,
    id                                               Text        Null,
    internal_mongo_id                                Text        Null,
    guid                                             Text        Null,

    timezone                                         Text        Null,
    timezone_offset                                  BIGINT      NULL,
    clock_drift_offset                               BIGINT      NULL,
    conversion_offset                                BIGINT      NULL,

    upload_id                                        Text        Null,
    user_id                                          Text        Null,

    payload                                          jsonb       Null,
    origin                                           jsonb       Null,
    annotations                                      jsonb       Null,

    active                                           boolean DEFAULT TRUE,

    revision                                         BIGINT      Null,

    insulin_on_board                                 jsonb       Null,
    carbohydrates_on_board                           jsonb       Null,

    blood_glucose_target_schedule                    jsonb       Null,
    blood_glucose_forecast                           jsonb       Null,
    blood_glucose_forecast_including_pending_insulin jsonb       Null,

    recommended_basal                                jsonb       Null,
    recommended_bolus                                jsonb       Null,
    units                                            jsonb       Null
);

SELECT create_hypertable('dosing_decision', 'time');



CREATE TABLE food
(
    time               TIMESTAMPTZ NOT NULL,

    archived_time      TIMESTAMPTZ NULL,
    created_time       TIMESTAMPTZ NULL,
    modified_time      TIMESTAMPTZ NULL,
    device_time        TIMESTAMPTZ NULL,

    device_id          TEXT        NULL,
    id                 Text        Null,
    internal_mongo_id  Text        Null,
    guid               Text        Null,

    timezone           Text        Null,
    timezone_offset    BIGINT      NULL,
    clock_drift_offset BIGINT      NULL,
    conversion_offset  BIGINT      NULL,

    upload_id          Text        Null,
    user_id            Text        Null,

    payload            jsonb       Null,
    origin             jsonb       Null,
    annotations        jsonb       Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT      Null,

    nutrition          jsonb       Null
);

SELECT create_hypertable('food', 'time');


CREATE TABLE insulin
(
    time               TIMESTAMPTZ NOT NULL,

    archived_time      TIMESTAMPTZ NULL,
    created_time       TIMESTAMPTZ NULL,
    modified_time      TIMESTAMPTZ NULL,
    device_time        TIMESTAMPTZ NULL,

    device_id          TEXT        NULL,
    id                 Text        Null,
    internal_mongo_id  Text        Null,
    guid               Text        Null,

    timezone           Text        Null,
    timezone_offset    BIGINT      NULL,
    clock_drift_offset BIGINT      NULL,
    conversion_offset  BIGINT      NULL,

    upload_id          Text        Null,
    user_id            Text        Null,

    payload            jsonb       Null,
    origin             jsonb       Null,
    annotations        jsonb       Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT      Null,

    dose               jsonb       Null
);

SELECT create_hypertable('insulin', 'time');


CREATE TABLE old_clinics_patients
(
    old_clinic_id TEXT,
    patient_id    TEXT
);


CREATE TABLE physical_activity
(
    time               TIMESTAMPTZ NOT NULL,

    archived_time      TIMESTAMPTZ NULL,
    created_time       TIMESTAMPTZ NULL,
    modified_time      TIMESTAMPTZ NULL,
    device_time        TIMESTAMPTZ NULL,

    device_id          TEXT        NULL,
    id                 Text        Null,
    internal_mongo_id  Text        Null,
    guid               Text        Null,

    timezone           Text        Null,
    timezone_offset    BIGINT      NULL,
    clock_drift_offset BIGINT      NULL,
    conversion_offset  BIGINT      NULL,

    upload_id          Text        Null,
    user_id            Text        Null,

    payload            jsonb       Null,
    origin             jsonb       Null,
    annotations        jsonb       Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT      Null,

    duration           jsonb       Null,
    distance           jsonb       Null,
    energy             jsonb       Null,
    name               Text        Null
);

SELECT create_hypertable('physical_activity', 'time');

CREATE TABLE pump_settings
(
    time                  TIMESTAMPTZ NOT NULL,

    archived_time         TIMESTAMPTZ NULL,
    created_time          TIMESTAMPTZ NULL,
    modified_time         TIMESTAMPTZ NULL,
    device_time           TIMESTAMPTZ NULL,

    device_id             TEXT        NULL,
    id                    Text        Null,
    internal_mongo_id     Text        Null,
    guid                  Text        Null,

    timezone              Text        Null,
    timezone_offset       BIGINT      NULL,
    clock_drift_offset    BIGINT      NULL,
    conversion_offset     BIGINT      NULL,

    upload_id             Text        Null,
    user_id               Text        Null,

    payload               jsonb       Null,
    origin                jsonb       Null,
    annotations           jsonb       Null,

    active                boolean DEFAULT TRUE,

    revision              BIGINT      Null,

    active_schedule       Text        Null,
    basal_schedules       jsonb       Null,
    bg_targets            jsonb       Null,
    carb_ratios           jsonb       Null,
    insulin_sensitivities jsonb       Null,
    units                 jsonb       Null,

    manufacturers         Text[],
    model                 Text        Null,
    serial_number         Text        Null
);

SELECT create_hypertable('pump_settings', 'time');


CREATE TABLE reported_state
(
    time               TIMESTAMPTZ NOT NULL,

    archived_time      TIMESTAMPTZ NULL,
    created_time       TIMESTAMPTZ NULL,
    modified_time      TIMESTAMPTZ NULL,
    device_time        TIMESTAMPTZ NULL,

    device_id          TEXT        NULL,
    id                 Text        Null,
    internal_mongo_id  Text        Null,
    guid               Text        Null,

    timezone           Text        Null,
    timezone_offset    BIGINT      NULL,
    clock_drift_offset BIGINT      NULL,
    conversion_offset  BIGINT      NULL,

    upload_id          Text        Null,
    user_id            Text        Null,

    payload            jsonb       Null,
    origin             jsonb       Null,
    annotations        jsonb       Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT      Null,

    states             jsonb       Null

);

SELECT create_hypertable('reported_state', 'time');



CREATE TABLE settings
(
    time                TIMESTAMPTZ NOT NULL,

    archived_time       TIMESTAMPTZ NULL,
    created_time        TIMESTAMPTZ NULL,
    modified_time       TIMESTAMPTZ NULL,
    device_time         TIMESTAMPTZ NULL,

    device_id           TEXT        NULL,
    id                  Text        Null,
    internal_mongo_id   Text        Null,
    guid                Text        Null,

    timezone            Text        Null,
    timezone_offset     BIGINT      NULL,
    clock_drift_offset  BIGINT      NULL,
    conversion_offset   BIGINT      NULL,

    upload_id           Text        Null,
    user_id             Text        Null,

    payload             jsonb       Null,
    origin              jsonb       Null,
    annotations         jsonb       Null,

    active              boolean DEFAULT TRUE,

    revision            BIGINT      Null,

    source              TEXT        NULL,
    group_id            TEXT        NULL,
    active_schedule     TEXT        NULL,
    units               jsonb       NULL,
    basal_schedules     jsonb       NULL,
    insulin_sensitivity jsonb       NULL,
    carb_ratio          jsonb       NULL
);

SELECT create_hypertable('settings', 'time');



CREATE TABLE smbg
(
    time               TIMESTAMPTZ      NOT NULL,

    archived_time      TIMESTAMPTZ      NULL,
    created_time       TIMESTAMPTZ      NULL,
    modified_time      TIMESTAMPTZ      NULL,
    device_time        TIMESTAMPTZ      NULL,

    device_id          TEXT             NULL,
    id                 Text             Null,
    internal_mongo_id  Text             Null,
    guid               Text             Null,

    timezone           Text             Null,
    timezone_offset    BIGINT           NULL,
    clock_drift_offset BIGINT           NULL,
    conversion_offset  BIGINT           NULL,

    upload_id          Text             Null,
    user_id            Text             Null,

    payload            jsonb            Null,
    origin             jsonb            Null,
    annotations        jsonb            Null,

    active             boolean DEFAULT TRUE,

    revision           BIGINT           Null,

    value              DOUBLE PRECISION NULL,

    units              Text             Null

);

SELECT create_hypertable('smbg', 'time');

CREATE TABLE upload
(
    time                 TIMESTAMPTZ NULL,

    archived_time        TIMESTAMPTZ NULL,
    created_time         TIMESTAMPTZ NULL,
    modified_time        TIMESTAMPTZ NULL,
    device_time          TIMESTAMPTZ NULL,

    device_id            TEXT        NULL,
    id                   Text        Null,
    internal_mongo_id    Text        Null,
    guid                 Text        Null,

    timezone             Text        Null,
    timezone_offset      BIGINT      NULL,
    clock_drift_offset   BIGINT      NULL,
    conversion_offset    BIGINT      NULL,

    upload_id            Text        Null,
    user_id              Text        Null,

    payload              jsonb       Null,
    origin               jsonb       Null,
    annotations          jsonb       Null,

    active               boolean DEFAULT TRUE,

    revision             BIGINT      Null,

    data_set_type        Text        Null,
    data_state           Text        Null,

    device_manufacturers Text[],
    device_model         Text        Null,
    device_serial_number Text        Null,

    device_tags          Text[],

    state                Text        Null,
    version              Text        Null
);


CREATE TABLE users
(
    user_id       TEXT,
    username      TEXT,
    authenticated bool
);


CREATE TABLE wizard
(
    time                TIMESTAMPTZ      NOT NULL,

    archived_time       TIMESTAMPTZ      NULL,
    created_time        TIMESTAMPTZ      NULL,
    modified_time       TIMESTAMPTZ      NULL,
    device_time         TIMESTAMPTZ      NULL,

    device_id           TEXT             NULL,
    id                  Text             Null,
    internal_mongo_id   Text             Null,
    guid                Text             Null,

    timezone            Text             Null,
    timezone_offset     BIGINT           NULL,
    clock_drift_offset  BIGINT           NULL,
    conversion_offset   BIGINT           NULL,

    upload_id           Text             Null,
    user_id             Text             Null,

    payload             jsonb            Null,
    origin              jsonb            Null,
    annotations         jsonb            Null,

    active              boolean DEFAULT TRUE,

    revision            BIGINT           Null,

    bolus               Text             Null,
    units               Text             Null,

    recommended         jsonb            NULL,

    bg_input            DOUBLE PRECISION Null,
    bg_target           jsonb            NULL,

    carb_input          BIGINT           NULL,
    insulin_carb_ratio  BIGINT           NULL,

    insulin_on_board    DOUBLE PRECISION NULL,
    insulin_sensitivity DOUBLE PRECISION NULL
);

SELECT create_hypertable('wizard', 'time');

