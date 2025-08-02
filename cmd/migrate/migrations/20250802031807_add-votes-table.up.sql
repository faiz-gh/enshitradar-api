CREATE TABLE
    votes (
        id uuid PRIMARY KEY DEFAULT gen_random_uuid (),
        channel_id uuid NOT NULL,
        user_id uuid NOT NULL,
        level level_enum NOT NULL,
        public_ip_address character varying,
        date_added timestamp
        with
            time zone DEFAULT NOW ()
    );

ALTER TABLE votes ADD CONSTRAINT fk_vote_channel FOREIGN KEY (channel_id) REFERENCES channels (id) ON DELETE CASCADE;

ALTER TABLE votes ADD CONSTRAINT fk_vote_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE;