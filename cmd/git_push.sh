#!/usr/bin/env bash

# shellcheck disable=SC2006
BRANCH=`git rev-parse --abbrev-ref HEAD`

echo "Pushing to branch ${BRANCH}"
echo "-----------------------------------------"
echo -n "Enter commit message: "
# shellcheck disable=SC2162
read COMMIT

git add .
git commit -m "$COMMIT"
git push origin "${BRANCH}"
