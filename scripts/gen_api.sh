#!/bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
ROOT_DIR=$(dirname "$SCRIPT_DIR")

if ! command -v openapi-generator &> /dev/null
then
    echo "openapi-generator could not be found"
    brew install openapi-generator
    if ! command -v oapi-codegen &> /dev/null
    then
      echo "unable to install openapi-generator"
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
  openapi-generator generate -i reference/SpaceTraders.json -g go -o "$ROOT_DIR/internal/api/"  --git-user-id nandiheath  --git-repo-id spacetraders --package-name api --additional-properties=isGoSubmodule=true --additional-properties=enumClassPrefix=true
  # remove go.mod as it is not a separate model
  rm "$ROOT_DIR/internal/api/go.mod"
)

# there is a bug on the generated client.go and this is a hacky fix for it..
cp -f "$ROOT_DIR/scripts/templates/client.go"  "$ROOT_DIR/internal/api"

