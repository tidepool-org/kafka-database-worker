CREATE TABLE basal (
  time               TIMESTAMPTZ       NOT NULL,

  upload_id          TEXT NULL,

  delivery_type      TEXT NULL,
  duration           INT,
  expected_duration  INT,
  rate               DOUBLE PRECISION  NULL,
  percent            DOUBLE PRECISION  NULL,
  schedule_name      TEXT NULL
);

SELECT create_hypertable('basal', 'time');

