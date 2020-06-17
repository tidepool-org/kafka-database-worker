CREATE TABLE basal (
  time        TIMESTAMPTZ       NOT NULL,

  uploadId    TEXT NULL,

  deliveryType    TEXT NULL,
  duration        INT,
  expectedDuration        INT,
  rate DOUBLE PRECISION  NULL,
  percent DOUBLE PRECISION  NULL,
  scheduleName    TEXT NULL
);

SELECT create_hypertable('basal', 'time');

