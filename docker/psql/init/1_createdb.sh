#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username postgres <<-EOSQL
    CREATE ROLE hotpotbot WITH LOGIN PASSWORD 'hotpotbot';
EOSQL
