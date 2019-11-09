#!/usr/bin/env bash

set -e

psql -v ON_ERROR_STOP=1 --username "${POSTGRES_USER}" <<-EOSQL
    CREATE USER "sentry" WITH PASSWORD 'sentry';
    CREATE DATABASE "sentry" WITH OWNER "sentry";
EOSQL
