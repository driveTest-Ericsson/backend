#!/bin/sh
set -e
make migrate-up
make seed
./main