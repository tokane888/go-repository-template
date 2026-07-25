#!/bin/bash
set -euxo pipefail

WORKSPACE_FOLDER="$1"
WORKSPACE_BASENAME="$2"
ENV_FILE="${WORKSPACE_FOLDER}/.devcontainer/.env"

MAIN_GIT_DIR=$(realpath "$(git rev-parse --git-common-dir)")

# Output environment variable to .env file.
# .env file is read by docker-compose.yml.
{
  echo "WORKSPACE_BASENAME=${WORKSPACE_BASENAME}"
  echo "MAIN_GIT_DIR=${MAIN_GIT_DIR}"
} > "${ENV_FILE}"
