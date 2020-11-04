

CREATE INDEX on basal (user_id, time);
CREATE INDEX on basal (internal_mongo_id, time);
CREATE INDEX on blood_ketone (user_id, time);
CREATE INDEX on blood_ketone (internal_mongo_id, time);
CREATE INDEX on bolus (user_id, time);
CREATE INDEX on bolus (internal_mongo_id, time);
CREATE INDEX on cbg (user_id, time, value);
CREATE INDEX on cbg (internal_mongo_id, time);
CREATE INDEX on cgm_settings (user_id, time);
CREATE INDEX on cgm_settings (internal_mongo_id, time);
CREATE INDEX on device_event (user_id, time);
CREATE INDEX on device_event (internal_mongo_id, time);
CREATE INDEX on device_meta (user_id, time);
CREATE INDEX on device_meta (internal_mongo_id, time);
CREATE INDEX on dosing_decision (user_id, time);
CREATE INDEX on dosing_decision (internal_mongo_id, time);
CREATE INDEX on food (user_id, time);
CREATE INDEX on food (internal_mongo_id, time);
CREATE INDEX on insulin (user_id, time);
CREATE INDEX on insulin (internal_mongo_id, time);
CREATE INDEX on physical_activity (user_id, time);
CREATE INDEX on physical_activity (internal_mongo_id, time);
CREATE INDEX on pump_settings (user_id, time);
CREATE INDEX on pump_settings (internal_mongo_id, time);
CREATE INDEX on reported_state (user_id, time);
CREATE INDEX on reported_state (internal_mongo_id, time);
CREATE INDEX on settings (user_id, time);
CREATE INDEX on settings (internal_mongo_id, time);
CREATE INDEX on smbg (user_id, time);
CREATE INDEX on smbg (internal_mongo_id, time);
CREATE INDEX on upload (user_id, time);
CREATE INDEX on upload (internal_mongo_id, time);
CREATE INDEX on wizard (user_id, time);
CREATE INDEX on wizard (internal_mongo_id, time);

