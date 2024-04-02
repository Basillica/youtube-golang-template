package sql

const (
	CREATE_TABLE = `
		BEGIN;
		CREATE TABLE IF NOT EXISTS sensor_data (
			timestamp TIMESTAMPTZ not null,
			sensor_id varchar(50) not null,
			name varchar(50) not null,
			topic varchar(50) not null,
			value varchar(8) not null
		);
		SELECT create_hypertable('sensor_data', 'timestamp');
		COMMIT;
	`

	PG_CRON = `CREATE EXTENSION IF NOT EXISTS pg_cron;`

	TIMESCALE = `CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;`

	STATS = `CREATE EXTENSION IF NOT EXISTS pg_stat_statements CASCADE;`

	HTTP = `CREATE EXTENSION IF NOT EXISTS http CASCADE;`

	AGGREGATION_MATERIAL_VIEW = `
		CREATE MATERIALIZED VIEW IF NOT EXISTS total_summary_daily 
		WITH (timescaledb.continuous) 
		AS SELECT 
		time_bucket(INTERVAL '1 day', timestamp) AS bucket, 
		AVG(total_amount), 
		MAX(total_amount),
		MIN(total_amount) 
		FROM sensor_data
		WHERE total_amount > 0
		GROUP BY bucket
		WITH NO DATA;
		
		SELECT add_continuous_aggregate_policy('total_summary_daily',
		start_offset => INTERVAL '1 year',
		end_offset => INTERVAL '1 month',
		schedule_interval => INTERVAL '1 hour');
	`

	TRIGGER = `
		CREATE OR REPLACE FUNCTION notify_trigger() RETURNS trigger AS $trigger$
		DECLARE
		rec RECORD;
		dat RECORD;
		payload TEXT;
		BEGIN

		-- Set record row depending on operation
		CASE TG_OP
		WHEN 'UPDATE' THEN
			rec := NEW;
			dat := OLD;
		WHEN 'INSERT' THEN
			rec := NEW;
		WHEN 'DELETE' THEN
			rec := OLD;
		ELSE
			RAISE EXCEPTION 'Unknown TG_OP: "%". Should not occur!', TG_OP;
		END CASE;
		
		-- Build the payload
		payload := json_build_object('timestamp',CURRENT_TIMESTAMP,'action',LOWER(TG_OP),'schema',TG_TABLE_SCHEMA,'public',TG_TABLE_NAME,'sensor_data',row_to_json(rec), 'old',row_to_json(dat));

		-- Notify the channel
		PERFORM pg_notify('core_db_event',payload);
		
		RETURN rec;
		END;
		$trigger$ LANGUAGE plpgsql;

		DROP TRIGGER IF EXISTS event_notify ON public.sensor_data;

		CREATE TRIGGER event_notify
		AFTER INSERT ON public.sensor_data
		FOR EACH ROW EXECUTE PROCEDURE notify_trigger();
	`
)
