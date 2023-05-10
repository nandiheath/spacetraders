#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
ROOT_DIR=$(dirname "$SCRIPT_DIR")

[[ ! -f "$ROOT_DIR/.env" ]] && echo "cannot find .env . Run `cp .env-sample .env` and configure it. " && exit 1;

source "$ROOT_DIR/.env"

[[ -z "$CALLSIGN" ]] && echo "CALLSIGN not set. check your .env" && exit 1;

response=$(curl --request POST \
 --url 'https://api.spacetraders.io/v2/register' \
 --header 'Content-Type: application/json' \
 --data '{
         "symbol": "'"$CALLSIGN"'",
         "faction": "COSMIC"
        }')

# TODO: jq the response :D
echo "$response"
echo "please save your access token to .env!"
