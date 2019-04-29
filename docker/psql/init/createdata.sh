#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE ROLE hotpotbot LOGIN PASSWORD 'hotpotbot';
    CREATE DATABASE hotpotbot_db;
    GRANT ALL ON DATABASE hotpotbot_db TO hotpotbot;
    CREATE TABLE member (NAME VARCHAR(32) NOT NULL, COMPANY VARCHAR(32) NOT NULL, JOBTYPE INT NOT NULL);
EOSQL
