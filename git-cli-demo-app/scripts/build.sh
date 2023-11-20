#!/usr/bin/env bash

set -o pipefail
set -eu

go build -o git-cli ./cmd/...
