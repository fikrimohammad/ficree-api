CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    guid UUID NOT NULL DEFAULT uuid_generate_v1(),
    email text NOT NULL UNIQUE,
    name text NOT NULL,
    phone_number text NOT NULL UNIQUE,
    profile_picture text,
    github_url text,
    linkedin_url text,
    twitter_url text,
    summary text,
    title text,
    created_at timestamp with time zone NOT NULL DEFAULT NOW(),
    updated_at timestamp with time zone NOT NULL DEFAULT NOW(),
    deleted_at timestamp with time zone
);