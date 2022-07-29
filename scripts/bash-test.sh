
#! /bin/bash
set -e
TT=$1
echo $1
#TAG_COPY=users-api-0.0.3
TAG=v${TT##*-}

echo "GORELEASER_CURRENT_TAG=v${TT##*-}" >> $GITHUB_ENV
#echo export TAG
#export $TAG
#echo $TAG
#MONGO_URI=$(gcloud secrets versions access latest --secret=MONGO_URI --project=mussia8)
#
#echo $MONGO_URI
#
#export MYSQL_PASSWORD=gcp:///MONGO_URI
#gcp-get-secret bash -c 'echo $MYSQL_PASSWORD'

#echo Branch name: ${GITHUB_REF##*/}
#echo HEAD_REF: $HEAD_REF
#echo sha: ${GITHUB_SHA::8}
#echo BRA: ${BRA}
#echo BRA: $BRA
#
#
#if [[ $BRA == *"/"* ]]; then
#    HEAD_REF=${GITHUB_SHA::8}
#    echo "It's there!"
#    echo $HEAD_REF
#  fi
