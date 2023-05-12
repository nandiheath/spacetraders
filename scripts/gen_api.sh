#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
ROOT_DIR=$(dirname "$SCRIPT_DIR")

if ! command -v  oapi-codegen &> /dev/null
then
    echo " oapi-codegen  could not be found"
    go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
    if ! command -v oapi-codegen &> /dev/null
    then
      echo "unable to install  oapi-codegen"
      exit 1
    fi
    echo "oapi-codegen installed"
fi

temp_dir=$(mktemp -d)

clean_up () {
#  rm -rf "$temp_dir"
  exit 0
}
trap clean_up EXIT

(
  cd "$temp_dir"
  git clone --depth=1 https://github.com/SpaceTradersAPI/api-docs
  cd "api-docs"
  oapi-codegen -generate types,client -package api -templates "$ROOT_DIR/oapi-gen-templates/" -o "$ROOT_DIR/internal/api/api.go" "$ROOT_DIR/api_ref.json"
)


