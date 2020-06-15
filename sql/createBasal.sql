CREATE TABLE basal (
  time        TIMESTAMPTZ       NOT NULL,

  uploadId    TEXT              NOT NULL,

  deliveryType    TEXT              NOT NULL,
  duration        INT,
  expectedDuration        INT,
  rate DOUBLE PRECISION  NULL,
  percent DOUBLE PRECISION  NULL,
  scheduleName    TEXT              NOT NULL
);

SELECT create_hypertable('basal', 'time');

