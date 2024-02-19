#!/bin/bash
set -a # automatically export all variables
source /bin/.env
set +a

exec /bin/TideTracker
