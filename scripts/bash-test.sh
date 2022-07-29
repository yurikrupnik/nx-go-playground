#! /bin/bash
set -e
echo "GORELEASER_CURRENT_TAG=v${$1##*-}" >> $GITHUB_ENV

