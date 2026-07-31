#!/usr/bin/env bash

CURRENT_BRANCH=$(git branch --show-current)

echo "${CURRENT_BRANCH}"

if [ -z "$1" ]; then
    echo "Error: You should add a commit message"
    exit 1
fi

git add -A
git commit -m "$1"
git push origin "${CURRENT_BRANCH}"