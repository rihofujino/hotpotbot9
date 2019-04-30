
#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username postgres --dbname hotpotbot_db <<-EOSQL   
    CREATE TABLE member (
        ID SERIAL,
        NAME VARCHAR(32) NOT NULL, 
        COMPANY VARCHAR(32), 
        JOBTYPE INT,
        CREATED_AT TIMESTAMP
    );
    CREATE TABLE survey (
        ID SERIAL,
        NAME VARCHAR(32) NOT NULL, 
        SATISFACTION INT NOT NULL,
        IMPRESSION TEXT, 
        EXPECT_THEME TEXT,
        CREATED_AT TIMESTAMP
    );
    GRANT ALL ON ALL TABLES IN SCHEMA public TO hotpotbot;
    GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public TO hotpotbot;
EOSQL



