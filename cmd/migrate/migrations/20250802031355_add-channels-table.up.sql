CREATE TYPE level_enum AS ENUM ('low', 'high', 'confirmed');

CREATE TABLE
    IF NOT EXISTS channels (
        id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
        name character varying NOT NULL,
        level level_enum DEFAULT 'low',
        description character varying,
        date_added timestamp
        with
            time zone DEFAULT NOW ()
    );