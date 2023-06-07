CREATE TABLE IF NOT EXISTS events(
    id SERIAL PRIMARY KEY,
    title VARCHAR,
    description VARCHAR,
    location VARCHAR,
    date TIMESTAMP,
    organiser_id BIGINT,
    created_at TIMESTAMP DEFAULT NOW(),

    CONSTRAINT fk_organiser
        FOREIGN KEY(organiser_id)
            REFERENCES users(id)
);