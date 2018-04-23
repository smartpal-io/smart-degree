#!/usr/bin/env bash
set -xeuo pipefail

docker build -t androiddocker.image -f ./scripts/Dockerfile .