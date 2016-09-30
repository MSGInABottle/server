CREATE DATABASE messageinabottle;

\c messageinabottle

CREATE TABLE MESSAGE(
    text TEXT NOT NULL,
    location POINT NOT NULL,
    expiry TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE USER goserver;
ALTER USER goserver WITH PASSWORD '';
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public to goserver;
