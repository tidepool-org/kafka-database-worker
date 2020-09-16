

CREATE INDEX on basal (user_id, time);
CREATE INDEX on bolus (user_id, time);
CREATE INDEX on cbg (user_id, time, value);
CREATE INDEX on cgm_settings (user_id, time);
CREATE INDEX on device_event (user_id, time);
CREATE INDEX on device_meta (user_id, time);
CREATE INDEX on food (user_id, time);
CREATE INDEX on physical_activity (user_id, time);
CREATE INDEX on pump_settings (user_id, time);
CREATE INDEX on smbg (user_id, time);
CREATE INDEX on upload (user_id, time);
CREATE INDEX on wizard (user_id, time);

