#! /bin/bash
set -e
# Tag passed to the script
TAG=$1
# set goreleaser env var for custom tag
echo "GORELEASER_CURRENT_TAG=v${TAG##*-}" >> $GITHUB_ENV

