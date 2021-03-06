#!/usr/bin/env bash

declare basePath="$1"
find "$basePath" -type d -name ".git" \
  | while read -r gitDirectoryPath; do
    declare repoPath="$(dirname "$gitDirectoryPath")"
    pushd "$repoPath" > /dev/null

    declare name="$(basename "$repoPath")"
    declare lastCommit=$(git rev-parse HEAD)
    printf '{"name":"%s","lastCommit":"%s"}\n' $name $lastCommit

    popd > /dev/null
  done \
  | paste -s -d"," \
  | sed -E "s/(.*)/[\1]/"
