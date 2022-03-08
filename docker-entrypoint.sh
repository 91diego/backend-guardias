#!/usr/bin/env bash
set -ex

wait-for-it "${DB_HOST}:${DB_PORT}"

exec "$@"
