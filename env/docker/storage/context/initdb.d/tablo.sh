#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "${POSTGRES_USER}" <<-EOSQL
    CREATE USER "tablo" WITH PASSWORD 'tablo';
    CREATE DATABASE "tablo" WITH OWNER "tablo";
    \c "tablo";
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL

psql -v ON_ERROR_STOP=1 -U tablo -d tablo -f /docker-entrypoint-initdb.d/prototype.sql -a
