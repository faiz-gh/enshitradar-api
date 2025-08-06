CREATE TYPE level_enum AS ENUM ('low', 'medium', 'high', 'confirmed');

CREATE TABLE
    IF NOT EXISTS channels (
        id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
        channel_id character varying NOT NULL,
        name character varying NOT NULL,
        level level_enum DEFAULT 'low',
        description character varying,
        date_added timestamp
        with
            time zone DEFAULT NOW ()
    );