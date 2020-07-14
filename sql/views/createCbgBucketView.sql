CREATE VIEW cbg_regular_ts_view
            WITH (timescaledb.continuous) AS
SELECT user_id,
       interpolate(avg(value)) as value,
       time_bucket(INTERVAL '1 hour', time) AS bucket
FROM cbg
GROUP BY user_id, bucket;
