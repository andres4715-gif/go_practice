#!/usr/bin/env bash

CURRENT_BRANCH=git branch --show-current

git add -A
git commit -m "$1"
git push origin ${CURRENT_BRANCH}
